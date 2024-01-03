package model

import (
	"bytes"
	"locconverter/internal/pkg/reader"
)

type Document struct {
	Header  Header
	Letters []Letter
	Numbers []Number
}

// TODO
func (doc *Document) Encode() []byte {
	return nil
}

func (doc *Document) Decode(r *bytes.Reader) error {
	if err := doc.Header.Decode(r); err != nil {
		return err
	}

	letterLength, err := reader.ReadAsInt32(r)
	if err != nil {
		return err
	}

	doc.Letters = make([]Letter, letterLength)
	for i := int32(0); i < letterLength; i++ {
		if err := doc.Letters[i].Decode(r); err != nil {
			return err
		}
	}

	numberLength, err := reader.ReadAsInt32(r)
	if err != nil {
		return err
	}

	doc.Numbers = make([]Number, numberLength)
	for i := int32(0); i < numberLength; i++ {
		if err := doc.Numbers[i].Decode(r); err != nil {
			return err
		}
	}

	return nil
}
