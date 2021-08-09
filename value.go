package bencode

var (
	_ Value = (*Integer)(nil)
	_ Value = (*String)(nil)
	_ Value = (*List)(nil)
	_ Value = (*Dictionary)(nil)
)

type Value interface {
	Bencode() []byte
}
