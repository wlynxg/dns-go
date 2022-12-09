package packet

import (
	"encoding/binary"
	"math/rand"
	"strings"
)

type Queries struct {
	Name  string
	QType QueryType
	Class Class
}

func serializeQueries(queries Queries) []byte {
	var (
		raw    = make([]byte, 512)
		offset = 0
	)

	domains := strings.Split(queries.Name, ".")
	for _, domain := range domains {
		raw[offset] = byte(len(domain))
		offset += 1
		copy(raw[offset:], domain)
		offset += len(domain)
	}
	raw[offset] = 0
	offset += 1

	binary.BigEndian.PutUint16(raw[offset:], uint16(queries.QType))
	offset += 2
	binary.BigEndian.PutUint16(raw[offset:], uint16(queries.Class))
	offset += 2

	return raw[:offset]
}

func NewRequest(name string) []byte {
	var (
		request [512]byte
		offset  int
	)

	header := MarshalHeader(DNSPacketHeader{
		TransactionID: rand.Intn(1 << 16),
		Flags:         0x0100,
		Questions:     1,
		AnswersRRs:    0,
		AuthorityRRs:  0,
		AdditionalRRs: 0,
	})
	copy(request[:], header)
	offset += len(header)

	queries := serializeQueries(Queries{
		Name:  name,
		QType: A,
		Class: IN,
	})
	copy(request[offset:], queries)
	offset += len(queries)

	result := make([]byte, offset)
	copy(result, request[:offset])
	return result
}
