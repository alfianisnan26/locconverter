package reader

import (
	"bytes"
	"io"
	"strings"
)

func ReadAsInt32(r *bytes.Reader) (int32, error) {
	var val int32
	for i := 0; i < 4; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		var bInt int32 = int32(b)
		for j := 0; j < i; j++ {
			bInt <<= 8
		}

		val += bInt
	}
	return val, nil
}

func ReadAsString(r *bytes.Reader) (string, error) {
	var len int

	for i := 0; i < 4; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return "", err
		}

		_b := b & 0x7f
		v := int(_b)
		for j := 0; j < i; j++ {
			v <<= 7
		}

		len += v
		if b == _b {
			break
		}
	}

	// fmt.Println(len)
	var val = make([]byte, 0, len)
	for i := 0; i < len; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return "", err
		}

		if b == 0x00 {
			return "", io.EOF
		}

		val = append(val, b)
	}

	res := strings.ReplaceAll(string(val), string([]byte{0x0D, 0x0A}), "\n")
	return res, nil
}
