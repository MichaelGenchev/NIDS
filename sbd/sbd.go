package sbd

import (
	"fmt"
	"log"

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
		fmt.Println("IN SBD")
		res, signature, err := sbd.CheckParsedPacket(packet)
		if err != nil {
			log.Println(err.Error())
			continue
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

func (sbd *SignatureBasedDetection) CheckParsedPacket(packet *parser.ParsedPacket) (bool, *Signature, error) {
	signatues, err := sbd.storage.FindAll()
	if err != nil {
		return false, nil, err
	}
	for _, signature := range signatues {
		res := signature.Match(packet)
		if res {
			return true, signature, nil
		}
	}
	return false, nil, nil
}
