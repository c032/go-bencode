package bencode_test

import (
	"testing"

	"github.com/c032/go-bencode"
)

func TestInteger_Bencode(t *testing.T) {
	testCases := []struct {
		Value           bencode.Integer
		ExpectedBencode []byte
	}{
		{
			Value:           bencode.Integer(0),
			ExpectedBencode: []byte("i0e"),
		},
		{
			Value:           bencode.Integer(1),
			ExpectedBencode: []byte("i1e"),
		},
		{
			Value:           bencode.Integer(-1),
			ExpectedBencode: []byte("i-1e"),
		},
	}

	for i, tc := range testCases {
		if got, expected := string(tc.Value.Bencode()), string(tc.ExpectedBencode); got != expected {
			t.Errorf("testCases[%d].Value.Bencode() = %#v; expected %#v", i, got, expected)
		}
	}
}
