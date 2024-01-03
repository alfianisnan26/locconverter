package model

import (
	"bytes"
	"fmt"
	"locconverter/internal/pkg/reader"
)

type Number struct {
	value         int32
	originalValue int32

	Container
}

// TODO
func (number *Number) Encode() []byte {
	return nil
}

func (number *Number) Decode(r *bytes.Reader) error {
	key, err := reader.ReadAsString(r)
	if err != nil {
		return err
	}

	if err := number.Parse(key); err != nil {
		return err
	}

	val, err := reader.ReadAsInt32(r)
	if err != nil {
		return err
	}

	number.originalValue = val
	number.value = val
	return nil
}

func (number *Number) Value() int32 {
	return number.value
}

func (number *Number) SetValue(value int32) {
	number.value = value
	number.isModified = number.value != number.originalValue
}

func (number *Number) OriginalValue() int32 {
	return number.originalValue
}

func (number *Number) String() string {
	return fmt.Sprintf("%s {%d}", number.Container.String(), number.value)
}
