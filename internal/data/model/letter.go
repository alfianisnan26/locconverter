package model

import (
	"bytes"
	"fmt"
	"locconverter/internal/pkg/reader"
)

type Letter struct {
	value         string
	originalValue string

	Container
}

// TODO
func (letter *Letter) Encode() []byte {
	return nil
}

func (letter *Letter) Decode(r *bytes.Reader) error {
	key, err := reader.ReadAsString(r)
	if err != nil {
		return err
	}

	if err := letter.Parse(key); err != nil {
		return err
	}

	val, err := reader.ReadAsString(r)
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
	letter.isModified = letter.value != letter.originalValue
}

func (letter *Letter) OriginalValue() string {
	return letter.originalValue
}

func (letter *Letter) String() string {
	return fmt.Sprintf("%s {%s}", letter.Container.String(), letter.value)
}
