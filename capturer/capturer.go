package capturer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Capturer struct{}

func (c *Capturer) Capture(chPacket chan gopacket.Packet) {
	// Find all available network interfaces
	interfaces, _ := pcap.FindAllDevs()

	// Print the name of each interface
	// for _, intf := range interfaces {
	// 	fmt.Println(intf.Name)
	// }

	// Open a handle to the first interface
	handle, _ := pcap.OpenLive(interfaces[0].Name, 65536, true, pcap.BlockForever)
	defer handle.Close()

	// Capture packets until the user stops the program
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		chPacket <- packet
	}
}
