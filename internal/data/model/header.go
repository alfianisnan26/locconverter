package model

import (
	"bytes"
	"locconverter/internal/pkg/reader"
)

type Header struct {
	FileVersion byte
	Language    string
	Short       string
	Locale      string
}

// TODO
func (header *Header) Encode() []byte {
	return nil
}

func (header *Header) Decode(r *bytes.Reader) error {
	var (
		err error
	)

	if header.FileVersion, err = r.ReadByte(); err != nil {
		return err
	}

	if _, err = r.ReadByte(); err != nil { // skip 0x00
		return err
	}

	if header.Language, err = reader.ReadAsString(r); err != nil {
		return err
	}

	if header.Short, err = reader.ReadAsString(r); err != nil {
		return err
	}

	if header.Locale, err = reader.ReadAsString(r); err != nil {
		return err
	}

	return nil
}
