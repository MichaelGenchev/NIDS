package sbd

import (
	"fmt"

	"github.com/MichaelGenchev/NIDS/parser"
)

type SignatureBasedDetection struct {
	storage SignatureStorage
}

func NewSBD(storage SignatureStorage) *SignatureBasedDetection {
	return &SignatureBasedDetection{storage: storage}
}
func (sbd *SignatureBasedDetection) AcceptParsedPackets(chPP chan *parser.ParsedPacket, chD chan DetectionEvent) {
	for {
		packet := <-chPP
		fmt.Println("FROM SBD   ", packet.SrcIP)
		res, signature, err := sbd.CheckParsedPacket(packet)
		if err != nil {
			fmt.Println(err)
		}
		if res {
			fmt.Println("DETECTED")
			event := DetectionEvent{
				Signature:    signature,
				ParsedPacket: packet,
			}
			chD <- event
		}

	}

}

func (sbd *SignatureBasedDetection) CheckParsedPacket(packet *parser.ParsedPacket) (bool, Signature, error) {
	signatues, err := sbd.storage.FindAll()
	if err != nil {
		return false, Signature{}, err
	}
	for _, signature := range signatues {
		res := signature.Match(packet)
		if res {
			return true, *signature, nil
		}
	}
	return false, Signature{}, nil
}
