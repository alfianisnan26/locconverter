package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"locconverter/internal/data/model"
	"os"
)

func main() {
	path := flag.String("path", "./sample/id-ID.loc", "set path to process")
	f, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	buff, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(buff)

	var doc model.Document

	if err := doc.Decode(r); err != nil {
		panic(err)
	}

	fmt.Println(doc.Header)
	for _, letter := range doc.Letters {
		fmt.Println(letter.String())
	}

	for _, number := range doc.Numbers {
		fmt.Println(number.String())
	}

	// r := bufio.NewReader(f)
	// ver, err := r.ReadByte() // first byte is version
	// r.Discard(1)

	// lenLanguage, err := r.ReadByte()
	// language := make([]byte, lenLanguage)
	// _, err = r.Read(language)

	// lenShort, err := r.ReadByte()
	// short := make([]byte, lenShort)
	// _, err = r.Read(short)

	// lenLocale, err := r.ReadByte()
	// locale := make([]byte, lenLocale)
	// _, err = r.Read(locale)

	// unknown1 := make([]byte, 4)
	// _, err = r.Read(unknown1) // START SECTION 1 : ? 41 00 00 (3F,45,3F,3F) why english using 45?

	// lenKey, err := r.ReadByte()
	// key := make([]byte, lenKey)
	// _, err = r.Read(key)

	// lenValue, err := r.ReadByte()
	// value := make([]byte, lenValue)
	// _, err = r.Read(value)

	// unknown2 := make([]byte, 4) // START SECTION 2 : (FC-)? 00 00 00

	// pola di section 2 value berupa int32

	// > 255 character start byte with 1 left
	// 10001001 00000101 [0x89 0x05]
	// dibaca => 0001001 00000101 => dibalik =>  00000101 0001001
	// nilai: 649
}
