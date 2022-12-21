package packet

import (
	"bytes"
	"encoding/binary"
	"net"
	"strings"
)

type Answer interface {
	i()
}

type AnswerBase struct {
	Name         string
	QType        QueryType
	Class        Class
	TTL          int
	ResourceSize int
	Resource     string
}

func (a AnswerBase) i() {}

type AnswerA struct {
	AnswerBase
	IP net.IP
}

type AnswerMX struct {
	AnswerBase
	Preference   int
	MailExchange string
}

func MarshalAnswers(answer Answer) []byte {
	switch answer.(type) {
	case AnswerA:
		as := answer.(AnswerA)
		var (
			raw    = make([]byte, FieldBuffSize)
			offset int
		)

		domains := strings.Split(as.Name, ".")
		for _, domain := range domains {
			raw[offset] = byte(len(domain))
			offset += 1
			copy(raw[offset:], domain)
			offset += len(domain)
		}
		raw[offset] = 0
		offset += 1

		binary.BigEndian.PutUint16(raw[offset:], uint16(as.QType))
		offset += 2
		binary.BigEndian.PutUint16(raw[offset:], uint16(as.Class))
		offset += 2
		binary.BigEndian.PutUint32(raw[offset:], uint32(as.TTL))
		offset += 4
		binary.BigEndian.PutUint16(raw[offset:], uint16(as.ResourceSize))
		offset += 2

		if as.QType == A && as.ResourceSize == 4 {
			ip := net.ParseIP(as.Resource)
			copy(raw[offset:], ip.To4())
			offset += 4
		}

		return raw[:offset]
	}
	return nil
}

func UnmarshalAnswers(raw []byte, answers *AnswerBase) (int, error) {
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
