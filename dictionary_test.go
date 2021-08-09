package bencode_test

import (
	"testing"

	"github.com/c032/go-bencode"
)

func makeDictionary(data map[string]bencode.Value) *bencode.Dictionary {
	d := bencode.NewDictionary()

	for key, value := range data {
		d.Set(bencode.String(key), value)
	}

	return d
}

func TestDictionary_Bencode(t *testing.T) {
	testCases := []struct {
		Value           *bencode.Dictionary
		ExpectedBencode []byte
		IsSafeString    bool
	}{
		{
			IsSafeString: true,
			Value: makeDictionary(map[string]bencode.Value{
				"foo": bencode.Integer(42),
				"bar": bencode.String("spam"),
			}),
			ExpectedBencode: []byte("d3:bar4:spam3:fooi42ee"),
		},
		{
			IsSafeString: true,
			Value: makeDictionary(map[string]bencode.Value{
				"lorem": bencode.String("ipsum"),
				"foo":   bencode.Integer(42),
				"bar":   bencode.String("spam"),
				"baz": bencode.List{
					bencode.Integer(1),
					bencode.Integer(0),
					bencode.Integer(-1),
				},
			}),
			ExpectedBencode: []byte("d3:bar4:spam3:bazli1ei0ei-1ee3:fooi42e5:lorem5:ipsume"),
		},
	}

	for i, tc := range testCases {
		if got, expected := tc.Value.Bencode(), tc.ExpectedBencode; !sameByteSlice(got, expected) {
			if tc.IsSafeString {
				t.Errorf("testCases[%d].Value.Bencode() = %#v; expected %#v", i, string(got), string(expected))
			} else {
				t.Errorf("testCases[%d].Value.Bencode() = %#v; expected %#v", i, got, expected)
			}
		}
	}
}
