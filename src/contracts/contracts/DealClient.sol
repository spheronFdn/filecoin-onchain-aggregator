// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import {Cid} from "./libraries/Cid.sol";
import {BytesLib} from "./libraries/BytesLib.sol";

struct RequestId {
    bytes32 requestId;
    bool valid;
}

struct RequestIdx {
    uint256 idx;
    bool valid;
}

struct DealRequest {
    bytes piece_cid;
    string label;
    string fetchLink;
}

// User request for this contract to make a deal. This structure is modelled after Filecoin's Deal
// Proposal, but leaves out the provider, since any provider can pick up a deal broadcast by this
// contract.
struct DealBatch {
    uint256 batch_id;
    uint256 commPcEndIndex;
    bytes32 commpa;
    bytes deal_id; // A unique identifier for the deal.
    bytes provider_id; // The storage provider that is storing the data for the deal.
    bytes market_actor_id; // The actor ID of the storage market in which the deal representation is stored. dealId is scoped by marketActorId
}

contract DealClient {
    using Cid for bytes32;
    using Cid for bytes;
    using BytesLib for bytes;

    bytes32 constant TRUNCATOR =
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3f;
    uint256 constant NodeSize = 32;
    uint256 constant ChecksumSize = 16;

    enum Status {
        None,
        RequestSubmitted,
        DealPublished,
        DealActivated,
        DealTerminated
    }

    address public owner;
    address public aggregator;

    DealBatch[] public batches;
    DealRequest[] public dealRequests;

    mapping(bytes32 => RequestIdx) public dealRequestIdx; // contract deal id -> deal index
    mapping(bytes => RequestId) public pieceRequests; // commP -> dealProposalID
    mapping(bytes => Status) public pieceStatus;
    mapping(bytes => uint256) public pieceToBatchId;

    event ReceivedDataCap(string received);
    event DealProposalCreate(bytes32 indexed id, bytes pieceCid);
    event BatchIdLeafNodes(
        uint256 indexed batchId,
        bytes32 indexed node,
        bytes32 pieceCidLeft,
        bytes32 pieceCidRight
    );

    constructor() {
        owner = msg.sender;
        aggregator = msg.sender;
    }

    modifier onlyAggregator() {
        require(aggregator == msg.sender, "!aggregator");
        _;
    }

    function setAggregator(address _newAggregator) external onlyAggregator {
        aggregator = _newAggregator;
    }

    function submit(
        bytes memory _cid,
        string memory label,
        string memory fetchLink
    ) public returns (bytes32) {
        DealRequest memory deal;

        deal.piece_cid = _cid;
        deal.label = label;
        deal.fetchLink = fetchLink;

        if (
            pieceStatus[deal.piece_cid] == Status.DealPublished ||
            pieceStatus[deal.piece_cid] == Status.DealActivated
        ) {
            revert("deal with this pieceCid already published");
        }

        uint256 index = dealRequests.length;

        dealRequests.push(deal);

        // creates a unique ID for the deal proposal -- there are many ways to do this
        bytes32 id = keccak256(
            abi.encodePacked(block.timestamp, msg.sender, index)
        );
        dealRequestIdx[id] = RequestIdx(index, true);

        pieceRequests[deal.piece_cid] = RequestId(id, true);
        pieceStatus[deal.piece_cid] = Status.RequestSubmitted;

        // writes the proposal metadata to the event log
        emit DealProposalCreate(id, deal.piece_cid);

        return id;
    }

    function complete(
        bytes memory _marketActorId,
        bytes memory _dealId,
        bytes memory _providerId,
        uint256 commPcStartIndex,
        uint256 commPcEndIndex
    )
        public
        onlyAggregator
        returns (bytes32 _root, uint256 _count, uint256 batchId)
    {
        require(msg.sender == owner);
        batchId = batches.length;

        DealBatch memory batch;

        batch.deal_id = _dealId;
        batch.market_actor_id = _marketActorId;
        batch.provider_id = _providerId;
        batch.batch_id = batchId;
        batch.commPcEndIndex = commPcEndIndex;

        bytes32[] memory nodes;

        uint256 nodeCount = 0;
        for (uint256 i = commPcStartIndex; i <= commPcEndIndex; i++) {
            bytes memory deal_piece_cid = dealRequests[i].piece_cid;

            nodes[2 * i] = bytes32(deal_piece_cid.slice(0, NodeSize));
            nodes[2 * i + 1] = bytes32(deal_piece_cid.slice(NodeSize, 0));

            pieceToBatchId[deal_piece_cid] = batchId;
            pieceStatus[deal_piece_cid] = Status.DealPublished;

            nodeCount = nodeCount + 2;
        }

        require(nodes.length > 0, "Leaf data is required");
        require(
            nodes.length < 2 ** 63,
            "Merkle trees with depths greater than 63 are not supported"
        );

        while (nodes.length > 1) {
            bytes32[] memory newNodes = new bytes32[](
                nodes.length / 2 + (nodes.length % 2)
            );
            for (uint256 i = 0; i < nodes.length; i += 2) {
                if (i + 1 < nodes.length) {
                    bytes32 leftLeaf = nodes[i];
                    bytes32 rightLeaf = nodes[i + 1];
                    bytes32 node = computeNode(leftLeaf, rightLeaf);
                    newNodes[i / 2] = node;
                    emit BatchIdLeafNodes(batchId, node, leftLeaf, rightLeaf);
                } else {
                    // If there's an odd number of nodes, duplicate the last one
                    bytes memory zeroComms;
                    bytes32 zeroCommNode = bytes32(
                        zeroComms.slice(32 * i, 32 * (i + 1))
                    );
                    bytes32 leftLeaf = nodes[i];
                    bytes32 rightLeaf = zeroCommNode;
                    bytes32 node = computeNode(leftLeaf, rightLeaf);
                    newNodes[i / 2] = node;
                    emit BatchIdLeafNodes(batchId, node, leftLeaf, rightLeaf);
                }
            }
            nodes = newNodes;
        }

        _root = nodes[0];

        batch.commpa = _root;
        batches.push(batch);

        return (_root, _count, batchId);
    }

    function getProposalIdSet(
        bytes memory cid
    ) public view returns (RequestId memory) {
        return pieceRequests[cid];
    }

    function dealsLength() public view returns (uint256) {
        return dealRequests.length;
    }

    function batchesLength() public view returns (uint256) {
        return batches.length;
    }

    function getDealByIndex(
        uint256 index
    ) public view returns (DealRequest memory) {
        return dealRequests[index];
    }

    /**
     * @notice Retrieves all deal IDs associated with a specified piece CID
     * @param _cid The piece CID to query
     * @return Array of deals corresponding to the piece CID
     */
    function getAllDeals(
        bytes memory _cid
    ) external view returns (DealBatch memory) {
        uint256 deal_batch_id = pieceToBatchId[_cid];
        DealBatch memory deal = batches[deal_batch_id];
        return deal;
    }

    // helper function to get deal request based from id
    function getDealRequest(
        bytes32 requestId
    ) internal view returns (DealRequest memory) {
        RequestIdx memory ri = dealRequestIdx[requestId];
        require(ri.valid, "proposalId not available");
        return dealRequests[ri.idx];
    }

    /**
     * @dev computeNode computes the parent node of two child nodes
     * @param _left is the left child
     * @param _right is the right child
     * @return the parent node
     */
    function computeNode(
        bytes32 _left,
        bytes32 _right
    ) public pure returns (bytes32) {
        bytes32 digest = sha256(abi.encodePacked(_left, _right));
        return truncate(digest);
    }

    function getCIDFromPieceCid(
        bytes32 piece_cid
    ) public pure returns (bytes memory) {
        return piece_cid.pieceCommitmentToCid();
    }

    function getPieceCidfromCid(
        bytes memory _cid
    ) public pure returns (bytes32) {
        return _cid.cidToPieceCommitment();
    }

    /**
     * @dev truncate truncates a node to 254 bits.
     * @param _n is the node
     * @return the truncated node
     */
    function truncate(bytes32 _n) internal pure returns (bytes32) {
        return _n & TRUNCATOR;
    }
}
