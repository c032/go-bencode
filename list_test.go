package bencode_test

import (
	"testing"

	"github.com/c032/go-bencode"
)

func TestList_Bencode(t *testing.T) {
	testCases := []struct {
		Value           bencode.List
		ExpectedBencode []byte
	}{
		{
			Value: bencode.List{
				bencode.String("spam"),
				bencode.Integer(42),
			},
			ExpectedBencode: []byte("l4:spami42ee"),
		},
	}

	for i, tc := range testCases {
		if got, expected := tc.Value.Bencode(), tc.ExpectedBencode; !sameByteSlice(got, expected) {
			t.Errorf("testCases[%d].Value.Bencode() = %#v; expected %#v", i, got, expected)
		}
	}
}
