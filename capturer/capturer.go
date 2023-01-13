package capturer


import (
	"fmt"
	"time"

	"github.com/MichaelGenchev/NIDS/parser"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Capture() {
	// Find all available network interfaces
	interfaces, _ := pcap.FindAllDevs()

	// Print the name of each interface
	for _, intf := range interfaces {
		fmt.Println(intf.Name)
	}

	// Open a handle to the first interface
	handle, _ := pcap.OpenLive("en0", 65536, true, pcap.BlockForever)
	defer handle.Close()

	// Capture packets until the user stops the program
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parsedPacked := parser.ParsePacket(packet)
		fmt.Println(parsedPacked)
		// fmt.Println(packet)
		time.Sleep(10 *time.Second)
	}
}