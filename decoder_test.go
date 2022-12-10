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
