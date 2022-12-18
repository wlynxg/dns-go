package packet

type ResultCode int

const (
	NOERROR ResultCode = iota
	FORMERR
	SERVFAIL
	NXDOMAIN
	NOTIMP
	REFUSED
)

type QueryType int

const (
	UnknowQueryType QueryType = 0
	A                         = 1
	NS                        = 2
	CNAME                     = 5
	SOA                       = 6
	PTR                       = 12
	MX                        = 15
	TXT                       = 16
	AAAA                      = 28
	SRV                       = 33
	NAPTR                     = 35
	CAA                       = 257
)

type Class int

const (
	UnkonwClass Class = iota
	IN
)

const (
	FieldBuffSize  = 256
	PacketBuffSize = 1024
)
