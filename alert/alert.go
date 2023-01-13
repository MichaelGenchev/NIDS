package alert

import (
	// "time"

	// "github.com/google/gopacket"
)


type Alerter struct {

}

//TODO CHANGE SIGNATURE TYPE
func (a *Alerter) GenerateAlert(packet string, signature string) error{

	// f := Alert{
	// 	ID: "testggr",
	// 	Timestamp: time.Now().Format(time.RFC3339),
	// 	Severity: 1,
	// 	SourceIP: "",
	// }
	panic("implement me")

}

func (a *Alerter) StoreAlert(alert *Alert) error{
	panic("g")
}

func (a *Alerter) SendAlert(alert *Alert)error {
	panic("implement me")
}
