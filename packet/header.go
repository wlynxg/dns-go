package packet

import (
	"encoding/binary"
	"errors"
)

type Header struct {
	TransactionID int // 16bits
	Flags         int // 16bits
	Questions     int // 16bits
	AnswersRRs    int // 16bits
	AuthorityRRs  int // 16bits
	AdditionalRRs int // 16bits
}

func MarshalHeader(header *Header) []byte {
	raw := make([]byte, 12)

	binary.BigEndian.PutUint16(raw, uint16(header.TransactionID))
	binary.BigEndian.PutUint16(raw[2:], uint16(header.Flags))
	binary.BigEndian.PutUint16(raw[4:], uint16(header.Questions))
	binary.BigEndian.PutUint16(raw[6:], uint16(header.AnswersRRs))
	binary.BigEndian.PutUint16(raw[8:], uint16(header.AuthorityRRs))
	binary.BigEndian.PutUint16(raw[10:], uint16(header.AdditionalRRs))
	return raw
}

func UnmarshalHeader(raw []byte, header *Header) (int, error) {
	if len(raw) < 12 {
		return -1, errors.New("this is not a complete DNS packet header")
	}

	header.TransactionID = int(binary.BigEndian.Uint16(raw[0:2]))
	header.Flags = int(binary.BigEndian.Uint16(raw[2:4]))
	header.Questions = int(binary.BigEndian.Uint16(raw[4:6]))
	header.AnswersRRs = int(binary.BigEndian.Uint16(raw[6:8]))
	header.AuthorityRRs = int(binary.BigEndian.Uint16(raw[8:10]))
	header.AdditionalRRs = int(binary.BigEndian.Uint16(raw[10:12]))
	return 12, nil
}
