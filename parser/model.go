package parser




type Packet struct{
	SourceMac string
	DestinationMac string
	SourceIP string
	DestinationIP string
	SourcePort string
	DestinationPort string
}

// func (p Packet) Print(){
// 	fmt.Printf("Source MAC: %s\n", p.SourceMac)
// 	fmt.Printf("Destination MAC: %s\n", p.DestinationMac)
// 	fmt.Printf("Source IP: %s\n", p.SourceIP)
// 	fmt.Printf("Destination IP: %s\n", p.DestinationIP)
// 	fmt.Printf("Source Port: %s\n", p.SourcePort)
// 	fmt.Printf("Destination Port: %s\n", p.DestinationPort)
// 	fmt.Println()
// }