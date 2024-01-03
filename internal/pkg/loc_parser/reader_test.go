package loc_parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	buff := []byte{0x80, 0x01}
	xbuf := "34Cxz%t}V,t2ktC}7J#v,ZKa_wA?}x!,tRmCE92cfLS#KcvR4*h0CW8V[RyWYrt3!v]3/iahQ_*:Wn{paRr+T3;CnRK_[_:p.W%g&?([,&&?%*cK?+Yq&y@6A8FCS688123456"
	buff = append(buff, []byte(xbuf)...)
	r := bytes.NewReader(buff)
	str, err := ReadAsString(r)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output:", str)

	buff = []byte{0x45, 0x41, 0, 0}
	r = bytes.NewReader(buff)
	i, err := ReadAsUint32(r)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output:", i)

	buff = []byte{0x45, 0x41, 0, 0}
	r = bytes.NewReader(buff)
	val, err := ReadAsString(r)
	fmt.Println("Output:", val)
	if err != io.EOF {
		t.Fatal(err)
	}

	buff = []byte{0x0C}
	buff = append(buff, []byte("line1")...)
	buff = append(buff, []byte{0x0D, 0x0A}...)
	buff = append(buff, []byte("line2")...)
	r = bytes.NewReader(buff)
	val, err = ReadAsString(r)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output", val)

	// buff = []byte{0xC6, 0x02}
	// r = bytes.NewReader(buff)
	// val, err = ReadAsString(r)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println("Output", val)
}

func TestWriter(t *testing.T) {
	var b []byte
	b = binary.AppendUvarint(b, 500)
	fmt.Println(b, len(b), cap(b))

	b = append(b, []byte("Hello World")...)

	r := bytes.NewReader(b)
	n, err := binary.ReadUvarint(r)
	if err != nil {
		panic(err)
	}

	rfb, _ := r.ReadByte()
	str, _ := io.ReadAll(r)

	fmt.Println("n", n)
	fmt.Printf("rfb: %c | %c%s\n", rfb, rfb, str)

	var buf bytes.Buffer

	WriteFromString("Nama:\nAlfian", &buf)

	for i, b := range buf.Bytes() {
		fmt.Printf("%d.\t%c\t%d\t%x\n", i, b, b, b)
	}

	if err := WriteFromUint32(5000, &buf); err != nil {
		return
	}

	fmt.Println(buf.Bytes())
}
