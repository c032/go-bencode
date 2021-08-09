package bencode

import (
	"fmt"
)

type String []byte

func (s String) Bencode() []byte {
	sBytes := []byte(s)
	prefix := []byte(fmt.Sprintf("%d:", len(sBytes)))

	raw := make([]byte, 0, len(prefix)+len(sBytes))

	raw = append(raw, prefix...)
	raw = append(raw, sBytes...)

	return raw
}

func (s String) BencodeKey() string {
	return fmt.Sprintf("%x", []byte(s))
}
