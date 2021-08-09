package bencode_test

import (
	"testing"

	"github.com/c032/go-bencode"
)

func TestString_Bencode(t *testing.T) {
	testCases := []struct {
		Value           bencode.String
		ExpectedBencode []byte
		ExpectedKey     string
	}{
		{
			Value:           bencode.String(""),
			ExpectedBencode: []byte("0:"),
			ExpectedKey:     "",
		},
		{
			Value:           bencode.String("Hello, world!"),
			ExpectedBencode: []byte("13:Hello, world!"),
			ExpectedKey:     "48656c6c6f2c20776f726c6421",
		},
	}

	for i, tc := range testCases {
		if got, expected := tc.Value.Bencode(), tc.ExpectedBencode; !sameByteSlice(got, expected) {
			t.Errorf("testCases[%d].Value.Bencode() = %#v; expected %#v", i, got, expected)
		}

		if got, expected := tc.Value.BencodeKey(), tc.ExpectedKey; got != expected {
			t.Errorf("testCases[%d].Value.BencodeKey() = %#v; expected %#v", i, got, expected)
		}
	}
}
