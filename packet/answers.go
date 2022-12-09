package packet

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Answers struct {
	Name         string
	QType        QueryType
	Class        Class
	TTL          int
	ResourceSize int
	Resource     string
}

func MarshalAnswers(ansers *Answers) []byte {
	return nil
}

func UnmarshalAnswers(raw []byte, answers *Answers) (int, error) {
	var (
		offset int
		domain bytes.Buffer
	)

	// The dns reply enables message compression
	// https://www.rfc-editor.org/rfc/rfc1035#section-4.1.4
	if raw[0] == 0xc0 && raw[1] == 0x0c {
		answers.Name = "0xc00c"
		offset += 2
	} else {
		for raw[offset] != 0 {
			length := int(raw[offset])
			offset++
			domain.Write(raw[offset : offset+length])
			domain.WriteByte('.')
			offset += length
		}
		offset++
		answers.Name = domain.String()[:domain.Len()-1]
	}

	answers.QType = QueryType(binary.BigEndian.Uint16(raw[offset : offset+2]))
	offset += 2
	answers.Class = Class(binary.BigEndian.Uint16(raw[offset : offset+2]))
	offset += 2
	answers.TTL = int(binary.BigEndian.Uint32(raw[offset : offset+4]))
	offset += 4
	answers.ResourceSize = int(binary.BigEndian.Uint16(raw[offset : offset+2]))
	offset += 2

	if answers.QType == A && answers.ResourceSize == 4 {
		ipv4 := net.IPv4(raw[offset], raw[offset+1], raw[offset+2], raw[offset+3])
		answers.Resource = ipv4.String()
	}
	offset += answers.ResourceSize
	return offset, nil
}
