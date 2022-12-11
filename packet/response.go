package packet

type Response struct {
	Header  Header
	Queries Queries
	Answers Answers
}

func MarshalResponse(res Response) []byte {
	return nil
}
