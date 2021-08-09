package bencode

type List []Value

func (l List) Bencode() []byte {
	raw := []byte("l")

	for _, v := range l {
		raw = append(raw, v.Bencode()...)
	}

	raw = append(raw, byte('e'))

	return raw
}
