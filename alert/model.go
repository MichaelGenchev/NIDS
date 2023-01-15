package alert

import (
	"github.com/MichaelGenchev/NIDS/sbd"
)

type Alert struct {
	ID              int
	Timestamp       string
	Severity        int
	SourceIP        string
	DestinationIP   string
	SourcePort      string
	DestinationPort string
	Protocol        string
	Signature       sbd.Signature
}