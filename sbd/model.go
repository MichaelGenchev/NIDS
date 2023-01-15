package sbd

import (
	"strings"

	"github.com/MichaelGenchev/NIDS/parser"
)

type DetectionEvent struct {
	Signature   Signature
	ParsedPacket *parser.ParsedPacket 
}
type Signature struct {
    ID          int     // unique identifier for the signature
	Severity int
    Name        string     // name of the signature
    Description string     // description of the signature
    Type        string     // type of the signature (ban IP, detect malware, etc.)
    Keywords    []string   // keywords used for matching
    Rule        string     // rule used for matching (regular expression, CIDR notation, etc.)
    Action      string     // action to be taken when the signature is matched (block, alert, etc.)
}

func (s *Signature) Match(parsedPacket *parser.ParsedPacket) bool {
    switch s.Type {
        case "ban IP":
            return parsedPacket.SrcIP == s.Rule || parsedPacket.DstIP == s.Rule
        case "detect malware":
            for _, keyword := range s.Keywords {
                if strings.Contains(string(parsedPacket.Payload), keyword) {
                    return true
                }
            }
            return false
        // Add additional cases for other types of signatures
        default:
            return false
    }
}