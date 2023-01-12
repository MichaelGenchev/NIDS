package parser

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)



func Parse(packet gopacket.Packet) (Packet, error ){

	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer == nil {
		return Packet{}, fmt.Errorf("nil ethernet layer")
	}
	ethernet := ethernetLayer.(*layers.Ethernet)

	// Extract the IP layer 
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return Packet{}, fmt.Errorf("nil ip layer")
	}
	ip := ipLayer.(*layers.IPv4)

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return Packet{}, fmt.Errorf("nil tcp layer") 
	}
	tcp := tcpLayer.(*layers.TCP)

	// Print the packet information
	currPacket := Packet{
		SourceMac: ethernet.SrcMAC.String(),
		DestinationMac: ethernet.DstMAC.String(),
		SourceIP: ip.SrcIP.String(),
		DestinationIP: ip.DstIP.String(),
		SourcePort: tcp.SrcPort.String(),
		DestinationPort: tcp.DstPort.String(),
	}

	fmt.Printf("Source MAC: %s\n", ethernet.SrcMAC)
	fmt.Printf("Destination MAC: %s\n", ethernet.DstMAC)
	fmt.Printf("Source IP: %s\n", ip.SrcIP)
	fmt.Printf("Destination IP: %s\n", ip.DstIP)
	fmt.Printf("Source Port: %d\n", tcp.SrcPort)
	fmt.Printf("Destination Port: %d\n", tcp.DstPort)
	fmt.Println()
	return currPacket, nil
}