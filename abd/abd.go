package abd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MichaelGenchev/NIDS/parser"
	"github.com/e-XpertSolutions/go-iforest/iforest"
)

type ABD struct {
	storage parser.ParsedPacketStorage
}

func (abd *ABD) Start() {
	packets, err := abd.storage.FindAll()
	if err != nil {
		log.Println(err.Error())
		return
	}

	trainData := [][]float64{}
	for _, p := range packets {
		floatSrcIP, _ := strconv.ParseFloat(p.SrcIP, 64)
		floatDstIP, _ := strconv.ParseFloat(p.DstIP, 64)
		floatSrcPort, _ := strconv.ParseFloat(p.SrcPort, 64)
		floatDstPort, _ := strconv.ParseFloat(p.DstPort, 64)
		t, err := time.Parse(time.RFC3339, p.Timestamp)
		if err != nil {
			fmt.Println(err)
			return
		}
		unixTime := float64(t.Unix())

		trainData = append(trainData, []float64{floatSrcIP, floatDstIP, floatSrcPort, floatDstPort, unixTime, float64(p.PayloadLenght)})
	}

	// input parameters
	treesNumber := 100
	subsampleSize := 256
	outliersRatio := 0.01
	// routinesNumber := 10

	//model initialization
	forest := iforest.NewForest(treesNumber, subsampleSize, outliersRatio)
	forest.Train(trainData)
}
