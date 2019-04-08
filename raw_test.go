package binaryio_test

import (
	"encoding/hex"
	"fmt"

	"github.com/andreyvit/binaryio"
)

func ExampleReader_ReadN() {
	data, err := hex.DecodeString("0102030405")
	if err != nil {
		panic(err)
	}
	r := binaryio.NewReader(data, binaryio.LittleEndian)

	b := r.ReadN(1)
	fmt.Printf("%x %v\n", b, r.Err())

	b = r.ReadN(3)
	fmt.Printf("%x %v\n", b, r.Err())

	b = r.ReadN(2)
	fmt.Printf("%x %v\n", b, r.Err())

	// Output: 01 <nil>
	// 020304 <nil>
	//  message too short at offset 4 decoding 01020304 <!> 05 (len=5)
}

func ExampleReader_ReadFull() {
	var r binaryio.Reader
	r.ResetErr(hex.DecodeString("0102030405"))

	var b [5]byte
	r.ReadFull(b[:1])
	fmt.Printf("%x %v\n", b[:1], r.Err())

	r.ReadFull(b[:3])
	fmt.Printf("%x %v\n", b[:3], r.Err())

	r.ReadFull(b[:2])
	fmt.Printf("%x %v\n", b[:2], r.Err())

	// Output: 01 <nil>
	// 020304 <nil>
	// 0203 message too short at offset 4 decoding 01020304 <!> 05 (len=5)
}

func ExampleReader_Read() {
	var r binaryio.Reader
	r.ResetErr(hex.DecodeString("0102030405"))

	var b [5]byte
	n, err := r.Read(b[:1])
	fmt.Printf("%x %v\n", b[:n], err)

	n, err = r.Read(b[:3])
	fmt.Printf("%x %v\n", b[:n], err)

	n, err = r.Read(b[:2])
	fmt.Printf("%x %v\n", b[:n], err)

	n, err = r.Read(b[:2])
	fmt.Printf("%d %v\n", n, err)

	// Output: 01 <nil>
	// 020304 <nil>
	// 05 <nil>
	// 0 EOF
}

func ExampleWriter_Write() {
	var w binaryio.Writer

	w.Write([]byte{1, 2, 3})
	fmt.Printf("%x\n", w.Bytes())

	w.Write([]byte{4, 5})
	fmt.Printf("%x\n", w.Bytes())

	// Output: 010203
	// 0102030405
}
