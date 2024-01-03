package model

import (
	"bytes"
	"fmt"
	"strconv"

	"locconverter/internal/pkg/loc_parser"
)

type Header struct {
	FileVersion uint16
	Language    string
	Short       string
	Locale      string
}

func (header *Header) Encode() ([]byte, error) {
	var buf bytes.Buffer

	if err := loc_parser.WriteFromUint16(header.FileVersion, &buf); err != nil {
		return nil, err
	}

	if err := loc_parser.WriteFromString(header.Language, &buf); err != nil {
		return nil, err
	}

	if err := loc_parser.WriteFromString(header.Short, &buf); err != nil {
		return nil, err
	}

	if err := loc_parser.WriteFromString(header.Locale, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (header *Header) Decode(r *bytes.Reader) error {
	var (
		err error
	)

	if header.FileVersion, err = loc_parser.ReadAsUint16(r); err != nil {
		return err
	}

	if header.Language, err = loc_parser.ReadAsString(r); err != nil {
		return err
	}

	if header.Short, err = loc_parser.ReadAsString(r); err != nil {
		return err
	}

	if header.Locale, err = loc_parser.ReadAsString(r); err != nil {
		return err
	}

	return nil
}

func (header *Header) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"File Version": header.FileVersion,
		"Language":     header.Language,
		"Short":        header.Short,
		"Locale":       header.Locale,
	}
}

func (header *Header) SetMapping(k string, v string) error {
	switch k {
	case "File Version":
		val, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		header.FileVersion = uint16(val)
	case "Language":
		header.Language = v
	case "Short":
		header.Short = v
	case "Locale":
		header.Locale = v
	default:
		return fmt.Errorf("unknown mapping %s", k)
	}

	return nil
}
