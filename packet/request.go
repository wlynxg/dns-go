package packet

import (
	"math/rand"
)

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

	queries := MarshalQueries(&Queries{
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
