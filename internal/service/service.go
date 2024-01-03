package service

import (
	"bytes"
	"io"
	"os"

	"github.com/xuri/excelize/v2"

	"locconverter/internal/data/model"
)

func LocToExcel(s *os.File, d *excelize.File) error {
	var doc model.Document

	b, err := io.ReadAll(s)
	if err != nil {
		panic(err)
	}

	if err := doc.DecodeLoc(bytes.NewReader(b)); err != nil {
		return err
	}

	if err := doc.EncodeExcel(d); err != nil {
		return err
	}

	return nil
}

func ExcelToLoc(s *excelize.File, d *os.File) error {
	d.Truncate(0)

	var doc model.Document

	if err := doc.DecodeExcel(s); err != nil {
		return err
	}

	buf, err := doc.EncodeLoc()
	if err != nil {
		return err
	}

	if _, err := io.Copy(d, buf); err != nil {
		return err
	}

	return nil
}
