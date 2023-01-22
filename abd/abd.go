package abd

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/MichaelGenchev/NIDS/parser"
	"github.com/e-XpertSolutions/go-iforest/iforest"
)

type ABD struct {
	forest  *iforest.Forest
	wg      *sync.WaitGroup
	storage parser.ParsedPacketStorage
}

type PacketSet []*parser.ParsedPacket

func (abd *ABD) AcceptParsedPackets(chPP chan *parser.ParsedPacket, chTesting, chPredicting chan PacketSet) {
	var packetSet PacketSet
	for {
		pp := <-chPP
		if len(packetSet) == 100 {
			chTesting <- packetSet[:40]
			chPredicting <- packetSet[40:]
			packetSet = nil
			continue
		}
		packetSet = append(packetSet, pp)
	}
}

func (abd *ABD) ParsePacketsToMatrix(packets []*parser.ParsedPacket) [][]float64 {
	trainData := [][]float64{}

	for _, p := range packets {
		floatSrcIP, _ := strconv.ParseFloat(p.SrcIP, 64)
		floatDstIP, _ := strconv.ParseFloat(p.DstIP, 64)
		floatSrcPort, _ := strconv.ParseFloat(p.SrcPort, 64)
		floatDstPort, _ := strconv.ParseFloat(p.DstPort, 64)
		t, err := time.Parse(time.RFC3339, p.Timestamp)
		if err != nil {
			log.Println("Error parsing time: ", err)
			return nil
		}
		unixTime := float64(t.Unix())
		trainData = append(trainData, []float64{floatSrcIP, floatDstIP, floatSrcPort, 
			floatDstPort, unixTime})
	}
	return trainData
}

func (abd *ABD) TrainForestFromMongoDB(){
	for {
		packets, err := abd.storage.FindAll()
		if err != nil {
			log.Println("Error getting new packets: ", err)
			continue
		}
		trainData := abd.ParsePacketsToMatrix(packets)

		abd.forest.Train(trainData)
	}
}

func (abd *ABD) TestForest(chTesting chan PacketSet){
	
}
// TODO PLAN

// 1. Train the Isolation tree with data from the database (7k + packets in there)
// 2. have two subsets of one packetSets used for training and predicting.
// 3. Use waitgroup to wait for training and testing, before Predicting
// 4. Connect ABD to the rest of NIDS
