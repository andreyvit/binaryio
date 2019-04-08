package binaryio_test

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/andreyvit/binaryio"
)

type Bar struct {
	ID    int
	Name  string
	Score uint64
	Time  time.Time
}

func (bar *Bar) Encode(w *binaryio.Writer) {
	w.WriteUvarint(2)
	w.WriteVarint(bar.ID)
	w.WriteString(bar.Name)
	w.WriteUvarint64(bar.Score)
	w.WriteTimeSec32(bar.Time)
}

func (bar *Bar) Decode(r *binaryio.Reader) {
	ver := r.ReadUvarint()
	bar.ID = r.ReadVarint()
	bar.Name = r.ReadString()
	bar.Score = r.ReadUvarint64()
	if ver >= 2 {
		bar.Time = r.ReadTimeSec32()
	}
}

func Example() {
	a := Bar{32, "Hello", 1000, time.Date(2016, 1, 2, 20, 18, 15, 0, time.UTC)}
	data := binaryio.Encode(&a, binaryio.LittleEndian)
	fmt.Printf("Encode() = %x\n", data)

	var b Bar
	err := binaryio.Decode(&b, data, binaryio.LittleEndian)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode(data) = %v\n", b)

	old, _ := hex.DecodeString("01400a48656c6c6fe807")

	var c Bar
	err = binaryio.Decode(&c, old, binaryio.LittleEndian)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode(old) = %v\n", c)

	// Output: Encode() = 02400a48656c6c6fe80787308856
	// Decode(data) = {32 Hello 1000 2016-01-02 20:18:15 +0000 UTC}
	// Decode(old) = {32 Hello 1000 0001-01-01 00:00:00 +0000 UTC}
}
