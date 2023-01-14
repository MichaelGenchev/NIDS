package alert


type AlertStorage interface {
	Save(parsedPacket *Alert) error
	FindByID(id string) (*Alert, error)
	FindAll() ([]*Alert, error)
	DeleteByID(id string) error
	Count() (int, error)
}