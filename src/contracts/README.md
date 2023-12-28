# On-Chain Aggregator Filecoin Contracts

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.js
```

### Functions

#### `submit` Function

The `submit` function allows users to submit a deal request by providing the necessary details:

```solidity
function submit(
    bytes memory _cid,
    string memory label,
    string memory fetchLink
) public returns (bytes32);
```

- **_cid_**: The content identifier (CID) of the data piece to be stored.
- **label**: A string label associated with the deal.
- **fetchLink**: A string representing the link to fetch the stored data. Can be a IPFS link as well.
  The function generates a unique deal ID and emits an event indicating the creation of a deal proposal.

#### `complete` Function

The `complete` function is used by the aggregator to finalize and publish a batch of deals:

```
function complete(
    bytes memory _marketActorId,
    bytes memory _dealId,
    bytes memory _providerId,
    uint256 commPcStartIndex,
    uint256 commPcEndIndex
) public onlyAggregator returns (bytes32 _root, uint256 _count, uint256 batchId);
```

**\_marketActorId:** The actor ID of the storage market.
**\_dealId:** A unique identifier for the deal.
**\_providerId:** The storage provider's identifier.
**commPcStartIndex:** The starting index of deal requests in the batch.
**commPcEndIndex:** The ending index of deal requests in the batch.

The function computes a Merkle tree root from the provided data and updates the contract's state. It returns the Merkle root, the count of processed deals, and the batch ID. Can onnly be called by the aggregator
