package model

import (
	"bytes"
	"fmt"

	"locconverter/internal/pkg/loc_parser"
)

type Letter struct {
	value         string
	originalValue string

	Container
}

func (letter *Letter) Encode() ([]byte, error) {
	var buf bytes.Buffer

	if err := loc_parser.WriteFromString(letter.Container.String(), &buf); err != nil {
		return nil, err
	}

	if err := loc_parser.WriteFromString(letter.value, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (letter *Letter) Decode(r *bytes.Reader) error {
	key, err := loc_parser.ReadAsString(r)
	if err != nil {
		return err
	}

	if err := letter.Parse(key); err != nil {
		return err
	}

	val, err := loc_parser.ReadAsString(r)
	if err != nil {
		return err
	}

	letter.originalValue = val
	letter.value = val
	return nil
}

func (letter *Letter) Value() string {
	return letter.value
}

func (letter *Letter) SetValue(value string) {
	letter.value = value
	if letter.originalValue == "" {
		letter.originalValue = value
	}

	letter.isModified = letter.value != letter.originalValue
}

func (letter *Letter) OriginalValue() string {
	return letter.originalValue
}

func (letter *Letter) String() string {
	return fmt.Sprintf("%s {%s}", letter.Container.String(), letter.value)
}
