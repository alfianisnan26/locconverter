package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"

	"locconverter/internal/service"
)

func main() {
	locFile := flag.String("loc", "./sample/id-ID.loc", "set loc file")
	xlsxFile := flag.String("xlsx", "./sample/id-ID.xlsx", "set xlsx file")

	floc, err := os.OpenFile(*locFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := floc.Close(); err != nil {
			panic(err)
		}
	}()

	fxlsx, err := excelize.OpenFile(*xlsxFile)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Mode: Loc --> Excel")

		fxlsx = excelize.NewFile()
		defer func() {
			if err := fxlsx.Close(); err != nil {
				panic(err)
			}
		}()

		if err := service.LocToExcel(floc, fxlsx); err != nil {
			panic(err)
		}

		if err := fxlsx.SaveAs(*xlsxFile); err != nil {
			panic(err)
		}

		return
	} else if err != nil {
		panic(err)
	}

	defer func() {
		if err := fxlsx.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Mode: Excel --> Loc")

	if err := service.ExcelToLoc(fxlsx, floc); err != nil {
		panic(err)
	}
}
