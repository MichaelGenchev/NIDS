package parser

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)


func ParsePacket(packet gopacket.Packet) *PacketInfo {
    var packetInfo PacketInfo
    // Extract packet timestamp
    packetInfo.Timestamp = packet.Metadata().Timestamp.String()

    // Extract packet IP layer
    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer != nil {
        ip := ipLayer.(*layers.IPv4)
        packetInfo.SrcIP = ip.SrcIP.String()
        packetInfo.DstIP = ip.DstIP.String()
        packetInfo.Protocol = ip.Protocol.String()
    }

    // Extract packet transport layer
    transportLayer := packet.TransportLayer()
    if transportLayer != nil {
        switch transportLayer.LayerType() {
        case layers.LayerTypeTCP:
            tcp := transportLayer.(*layers.TCP)
            packetInfo.SrcPort = fmt.Sprintf("%d", tcp.SrcPort)
            packetInfo.DstPort = fmt.Sprintf("%d", tcp.DstPort)
        case layers.LayerTypeUDP:
            udp := transportLayer.(*layers.UDP)
            packetInfo.SrcPort = fmt.Sprintf("%d", udp.SrcPort)
            packetInfo.DstPort = fmt.Sprintf("%d", udp.DstPort)
        }
    }

    // Extract packet payload
    applicationLayer := packet.ApplicationLayer()
    if applicationLayer != nil {
        packetInfo.Payload = applicationLayer.Payload()
        packetInfo.PayloadLenght = len(applicationLayer.Payload())
    }
    return &packetInfo
}