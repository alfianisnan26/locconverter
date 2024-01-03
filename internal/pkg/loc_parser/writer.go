package loc_parser

import (
	"encoding/binary"
	"io"
	"strings"
)

func WriteFromUint32(v uint32, r io.Writer) error {
	var b = make([]byte, 0, 4)
	b = binary.LittleEndian.AppendUint32(b, v)
	_, err := r.Write(b)
	return err
}

func WriteFromUint16(v uint16, r io.Writer) error {
	var b = make([]byte, 0, 2)
	b = binary.LittleEndian.AppendUint16(b, v)
	_, err := r.Write(b)
	return err
}

func WriteFromString(v string, r io.Writer) error {
	v = strings.ReplaceAll(v, "\n", string([]byte{0x0D, 0x0A}))
	var b []byte
	b = binary.AppendUvarint(b, uint64(len(v)))

	if _, err := r.Write(b); err != nil {
		return err
	}
	if _, err := r.Write([]byte(v)); err != nil {
		return err
	}

	return nil
}
