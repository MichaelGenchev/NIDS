package parser

type ParsedPacket struct {
	ID            int
	Timestamp     string
	SrcIP         string
	DstIP         string
	SrcPort       string
	DstPort       string
	Protocol      string
	Payload       []byte
	PayloadLenght int
}
