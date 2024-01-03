package loc_parser

import (
	"bytes"
	"encoding/binary"
	"strings"
)

func ReadAsUint32(r *bytes.Reader) (uint32, error) {
	var b = make([]byte, 4)
	_, err := r.Read(b)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(b), nil
}

func ReadAsUint16(r *bytes.Reader) (uint16, error) {
	var b = make([]byte, 2)
	_, err := r.Read(b)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(b), nil
}

func ReadAsString(r *bytes.Reader) (string, error) {

	length, err := binary.ReadUvarint(r)
	if err != nil {
		return "", err
	}

	var val = make([]byte, length)
	if _, err := r.Read(val); err != nil {
		return "", err
	}

	res := strings.ReplaceAll(string(val), string([]byte{0x0D, 0x0A}), "\n")
	return res, nil
}
