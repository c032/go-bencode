package bencode_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/c032/go-bencode"
)

func TestDecoder_Token(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("d1:nli255ei-1ei0eee"))

	d := bencode.NewDecoder(rawData)

	var (
		err   error
		token bencode.Token
	)

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenDictionaryStart); ok {
		if got, want := parsedToken.Offset(), int64(0); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 1; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "d"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenDictionaryStart", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenString); ok {
		if got, want := parsedToken.Offset(), int64(1); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 3; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "1:n"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Value), "n"; got != want {
			t.Errorf("string(parsedToken.Value)) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenString", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenListStart); ok {
		if got, want := parsedToken.Offset(), int64(4); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 1; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "l"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenListStart", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenInteger); ok {
		if got, want := parsedToken.Offset(), int64(5); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 5; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "i255e"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := parsedToken.Value, int64(255); got != want {
			t.Errorf("string(parsedToken.Value)) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenInteger", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenInteger); ok {
		if got, want := parsedToken.Offset(), int64(10); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 4; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "i-1e"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := parsedToken.Value, int64(-1); got != want {
			t.Errorf("string(parsedToken.Value)) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenInteger", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenInteger); ok {
		if got, want := parsedToken.Offset(), int64(14); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 3; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "i0e"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := parsedToken.Value, int64(0); got != want {
			t.Errorf("string(parsedToken.Value)) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenInteger", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenEnd); ok {
		if got, want := parsedToken.Offset(), int64(17); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 1; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "e"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenEnd", token)
	}

	token, err = d.Token()
	if err != nil {
		t.Fatalf("unexpected err: %#v", err)
	}

	if parsedToken, ok := token.(*bencode.TokenEnd); ok {
		if got, want := parsedToken.Offset(), int64(18); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 1; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "e"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenEnd", token)
	}

	token, err = d.Token()
	if token != nil {
		t.Errorf("d.Token = %#v; want nil", token)
	}
	if err != io.EOF {
		t.Errorf("expected err = io.EOF; got %#v", err)
	}

	token, err = d.Token()
	if token != nil {
		t.Errorf("d.Token = %#v; want nil", token)
	}
	if err != io.EOF {
		t.Errorf("expected err = io.EOF; got %#v", err)
	}
}

func TestDecoder_Token_NegativeZero(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("i-0e"))

	d := bencode.NewDecoder(rawData)

	var (
		err   error
		token bencode.Token
	)

	token, err = d.Token()
	if err == nil {
		t.Error("unexpected nil err")
	} else {
		if parsedError, ok := err.(*bencode.ErrInvalidToken); ok {
			if got, want := parsedError.Offset, int64(0); got != want {
				t.Errorf("err.(*bencode.ErrInvalidToken).Offset = %#v; want %#v", got, want)
			}
		} else {
			t.Errorf("unexpected error %#v; want `*bencode.ErrInvalidToken`", err)
		}
	}
	if token != nil {
		t.Error("unexpected token; expected nil")
	}
}

func TestDecoder_Token_Zero(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("i0e"))

	d := bencode.NewDecoder(rawData)

	var (
		err   error
		token bencode.Token
	)

	token, err = d.Token()
	if err != nil {
		t.Fatal(err)
	}
	if parsedToken, ok := token.(*bencode.TokenInteger); ok {
		if got, want := parsedToken.Offset(), int64(0); got != want {
			t.Errorf("parsedToken.Offset() = %#v; want %#v", got, want)
		}
		if got, want := len(parsedToken.Raw()), 3; got != want {
			t.Errorf("len(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := string(parsedToken.Raw()), "i0e"; got != want {
			t.Errorf("string(parsedToken.Raw()) = %#v; want %#v", got, want)
		}
		if got, want := parsedToken.Value, int64(0); got != want {
			t.Errorf("string(parsedToken.Value)) = %#v; want %#v", got, want)
		}
	} else {
		t.Fatalf("unexpected token %#v; want bencode.TokenInteger", token)
	}
}

func TestDecoder_Decode_IntegerZero(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("i0e"))

	d := bencode.NewDecoder(rawData)

	value, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value.(int64), int64(0); got != want {
		t.Errorf("value = %#v; want %#v", got, want)
	}
}

func TestDecoder_Decode_IntegerOne(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("i1e"))

	d := bencode.NewDecoder(rawData)

	value, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value.(int64), int64(1); got != want {
		t.Errorf("value = %#v; want %#v", got, want)
	}
}

func TestDecoder_Decode_IntegerMinusOne(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("i-1e"))

	d := bencode.NewDecoder(rawData)

	value, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value.(int64), int64(-1); got != want {
		t.Errorf("value = %#v; want %#v", got, want)
	}
}

func TestDecoder_Decode_Integer256(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("i256e"))

	d := bencode.NewDecoder(rawData)

	value, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value.(int64), int64(256); got != want {
		t.Errorf("value = %#v; want %#v", got, want)
	}
}

func TestDecoder_Decode_StringEmpty(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("0:"))

	d := bencode.NewDecoder(rawData)

	value, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(value.([]byte)), ""; got != want {
		t.Errorf("string(value)) = %#v; want %#v", got, want)
	}
}

func TestDecoder_Decode_String(t *testing.T) {
	rawData := bytes.NewBuffer([]byte("1:a"))

	d := bencode.NewDecoder(rawData)

	value, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(value.([]byte)), "a"; got != want {
		t.Errorf("string(value)) = %#v; want %#v", got, want)
	}
}

func TestDecoder_Decode_Bytes(t *testing.T) {
	tests := map[string][]byte{
		"1:\x00":         []byte{'\x00'},
		"3:\x00\x01\x02": []byte{'\x00', '\x01', '\x02'},
	}

	for input, wantOutput := range tests {
		rawData := bytes.NewBuffer([]byte(input))

		d := bencode.NewDecoder(rawData)

		rawValue, err := d.Decode()
		if err != nil {
			t.Fatal(err)
		}

		value := rawValue.([]byte)

		if got, want := len(value), len(wantOutput); got != want {
			t.Errorf("len(Decode(%#v)) = %d; want %d", input, got, want)

			continue
		}

		for i := 0; i < len(wantOutput); i++ {
			if got, want := wantOutput[i], value[i]; got != want {
				t.Errorf("Decode(%#v)[%d] = %#v; want %#v", input, i, got, want)

				break
			}
		}
	}
}

func TestDecoder_Decode_ListSimple(t *testing.T) {
	tests := map[string][]interface{}{
		"le":         []interface{}{},
		"li0ee":      []interface{}{int64(0)},
		"li1ee":      []interface{}{int64(1)},
		"li1ei255ee": []interface{}{int64(1), int64(255)},
	}

	for input, wantSlice := range tests {
		rawData := bytes.NewBuffer([]byte(input))

		d := bencode.NewDecoder(rawData)

		rawValue, err := d.Decode()
		if err != nil {
			t.Fatal(err)
		}

		gotSlice := rawValue.([]interface{})

		if got, want := len(gotSlice), len(wantSlice); got != want {
			t.Errorf("len(Decode(%#v)) = %d; want %d", input, got, want)

			continue
		}

		for i := 0; i < len(wantSlice); i++ {
			if got, want := gotSlice[i], wantSlice[i]; got != want {
				t.Errorf("Decode(%#v)[%d] = %#v; want %#v", input, i, got, want)
			}
		}
	}
}

func keyDiff(left map[string]interface{}, right map[string]interface{}) (onlyLeft []string, onlyRight []string) {
	for key, _ := range left {
		if _, ok := right[key]; !ok {
			onlyLeft = append(onlyLeft, key)
		}
	}

	for key, _ := range right {
		if _, ok := left[key]; !ok {
			onlyRight = append(onlyRight, key)
		}
	}

	return
}

func TestDecoder_Decode_DictionarySimple(t *testing.T) {
	tests := map[string]map[string]interface{}{
		"de": map[string]interface{}{},
		"d0:i0ee": map[string]interface{}{
			"": int64(0),
		},
		"d1:\x00i0ee": map[string]interface{}{
			"\x00": int64(0),
		},
		"d3:onei1ee": map[string]interface{}{
			"one": int64(1),
		},
	}

	for input, wantMap := range tests {
		rawData := bytes.NewBuffer([]byte(input))

		d := bencode.NewDecoder(rawData)

		rawValue, err := d.Decode()
		if err != nil {
			t.Fatal(err)
		}

		gotMap := rawValue.(map[string]interface{})

		onlyLeft, onlyRight := keyDiff(gotMap, wantMap)
		if got, want := len(onlyLeft), 0; got != want {
			t.Errorf("Decode(%#v) has unexpected keys: %#v", input, onlyLeft)
		}

		if got, want := len(onlyRight), 0; got != want {
			t.Errorf("Decode(%#v) has missing keys: %#v", input, onlyRight)
		}

		for key, _ := range wantMap {
			if got, want := gotMap[key], wantMap[key]; got != want {
				t.Errorf("Decode(%#v)[%#v] = %#v; want %#v", input, key, got, want)
			}
		}
	}
}
func TestDecoder_Decode_DictionaryNested(t *testing.T) {
	// { "n": [255, -1, 0] }
	input := "d1:nli255ei-1ei0eee"
	rawData := bytes.NewBuffer([]byte(input))

	d := bencode.NewDecoder(rawData)

	rawValue, err := d.Decode()
	if err != nil {
		t.Fatal(err)
	}

	gotMap := rawValue.(map[string]interface{})
	wantMap := map[string]interface{}{
		"n": []interface{}{
			int64(255),
			int64(-1),
			int64(0),
		},
	}

	onlyLeft, onlyRight := keyDiff(gotMap, wantMap)
	if got, want := len(onlyLeft), 0; got != want {
		t.Errorf("Decode(%#v) has unexpected keys: %#v", input, onlyLeft)
	}

	if got, want := len(onlyRight), 0; got != want {
		t.Errorf("Decode(%#v) has missing keys: %#v", input, onlyRight)
	}

	const key = "n"
	gotList := gotMap[key].([]interface{})
	wantList := wantMap[key].([]interface{})
	for i, _ := range wantList {
		if got, want := gotList[i], wantList[i]; got != want {
			t.Errorf("Decode(%#v)[%#v][%d] = %#v; want %#v", input, key, i, got, want)
		}
	}
}
