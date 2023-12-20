package dealer

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"storage/pkg/constants"
	"storage/pkg/networks"
	web3 "storage/pkg/web3"
	"time"

	cid "github.com/ipfs/go-cid"

	datasegment "github.com/filecoin-project/go-data-segment/datasegment"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	abi "github.com/filecoin-project/go-state-types/abi"
)

func RunDealer() {
	networks.ConnectToFilecoin()

	for {
		// Should be time.Minute * 10 or any other greater duration
		deals, totalDealSize, dealStartIndex, dealEndIndex, err := web3.GetTotalDeals()
		if err != nil {
			log.Println("No Batch has reached threshold")
			log.Println(err)
		}
		log.Println("Found deals total: ", deals, totalDealSize)

		if deals != nil {
			dealCID, _ := CreateDealForBatch(deals, dealStartIndex, dealEndIndex)
			log.Println("Deal : ", dealCID)
		}

		time.Sleep(constants.Polling_Rate)
	}
}

func CreateDealForBatch(deals *[]web3.DealRequest, dealStartIndex uint64, dealEndIndex uint64) (*string, error) {
	var files [][]byte
	fmt.Print(len(*deals))
	for i := 0; i < len(*deals); i++ {
		deal := (*deals)[i]
		files = append(files, deal.PieceCid[:])
	}

	currentTimestamp := time.Now().GoString()
	batchFileName := string(currentTimestamp)
	fmt.Println(batchFileName)
	if err := AggregateFiles(batchFileName, files); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Aggregated File:", batchFileName)

	dealCID, dealCostPerEpoch, dealEpochDuration, minerAddress, err := networks.CreateDealForFile(fmt.Sprintf("%s.car", batchFileName))
	log.Println("dealCostPerEpoch: ", dealCostPerEpoch)
	log.Println("dealCID: ", dealCID)
	log.Println("dealEpochDuration: ", dealEpochDuration)

	if err != nil || dealCID == "" || dealCostPerEpoch == nil || dealEpochDuration == 0 {
		log.Println(err)
		return nil, err
	}

	txReceipt, err := web3.UpdateDeal(minerAddress.Bytes(), []byte(dealCID), minerAddress.Bytes(), big.NewInt(int64(dealStartIndex)), big.NewInt(int64(dealEndIndex)))
	if err != nil {
		fmt.Println("couldn't update on SC deal status", err)
	}
	fmt.Println(txReceipt, "txReceipt")
	return nil, nil
}

func AggregateFiles(filename string, files [][]byte) error {
	log.Println("filename: ", filename)

	pieceInfos := []abi.PieceInfo{}
	reader := []io.Reader{}

	// Add files to zip
	dealSizeBytes := uint64(0)
	for _, f := range files {
		_, pieceCid, _ := cid.CidFromBytes(f)
		fileName := hex.EncodeToString(f)

		fileData, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error getting file:", err)
			continue
		}
		reader = append(reader, fileData)

		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Error getting file information:", err)
			continue
		}
		pieceInfos = append(pieceInfos, abi.PieceInfo{
			PieceCID: pieceCid,
			Size:     abi.PaddedPieceSize(524288),
		})

		dealSizeBytes = dealSizeBytes + uint64(abi.UnpaddedPieceSize(fileInfo.Size()).Padded())
	}

	// dealSize := abi.PaddedPieceSize(dealSizeBytes)
	dealSize := abi.PaddedPieceSize(1 << 20)

	a, err := datasegment.NewAggregate(dealSize, pieceInfos)
	if err != nil {
		fmt.Println("Error creating Aggregate:", err)
		return err
	}
	objectReader, err := a.AggregateObjectReader(reader)
	if err != nil {
		fmt.Println("Error creating Aggregate object reader:", err)
		return err
	}
	commpHasher := commp.Calc{}
	_, err = io.CopyBuffer(&commpHasher, objectReader, make([]byte, commpHasher.BlockSize()*128))
	if err != nil {
		fmt.Println("Error creating CopyBuffer:", err)
		return err
	}
	batchFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating batchFile:", err)
		return err
	}
	defer batchFile.Close()
	_, err = io.CopyBuffer(batchFile, objectReader, make([]byte, commpHasher.BlockSize()*128))
	if err != nil {
		fmt.Println("Error CopyBuffer:", err)
		return err
	}
	commp, _, err := commpHasher.Digest()
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	fmt.Println("commp for file", hex.EncodeToString(commp))
	pieceCid, _ := commcid.PieceCommitmentV1ToCID(commp)
	fmt.Println("commp pieceCid for file", pieceCid)
	return nil
}
