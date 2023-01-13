package alert

import (
	"time"

	"github.com/google/uuid"
	"github.com/MichaelGenchev/NIDS/parser"
	"github.com/MichaelGenchev/NIDS/sbd"
)


type Alerter struct {
	// storage interface{}

}

//TODO CHANGE SIGNATURE TYPE
func (a *Alerter) GenerateAlert(packet *parser.PacketInfo, signature sbd.Signature) (*Alert, error){
	requestID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	alert := Alert{
		ID: int(requestID.ID()),
		Timestamp: time.Now().Format(time.RFC3339),
		Protocol: packet.Protocol,
		Signature: signature,
		Severity: signature.Severity,
		SourceIP: packet.SrcIP,
		DestinationIP: packet.DstIP,
		SourcePort: packet.SrcIP,
		DestinationPort: packet.DstIP,

	}
	return &alert, nil
}

func (a *Alerter) StoreAlert(alert *Alert) error{
	panic("lf")
}


