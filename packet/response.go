package packet

import "dns-go/record"

type Response struct {
	Header  Header
	Queries Queries
	Answers Answers
}

func MarshalResponse(res *Response) []byte {
	raw := make([]byte, PacketBuffSize)
	offset := 0

	data := MarshalHeader(&res.Header)
	copy(raw[offset:], data)
	offset += len(data)

	data = MarshalQueries(&res.Queries)
	copy(raw[offset:], data)
	offset += len(data)

	data = MarshalAnswers(&res.Answers)
	copy(raw[offset:], data)
	offset += len(data)
	return raw[:offset]
}

func UnmarshalResponse(raw []byte, res *Response) (int, error) {
	var offset int

	n, err := UnmarshalHeader(raw[offset:], &res.Header)
	if err != nil {
		return -1, err
	}
	offset += n

	n, err = UnmarshalQueries(raw[offset:], &res.Queries)
	if err != nil {
		return -1, err
	}
	offset += n

	n, err = UnmarshalAnswers(raw[offset:], &res.Answers)
	if err != nil {
		return -1, err
	}
	offset += n
	return offset, nil
}

func NewResponse(req *Request) ([]byte, error) {
	ip, err := record.Query(req.Queries.Name, req.Queries.QType)
	if err != nil {
		return nil, err
	}

	res := Response{
		Header: Header{
			TransactionID: req.Header.TransactionID,
			Flags:         0x8180,
			Questions:     1,
			AnswersRRs:    1,
			AuthorityRRs:  0,
			AdditionalRRs: 0,
		},
		Queries: req.Queries,
		Answers: Answers{
			Name:         req.Queries.Name,
			QType:        req.Queries.QType,
			Class:        req.Queries.Class,
			TTL:          600,
			ResourceSize: 4,
			Resource:     ip.String(),
		},
	}

	return MarshalResponse(&res), nil
}
