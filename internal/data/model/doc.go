package model

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"

	"locconverter/internal/pkg/loc_parser"
)

const (
	sheetNameHeader  = "header"
	sheetNameLetters = "letters"
	sheetNameNumbers = "numbers"
)

type Document struct {
	Header  Header
	Letters []Letter
	Numbers []Number
}

func (doc *Document) EncodeLoc() (*bytes.Buffer, error) {
	var buff bytes.Buffer

	// Write Header
	b, err := doc.Header.Encode()
	if err != nil {
		return nil, err
	}
	if _, err := buff.Write(b); err != nil {
		return nil, err
	}

	// Write Length Letters
	if err := loc_parser.WriteFromUint32(uint32(len(doc.Letters)), &buff); err != nil {
		return nil, err
	}

	// Write Letters
	for _, letter := range doc.Letters {
		b, err := letter.Encode()
		if err != nil {
			return nil, err
		}
		if _, err := buff.Write(b); err != nil {
			return nil, err
		}
	}

	// Write Length Numbers
	if err := loc_parser.WriteFromUint32(uint32(len(doc.Numbers)), &buff); err != nil {
		return nil, err
	}

	// Write Number
	for _, number := range doc.Numbers {
		b, err := number.Encode()
		if err != nil {
			return nil, err
		}
		if _, err := buff.Write(b); err != nil {
			return nil, err
		}
	}

	return &buff, nil
}

func (doc *Document) DecodeLoc(r *bytes.Reader) error {
	if err := doc.Header.Decode(r); err != nil {
		return err
	}

	letterLength, err := loc_parser.ReadAsUint32(r)
	if err != nil {
		return err
	}

	doc.Letters = make([]Letter, letterLength)
	for i := uint32(0); i < letterLength; i++ {
		if err := doc.Letters[i].Decode(r); err != nil {
			return err
		}
	}

	numberLength, err := loc_parser.ReadAsUint32(r)
	if err != nil {
		return err
	}

	doc.Numbers = make([]Number, numberLength)
	for i := uint32(0); i < numberLength; i++ {
		if err := doc.Numbers[i].Decode(r); err != nil {
			return err
		}
	}

	return nil
}

func (doc *Document) EncodeExcel(f *excelize.File) error {
	// Doc -> Excel

	for _, s := range f.GetSheetList() {
		if err := f.DeleteSheet(s); err != nil {
			return err
		}
	}
	// Write Header Sheet
	if _, err := f.NewSheet(sheetNameHeader); err != nil {
		return err
	}
	var rowIndex int

	for k, v := range doc.Header.ToMap() {
		rowIndex++
		if err := writeCell(f, sheetNameHeader, rowIndex, k, v); err != nil {
			return err
		}
	}

	// Write Letters Sheet

	if _, err := f.NewSheet(sheetNameLetters); err != nil {
		return err
	}
	for i, letter := range doc.Letters {
		if err := writeCell(f, sheetNameLetters, i+1, letter.Container.String(), letter.value); err != nil {
			return err
		}
	}

	// Write Number Sheet

	if _, err := f.NewSheet(sheetNameNumbers); err != nil {
		return err
	}
	for i, number := range doc.Numbers {
		if err := writeCell(f, sheetNameNumbers, i+1, number.Container.String(), number.value); err != nil {
			return err
		}
	}

	return nil
}

func (doc *Document) DecodeExcel(f *excelize.File) error {
	// Excel -> Doc

	// Read Header Sheet
	for i := 0; ; i++ {
		k, v, err := readCell(f, sheetNameHeader, i+1)
		if err != nil {
			return err
		}

		if k == "" || v == "" {
			break
		}

		if err := doc.Header.SetMapping(k, v); err != nil {
			return err
		}
	}

	// Read Letters Sheet
	for i := 0; ; i++ {
		k, v, err := readCell(f, sheetNameLetters, i+1)
		if err != nil {
			return err
		}

		if k == "" || v == "" {
			break
		}

		var letter Letter
		letter.SetValue(v)
		letter.Container.Parse(k)

		doc.Letters = append(doc.Letters, letter)
	}

	// Read Number Sheet
	for i := 0; ; i++ {
		k, v, err := readCell(f, sheetNameNumbers, i+1)
		if err != nil {
			return err
		}

		if k == "" || v == "" {
			break
		}

		var number Number

		val, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return err
		}
		number.SetValue(uint32(val))
		number.Container.Parse(k)

		doc.Numbers = append(doc.Numbers, number)
	}

	return nil
}

func writeCell(f *excelize.File, sheet string, row int, key string, value interface{}) error {
	if err := f.SetCellValue(sheet, fmt.Sprintf("A%d", row), key); err != nil {
		return err
	}

	if err := f.SetCellValue(sheet, fmt.Sprintf("B%d", row), value); err != nil {
		return err
	}
	return nil
}

func readCell(f *excelize.File, sheet string, row int) (key string, value string, err error) {

	if key, err = f.GetCellValue(sheet, fmt.Sprintf("A%d", row)); err != nil {
		return "", "", err
	}
	if value, err = f.GetCellValue(sheet, fmt.Sprintf("B%d", row)); err != nil {
		return "", "", err
	}

	return key, value, nil
}
