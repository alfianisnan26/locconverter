package model

import (
	"bytes"
	"fmt"

	"locconverter/internal/pkg/loc_parser"
)

type Number struct {
	value         uint32
	originalValue uint32

	Container
}

func (number *Number) Encode() ([]byte, error) {
	var buf bytes.Buffer

	if err := loc_parser.WriteFromString(number.Container.String(), &buf); err != nil {
		return nil, err
	}

	if err := loc_parser.WriteFromUint32(number.value, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (number *Number) Decode(r *bytes.Reader) error {
	key, err := loc_parser.ReadAsString(r)
	if err != nil {
		return err
	}

	if err := number.Parse(key); err != nil {
		return err
	}

	val, err := loc_parser.ReadAsUint32(r)
	if err != nil {
		return err
	}

	number.originalValue = val
	number.value = val
	return nil
}

func (number *Number) Value() uint32 {
	return number.value
}

func (number *Number) SetValue(value uint32) {
	number.value = value
	number.isModified = number.value != number.originalValue
}

func (number *Number) OriginalValue() uint32 {
	return number.originalValue
}

func (number *Number) String() string {
	return fmt.Sprintf("%s {%d}", number.Container.String(), number.value)
}
