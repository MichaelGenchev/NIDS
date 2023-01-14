package parser

import (
	"errors"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/uuid"
)

var (
	ErrIPLayer          = errors.New("ipLayer is nil")
	ErrTransportLayer   = errors.New("transport layer is nil")
	ErrApplicationLayer = errors.New("application layer is nil")
)

type Parser struct {
	storage ParsedPacketStorage 
}

func NewParser(storage ParsedPacketStorage) *Parser {
	return &Parser{storage: storage}
}

func (p *Parser)ParsePacket(packet gopacket.Packet) (*ParsedPacket, error) {
	var parsedPacked ParsedPacket
	// Extract packet timestamp
	parsedPacked.Timestamp = packet.Metadata().Timestamp.String()

	// Extract packet IP layer
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return nil, ErrIPLayer
	}
	ip := ipLayer.(*layers.IPv4)
	parsedPacked.SrcIP = ip.SrcIP.String()
	parsedPacked.DstIP = ip.DstIP.String()
	parsedPacked.Protocol = ip.Protocol.String()

	// Extract packet transport layer
	transportLayer := packet.TransportLayer()
	if transportLayer == nil {
		return nil, ErrTransportLayer
	}
	switch transportLayer.LayerType() {
	case layers.LayerTypeTCP:
		tcp := transportLayer.(*layers.TCP)
		parsedPacked.SrcPort = fmt.Sprintf("%d", tcp.SrcPort)
		parsedPacked.DstPort = fmt.Sprintf("%d", tcp.DstPort)
	case layers.LayerTypeUDP:
		udp := transportLayer.(*layers.UDP)
		parsedPacked.SrcPort = fmt.Sprintf("%d", udp.SrcPort)
		parsedPacked.DstPort = fmt.Sprintf("%d", udp.DstPort)
	}

	// Extract packet payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer == nil {
		return nil, ErrApplicationLayer
	}
	parsedPacked.Payload = applicationLayer.Payload()
	parsedPacked.PayloadLenght = len(applicationLayer.Payload())
	requestID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	parsedPacked.ID = int(requestID.ID())

	// Saves the parsedPacket to database
	err = p.storage.Save(&parsedPacked)
	if err != nil {
		return nil, err
	}

	return &parsedPacked, nil
}
