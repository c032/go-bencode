package bencode

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type ErrStringTooLong struct {
	TokenOffset     int64
	NextTokenOffset int64
}

func (e *ErrStringTooLong) Error() string {
	return fmt.Sprintf("string too long at offset %d; next token starts at %d", e.TokenOffset, e.NextTokenOffset)
}

type ErrInvalidToken struct {
	Offset int64
}

func (e *ErrInvalidToken) Error() string {
	return fmt.Sprintf("unexpected byte at offset %d", e.Offset)
}

type ErrUnexpectedByte struct {
	Offset   int64
	Got      byte
	Expected byte
}

func (e *ErrUnexpectedByte) Error() string {
	return fmt.Sprintf("unexpected byte %#v at offset %d, expected %#v", e.Got, e.Offset, e.Expected)
}

var _ error = (*ErrUnexpectedByte)(nil)

var (
	_ Token = (*TokenDictionaryStart)(nil)
	_ Token = (*TokenListStart)(nil)
	_ Token = (*TokenString)(nil)
	_ Token = (*TokenInteger)(nil)
	_ Token = (*TokenEnd)(nil)
)

var _ TokenReader = (*Decoder)(nil)

// Token is an interface holding one of the token types: TokenDictionaryStart,
// TokenListStart, TokenString, TokenInteger, TokenEnd.
type Token interface {
	Offset() int64
	Raw() []byte
}

type TokenReader interface {
	Token() (Token, error)
}

type baseToken struct {
	offset int64
	raw    []byte
}

func (t *baseToken) Offset() int64 {
	return t.offset
}

func (t *baseToken) Raw() []byte {
	return t.raw
}

type TokenDictionaryStart struct {
	baseToken
}

type TokenListStart struct {
	baseToken
}

type TokenString struct {
	offset int64
	raw    []byte
	Value  []byte
}

func (ts *TokenString) setValue(value []byte) {
	prefix := []byte(fmt.Sprintf("%d", len(value)))
	parts := [][]byte{
		prefix,
		value,
	}
	ts.raw = bytes.Join(parts, []byte(":"))
	ts.Value = ts.raw[len(prefix)+1:]
}

func (ts *TokenString) Offset() int64 {
	return ts.offset
}

func (ts *TokenString) Raw() []byte {
	return ts.raw
}

type TokenInteger struct {
	baseToken
	Value int64
}

type TokenEnd struct {
	baseToken
}

var DefaultDecoderOptions = DecoderOptions{
	MaxIntegerLength: 64 * 1024,
	MaxStringLength:  16 * 1024 * 1024,
}

type DecoderOptions struct {
	// MaxIntegerLength is the maximum amount of bytes that will be used when
	// reading integer values, excluding delimiters.
	//
	// For example, `i10e` needs `MaxIntegerLength >= 2` because `10` is 2
	// bytes long.
	//
	// This value does NOT affect the parsing of strings.
	MaxIntegerLength int

	// MaxStringLength is the maximum length of the content of a string that
	// can be parsed, excluding its prefix.
	//
	// For example, `4:test` needs `MaxStringLength >= 4` becase `test` is 4
	// bytes long.
	//
	// The parsing of byte strings is NOT affected by `MaxIntegerLength`. When
	// reading a byte string's prefix, the decoder will read at most
	// `len(fmt.Sprintf("%d",MaxStringLength))` bytes for the length part, and
	// will expect a `:` byte after that.
	//
	// For example, if `MaxStringLength=9`, then `12:Lorem ipsum.` will return
	// an error because it will read at most 1 byte for the length (because
	// `len("9")==1`), but since the next byte is not a delimiter (it's a `2`
	// instead of `:`), it will stop reading right there and return an error.
	MaxStringLength int64
}

type Decoder struct {
	r io.Reader

	options DecoderOptions

	offset int64
	isEOF  bool
}

// Token decodes a new token.
//
// One of the returned values is always nil. That means it returns _either_ a
// valid token and a nil error, or a nil token and a non-nil error.
func (d *Decoder) Token() (Token, error) {
	if d.isEOF {
		return nil, io.EOF
	}

	var (
		err error
		n   int
	)

	prefix := make([]byte, 1)

	n, err = d.r.Read(prefix)
	if err == io.EOF {
		d.isEOF = true

		if n == 0 {
			return nil, io.EOF
		}

		// If something was read, continue.
		err = nil
	}
	if err != nil {
		return nil, fmt.Errorf("error reading: %#v", err)
	}

	tokenOffset := d.offset
	d.offset += int64(n)

	c := prefix[0]
	if c == 'e' {
		t := &TokenEnd{
			baseToken: baseToken{
				offset: tokenOffset,
				raw:    []byte{'e'},
			},
		}

		return t, nil
	} else if c == 'i' {
		rawNumber := []byte{}

		readBuffer := make([]byte, 1)

		lastReadByte := c
		for len(rawNumber) <= d.options.MaxIntegerLength {
			isFirstByte := len(rawNumber) == 0

			n, err = d.r.Read(readBuffer)
			d.offset += int64(n)
			if err != nil {
				if err == io.EOF {
					d.isEOF = true

					return nil, fmt.Errorf("unexpected EOF: %w", err)
				} else {
					return nil, fmt.Errorf("read error: %w", err)
				}
			}

			lastReadByte = readBuffer[0]
			if lastReadByte == 'e' {
				break
			}

			isValid := (isFirstByte && lastReadByte == '-') || (lastReadByte >= '0' && lastReadByte <= '9')
			if !isValid {
				err := &ErrInvalidToken{
					Offset: tokenOffset,
				}

				return nil, err
			}

			rawNumber = append(rawNumber, lastReadByte)
		}

		if lastReadByte != 'e' {
			err := &ErrUnexpectedByte{
				Offset:   tokenOffset,
				Got:      lastReadByte,
				Expected: 'e',
			}

			return nil, err
		}

		var parsedNumber int64

		parsedNumber, err = strconv.ParseInt(string(rawNumber), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse byte string length: %#v", err)
		}

		raw := []byte{'i'}
		raw = append(raw, []byte(fmt.Sprintf("%d", parsedNumber))...)
		raw = append(raw, 'e')

		if fmt.Sprintf("i%se", string(rawNumber)) != string(raw) {
			err := &ErrInvalidToken{
				Offset: tokenOffset,
			}

			return nil, err
		}

		t := &TokenInteger{
			baseToken: baseToken{
				offset: tokenOffset,
				raw:    raw,
			},
			Value: parsedNumber,
		}

		return t, nil
	} else if c == 'd' {
		t := &TokenDictionaryStart{
			baseToken: baseToken{
				offset: tokenOffset,
				raw:    []byte{'d'},
			},
		}

		return t, nil
	} else if c == 'l' {
		t := &TokenListStart{
			baseToken: baseToken{
				offset: tokenOffset,
				raw:    []byte{'l'},
			},
		}

		return t, nil
	} else if c >= '0' && c <= '9' {
		maxLengthBytes := len(fmt.Sprintf("%d", d.options.MaxStringLength))
		rawLength := []byte{c}

		readBuffer := make([]byte, 1)

		lastReadByte := c
		for len(rawLength) <= maxLengthBytes {
			n, err = d.r.Read(readBuffer)
			d.offset += int64(n)
			if err != nil {
				if err == io.EOF {
					d.isEOF = true

					return nil, fmt.Errorf("unexpected EOF: %w", err)
				} else {
					return nil, fmt.Errorf("read error: %w", err)
				}
			}

			lastReadByte = readBuffer[0]
			if lastReadByte == ':' {
				break
			}

			isValid := (lastReadByte >= '0' && lastReadByte <= '9')
			if !isValid {
				err := &ErrInvalidToken{
					Offset: tokenOffset,
				}

				return nil, err
			}

			rawLength = append(rawLength, lastReadByte)
		}

		if lastReadByte != ':' {
			err := &ErrUnexpectedByte{
				Offset:   tokenOffset,
				Got:      lastReadByte,
				Expected: ':',
			}

			return nil, err
		}

		var parsedLength int64

		parsedLength, err = strconv.ParseInt(string(rawLength), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse byte string length: %#v", err)
		}
		if parsedLength > d.options.MaxStringLength {
			err := &ErrStringTooLong{
				TokenOffset:     tokenOffset,
				NextTokenOffset: d.offset + parsedLength,
			}

			return nil, err
		}

		dst := &bytes.Buffer{}

		var copiedBytes int64

		copiedBytes, err = io.CopyN(dst, d.r, parsedLength)
		d.offset += copiedBytes
		if err != nil {
			return nil, fmt.Errorf("read error: %w", err)
		}
		if copiedBytes != parsedLength {
			return nil, fmt.Errorf("unexpected read of %d bytes; wanted %d bytes", copiedBytes, parsedLength)
		}

		t := &TokenString{
			offset: tokenOffset,
		}

		t.setValue(dst.Bytes())

		return t, nil
	} else {
		err := &ErrInvalidToken{
			Offset: tokenOffset,
		}

		return nil, err
	}
}

func NewDecoder(r io.Reader) *Decoder {
	return NewDecoderWithOptions(r, DefaultDecoderOptions)
}

func NewDecoderWithOptions(r io.Reader, options DecoderOptions) *Decoder {
	return &Decoder{
		r: r,

		options: options,
	}
}
