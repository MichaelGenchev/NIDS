package parser

type ParsedPacketStorage interface {
	Save(parsedPacket *ParsedPacket) error
	FindByID(id string) (*ParsedPacket, error)
	FindAll() ([]*ParsedPacket, error)
	DeleteByID(id string) error
	Count() (int, error)
}
