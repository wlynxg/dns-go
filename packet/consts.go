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
	UnknowQueryType QueryType = iota
	A
)

type Class int

const (
	UnkonwClass Class = iota
	IN
)
