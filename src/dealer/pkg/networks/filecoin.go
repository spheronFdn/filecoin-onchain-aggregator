package networks

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"storage/pkg/config"
	"storage/pkg/constants"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

var FileApi lotusapi.FullNodeStruct

func ConnectToFilecoin() {
	fileCoinNode := config.FIL_NODE
	log.Println("Connecting to Filecoin Node... ", fileCoinNode)

	headers := http.Header{"Authorization": []string{"Bearer " + config.BEARER_TOKEN_FILECOIN_NODE}}
	_, err := jsonrpc.NewMergeClient(context.Background(), "http://"+fileCoinNode+"/rpc/v0", "Filecoin", []interface{}{&FileApi.Internal, &FileApi.CommonStruct.Internal}, headers)
	if err != nil {
		panic("connecting failed")
	}
	// defer closer()
	// Now you can call any API you're interested in.
	tipset, err := FileApi.ChainHead(context.Background())
	if err != nil {
		panic("calling chain head")
	}
	log.Println("Successfully connected to Filecoin Node, Current chain head is: \n\n", tipset.String())
}

func CreateDealForFile(fileName string) (string, *big.Int, uint64, address.Address, error) {
	log.Println("Importing file to Client", fileName)
	ir, err := FileApi.ClientImport(context.Background(), lotusapi.FileRef{Path: fmt.Sprintf("%s/%s", config.WORKING_DIR, "output.car"), IsCAR: true})
	if err != nil {
		log.Println("Import file err:", err)
		return "", nil, 0, address.Address{}, err
	}
	log.Println("import_result cid: ", ir.Root)

	lowestBid := types.NewInt(math.MaxInt32)
	var cheapestMiner address.Address

	for _, minerID := range constants.TopMiners["Devnet"] {
		log.Println("minerID: ", minerID)
		// 2. query storage provider's offer for storing this file
		minerAddr, err := string2Address(minerID)
		if err != nil {
			return "", nil, 0, address.Address{}, err
		}
		log.Println("minerAddr: ", minerAddr)
		qo, err := FileApi.ClientMinerQueryOffer(context.Background(), *minerAddr, ir.Root, nil)
		log.Println("qo: ", qo)
		if err != nil || qo.Err != "" {
			log.Println("miner query offer err: \n", err)
			return "", nil, 0, address.Address{}, err
		}

		if qo.MinPrice.Cmp(lowestBid.Int) == -1 {
			lowestBid = qo.MinPrice
			cheapestMiner = *minerAddr
		}
	}
	log.Println("lowestBid: ", lowestBid)
	log.Println("cheapestMiner: ", cheapestMiner)

	// 3. start storage deal with SP
	if ir.Root.String() != "" {
		defaultWallet, err := FileApi.WalletDefaultAddress(context.Background())
		if err != nil {
			log.Println("get default wallet err :\n", err)
			return "", nil, 0, address.Address{}, err
		}
		log.Println("default wallet address is: \n", defaultWallet)

		minBlocksDuration := constants.Min_Deal_Duration

		param := &lotusapi.StartDealParams{
			Data: &storagemarket.DataRef{
				TransferType: "graphsync",
				Root:         ir.Root,
			},
			Wallet:            defaultWallet,
			Miner:             cheapestMiner,
			EpochPrice:        lowestBid,
			MinBlocksDuration: uint64(minBlocksDuration),
			VerifiedDeal:      false,
		}

		dealCID, err := FileApi.ClientStartDeal(context.Background(), param)
		if err != nil {
			log.Println("make storage deal err: \n", err)
			return "", nil, 0, address.Address{}, err
		}
		log.Println("deal cid: \n", dealCID.String())

		return dealCID.String(), lowestBid.Int, uint64(minBlocksDuration), cheapestMiner, nil
	}

	return "", nil, 0, address.Address{}, nil
}

func string2Address(ID string) (*address.Address, error) {
	addr, err := address.NewFromString(ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &addr, nil
}
