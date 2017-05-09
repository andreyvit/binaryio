package binaryio_test

import (
	"encoding/hex"
	"fmt"
	"github.com/andreyvit/binaryio"
)

func ExampleReader() {
	data, _ := hex.DecodeString("02400a48656c6c6fe80787308856")

	r := binaryio.NewReader(data)
	ver := r.ReadUvarint()
	id := r.ReadVarint()
	name := r.ReadString()
	score := r.ReadUvarint64()
	time := r.ReadTimeSec32()
	r.ExpectEOF()

	if err := r.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("ver=%v id=%v name=%v score=%v time=%v\n", ver, id, name, score, time)
	// Output: ver=2 id=32 name=Hello score=1000 time=2016-01-02 20:18:15 +0000 UTC
}
