package packet

type Response struct {
	Header  Header
	Queries Queries
	Answers Answers
}

func MarshalResponse(res Response) []byte {
	return nil
}

func UnmarshalResponse(raw []byte, res *Response) (int, error) {
	return -1, nil
}
