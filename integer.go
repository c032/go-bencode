package bencode

import (
	"fmt"
)

type Integer int64

func (i Integer) Bencode() []byte {
	iStr := fmt.Sprintf("i%de", int64(i))

	return []byte(iStr)
}
