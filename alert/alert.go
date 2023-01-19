package alert

import (
	"log"
	"time"

	"github.com/MichaelGenchev/NIDS/cli"
	"github.com/MichaelGenchev/NIDS/parser"
	"github.com/MichaelGenchev/NIDS/sbd"
	"github.com/google/uuid"
)

type Alerter struct {
	storage AlertStorage
}

func NewAlerter(storage AlertStorage) *Alerter {
	return &Alerter{
		storage: storage,
	}
}

func (a *Alerter) ListenForDetectionEvents(chD chan sbd.DetectionEvent, chInfo chan cli.Info) {
	for {
		event := <-chD
		_, err := a.GenerateAlert(event.ParsedPacket, event.Signature)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		info := cli.Info{
			Captured: true,
			Parsed: true,
			SBD: true,
			Alerted: true,
			Ended: true,
		}
		chInfo <- info
	}
}

// TODO CHANGE SIGNATURE TYPE
func (a *Alerter) GenerateAlert(packet *parser.ParsedPacket, signature *sbd.Signature) (*Alert, error) {
	requestID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	alert := Alert{
		ID:              int(requestID.ID()),
		Timestamp:       time.Now().Format(time.RFC3339),
		Protocol:        packet.Protocol,
		Signature:       *signature,
		Severity:        signature.Severity,
		SourceIP:        packet.SrcIP,
		DestinationIP:   packet.DstIP,
		SourcePort:      packet.SrcPort,
		DestinationPort: packet.DstPort,
	}
	err = a.storage.Save(&alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func (a *Alerter) StoreAlert(alert *Alert) error {
	return a.storage.Save(alert)
}
