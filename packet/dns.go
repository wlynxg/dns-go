package packet

type DNSPacketHeaderRaw [12]byte
type DNSPacketHeader struct {
	TransactionID int
	Flags         int
	Questions     int
	AnswersRRs    int
	AuthorityRRs  int
	AdditionalRRs int
}
