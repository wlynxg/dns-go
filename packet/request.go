package packet

import (
	"math/rand"
	"unsafe"
)

type DNSPacketHeader struct {
	TransactionID int // 16bits
	Flags         int // 16bits
	Questions     int // 16bits
	AnswersRRs    int // 16bits
	AuthorityRRs  int // 16bits
	AdditionalRRs int // 16bits
}

type Queries struct {
	Name  string
	QType QueryType
	Class Class
}

func generateHeader(header DNSPacketHeader) [12]byte {
	var raw [12]byte

	*(*int16)(unsafe.Pointer(&raw[0])) = int16(header.TransactionID)
	*(*int16)(unsafe.Pointer(&raw[2])) = int16(header.Flags)
	*(*int16)(unsafe.Pointer(&raw[4])) = int16(header.Questions)
	*(*int16)(unsafe.Pointer(&raw[6])) = int16(header.AnswersRRs)
	*(*int16)(unsafe.Pointer(&raw[8])) = int16(header.AuthorityRRs)
	*(*int16)(unsafe.Pointer(&raw[10])) = int16(header.AdditionalRRs)
	return raw
}

func generateQueries(queries Queries) []byte {
	var raw [512]byte

	copy(raw[:], queries.Name)
	offset := len(queries.Name)

	*(*int16)(unsafe.Pointer(&raw[offset])) = int16(queries.QType)
	offset += 2
	*(*int16)(unsafe.Pointer(&raw[offset])) = int16(queries.Class)
	offset += 2

	result := make([]byte, offset)
	copy(result, raw[:offset])
	return result
}

func NewRequest(name string) []byte {
	var (
		request [512]byte
		offset  int
	)

	header := generateHeader(DNSPacketHeader{
		TransactionID: rand.Intn(1 << 16),
		Flags:         0x0100,
		Questions:     1,
		AnswersRRs:    0,
		AuthorityRRs:  0,
		AdditionalRRs: 1,
	})
	copy(request[:], header[:])
	offset += len(header)

	queries := generateQueries(Queries{
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
