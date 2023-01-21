package abd

import (
	"github.com/MichaelGenchev/NIDS/parser"
	// "github.com/e-XpertSolutions/go-iforest/iforest"
)

type ABD struct {
	storage parser.ParsedPacketStorage
}

type PacketSet []*parser.ParsedPacket

func (abd *ABD) AcceptParsedPackets(chPP chan *parser.ParsedPacket, chTraining, chTesting, chPredicting chan PacketSet) {
	var packetSet PacketSet
	for {
		pp := <-chPP
		if len(packetSet) == 100 {
			chTraining <- packetSet[:50]
			chTesting <- packetSet[50:75]
			chPredicting <- packetSet[75:]
			packetSet = nil
			continue
		}
		packetSet = append(packetSet, pp)
	}
}

// TODO PLAN

// 1. Train the Isolation tree with data from the database (7k + packets in there)
// 2. have two subsets of one packetSets used for training and predicting.
// 3. Use waitgroup to wait for training and testing, before Predicting
// 4. Connect ABD to the rest of NIDS
