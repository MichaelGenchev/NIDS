package parser




type PacketInfo struct {
	ID int
    Timestamp    string
    SrcIP        string
    DstIP        string
    SrcPort      string
    DstPort      string
    Protocol     string
    Payload      []byte
    PayloadLenght int
}

