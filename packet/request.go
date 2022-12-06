package packet

import "unsafe"

type DNSPacketHeaderRaw [12]byte
type DNSPacketHeader struct {
	TransactionID int // 16bits

	RecursionDesired    bool // 1bit
	TruncatedMessage    bool // 1bit
	AuthoritativeAnswer bool // 1bit
	Opcode              int  // 4bits
	Response            bool // 1bit

	ResCode            ResultCode // 4bits
	CheckingDisabled   bool       // 1bit
	AuthedData         bool       // 1bit
	Z                  bool       // 1bit
	RecursionAvailable bool       // 1bit

	Questions     int // 16bits
	AnswersRRs    int // 16bits
	AuthorityRRs  int // 16bits
	AdditionalRRs int // 16bits
}

func GenearateHeader(header DNSPacketHeader) DNSPacketHeaderRaw {
	var raw DNSPacketHeaderRaw

	var flag uint16
	if header.RecursionDesired {
		flag += 16 << 1
	}
	if header.TruncatedMessage {
		flag += 15 << 1
	}

	*(*int16)(unsafe.Pointer(&raw[0])) = int16(header.TransactionID)
	*(*uint16)(unsafe.Pointer(&raw[2])) = uint16(int16(flag))
	*(*int16)(unsafe.Pointer(&raw[4])) = int16(header.Questions)
	*(*int16)(unsafe.Pointer(&raw[6])) = int16(header.AnswersRRs)
	*(*int16)(unsafe.Pointer(&raw[8])) = int16(header.AuthorityRRs)
	*(*int16)(unsafe.Pointer(&raw[10])) = int16(header.AdditionalRRs)
	return raw
}
