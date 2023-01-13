package parser




type PacketInfo struct {
    Timestamp    string
    SrcIP        string
    DstIP        string
    SrcPort      string
    DstPort      string
    Protocol     string
    Payload      []byte
    PayloadLenght int
}

