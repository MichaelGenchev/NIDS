package alert

type AlertStorage interface {
	Save(alert *Alert) error
	FindByID(id string) (*Alert, error)
	FindAll() ([]*Alert, error)
	DeleteByID(id string) error
	Count() (int, error)
}
