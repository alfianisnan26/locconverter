package reader

import (
	"bytes"
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
	i, err := ReadAsInt32(r)
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
