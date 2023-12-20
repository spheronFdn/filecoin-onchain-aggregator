package web3

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"storage/pkg/config"
	"storage/pkg/constants"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

var privKey string = config.PRIV_KEY
var dealClientContract = config.DEAL_CLIENT_CONTRACT
var rpc = config.CHAIN_RPC

// get user token balance from contract
func GetTotalDeals() (*[]DealRequest, uint64, uint64, uint64, error) {
	client, err := ethclient.Dial(rpc)
	EthClient = client
	if err != nil {
		log.Println(err)
		return nil, 0, 0, 0, err
	}

	contractAddress := common.HexToAddress(dealClientContract)
	instance, err := NewWeb3(contractAddress, EthClient)
	if err != nil {
		log.Println(err)
		return nil, 0, 0, 0, err
	}

	aggregator, err := instance.Aggregator(nil)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	fmt.Println(aggregator)

	totalBatches, err := instance.BatchesLength(nil)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	totalDeals, err := instance.DealsLength(nil)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	fmt.Printf("TotalBatches: %s\n", totalBatches)
	fmt.Printf("TotalDeals: %s\n", totalDeals)

	dealStartIndex := 0

	if totalBatches.Uint64() == 0 && totalDeals.Uint64() == 0 {
		return nil, 0, 0, 0, nil
	}
	if totalBatches.Uint64() == 0 && totalDeals.Uint64() > 0 {
		dealStartIndex = 0
	}
	if int64(totalBatches.Uint64()) > 0 {
		latestBatch, _ := instance.Batches(&bind.CallOpts{}, big.NewInt(int64(totalBatches.Uint64())-1))
		dealStartIndex = int(latestBatch.CommPcEndIndex.Uint64())
	}

	totalSize := uint64(0)
	dealEndIndex := uint64(0)
	var deals []DealRequest

	for i := dealStartIndex; i < int(totalDeals.Int64()); i++ {
		deal, err := instance.DealRequests(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Println(err)
			return nil, 0, 0, 0, err
		}
		deals = append(deals, deal)

		fileName := hex.EncodeToString(deal.PieceCid[:])
		// Check if file already exists
		if _, err := os.Stat(fileName); err == nil {
			fmt.Printf("File '%s' already exists. Skipping download.\n", fileName)
			// Get file size after download
			fileInfo, err := os.Stat(fileName)
			if err != nil {
				fmt.Println("Error getting file information:", err)
				continue
			}
			totalSize = totalSize + uint64(fileInfo.Size())
			continue
		}

		out, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer out.Close()

		resp, err := http.Get(deal.FetchLink)
		if err != nil {
			fmt.Println("Error downloading file:", err)
			continue
		}
		defer resp.Body.Close()

		n, err := io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("Error copying content to file:", err)
			continue
		}

		fmt.Printf("Downloaded %d bytes and saved as '%s'\n", n, fileName)

		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Error getting file information:", err)
			continue
		}
		totalSize = totalSize + uint64(fileInfo.Size())

		if totalSize >= uint64(constants.Max_Batch_Size) {
			dealEndIndex = uint64(i)
			break
		}
	}
	return &deals, totalSize, uint64(dealStartIndex), dealEndIndex, nil
}

func UpdateDeal(_marketActorId []byte, _dealId []byte, _providerId []byte, commPcStartIndex *big.Int, commPcEndIndex *big.Int) (*types.Receipt, error) {
	client, err := ethclient.Dial(rpc)
	EthClient = client
	auth := GetUserFromPrivKey()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	contractAddress := common.HexToAddress(dealClientContract)
	instance, err := NewWeb3(contractAddress, client)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tx, err := instance.Complete(auth, _marketActorId, _dealId, _providerId, commPcStartIndex, commPcEndIndex)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	receipt, err := bind.WaitMined(context.Background(), EthClient, tx)
	log.Println("receipt: ", receipt)
	log.Println("receiptStatus: ", receipt.Status)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}

	return receipt, err
}

func GetUserFromPrivKey() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(4))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	return auth
}
