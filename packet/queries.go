package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

type Queries struct {
	Name  string
	QType QueryType
	Class Class
}

/*
raw: 06 67 6f 6f 67 6c 65 | 03 63 6f 6d | 00 | 00 01 | 00 01                                     ..
->:  6  g  o  o  g  l  e  | 3  c  o  m  | 0    A       IN
	 ^                      ^  		      ^
len("google")           len("com")       EOF
*/

func MarshalQueries(queries *Queries) []byte {
	var (
		raw    = make([]byte, 512)
		offset = 0
	)

	domains := strings.Split(queries.Name, ".")
	for _, domain := range domains {
		raw[offset] = byte(len(domain))
		offset += 1
		copy(raw[offset:], domain)
		offset += len(domain)
	}
	raw[offset] = 0
	offset += 1

	binary.BigEndian.PutUint16(raw[offset:], uint16(queries.QType))
	offset += 2
	binary.BigEndian.PutUint16(raw[offset:], uint16(queries.Class))
	offset += 2

	return raw[:offset]
}

func UnmarshalQueries(raw []byte, queries *Queries) error {
	if len(raw) < 6 {
		return errors.New("this is not a valid queries slice")
	}

	var (
		offset int
		domain bytes.Buffer
	)

	for raw[offset] != 0 {
		length := int(raw[offset])
		offset++
		domain.Write(raw[offset : offset+length])
		domain.WriteByte('.')
		offset += length
	}
	offset++
	queries.Name = domain.String()[:domain.Len()-1]
	queries.QType = QueryType(binary.BigEndian.Uint16(raw[offset : offset+2]))
	offset += 2
	queries.Class = Class(binary.BigEndian.Uint16(raw[offset : offset+2]))
	return nil
}
