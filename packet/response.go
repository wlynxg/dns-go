package packet

type Response struct {
	Header  Header
	Queries Queries
	Answers Answers
}

func MarshalResponse(res Response) []byte {
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
