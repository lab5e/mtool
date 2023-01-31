// Package util is a cornucopia of miscness all executed in glory clumsiness.
package util

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"strings"
)

// errors
var (
	ErrEmptyBuffer     = errors.New("empty buffer")
	ErrNotAMultipleOf2 = errors.New("buffer length is not a multiple of 2")
)

// BytesToInt16StringArray converts a byte array to a sensible text representation.
func BytesToInt16StringArray(data []byte, base int) ([]string, error) {
	n := len(data)

	// deal with empty buffer first
	if n == 0 {
		return nil, ErrEmptyBuffer
	}

	// make sure buffer length is multiple of 2
	if (n % 2) != 0 {
		return nil, ErrNotAMultipleOf2
	}

	elements := make([]string, n/2)
	reader := bytes.NewReader(data)
	for i := 0; ; i++ {
		var readInt16 int16
		err := binary.Read(reader, binary.BigEndian, &readInt16)
		switch err {
		// EOF means we should return the result
		case io.EOF:
			return elements, nil

		// nil means no error occurred and we add the element
		case nil:
			elements[i] = padZero(big.NewInt(int64(readInt16)).Text(base), base)
			if base == 16 {
				elements[i] = "0x" + elements[i]
			}

		// default means that err was non-nil
		default:
			return nil, err
		}
	}
}

// padZero pads binary and hexa
func padZero(s string, base int) string {
	switch base {
	// handle binary
	case 2:
		if len(s) < 16 {
			return strings.Repeat("0", 16-len(s)) + s
		}
		return s

	// handle hex
	case 16:
		if len(s) < 4 {
			return strings.Repeat("0", 4-len(s)) + s
		}
		return s

	// do nothing for other bases
	default:
		return s
	}
}

// BytesToStringArray takes a byte array and a base and returns an array of
// strings suitable for showing the user.
func BytesToStringArray(data []byte, base int) []string {
	ss := make([]string, len(data))
	for i, d := range data {
		ss[i] = big.NewInt(int64(d)).Text(base)

		switch base {
		case 2:
			if len(ss[i]) < 8 {
				ss[i] = strings.Repeat("0", 16-len(ss[i])) + ss[i]
			}

		case 16:
			if len(ss[i]) < 2 {
				ss[i] = "0" + ss[i]
			}
			ss[i] = "0x" + ss[i]
		}
	}
	return ss
}
