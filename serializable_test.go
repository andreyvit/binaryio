package binaryio_test

import (
	"encoding/hex"
	"fmt"
	"github.com/andreyvit/binaryio"
	"time"
)

type Foo struct {
	ID    int
	Name  string
	Score uint64
	Time  time.Time
}

func (foo *Foo) Serialize(c binaryio.Coder) {
	ver := c.Ver(2)
	c.Varint(&foo.ID)
	c.String(&foo.Name)
	c.Uvarint64(&foo.Score)
	if ver >= 2 {
		c.TimeSec32(&foo.Time)
	}
}

func Example_serializable() {
	a := Foo{32, "Hello", 1000, time.Date(2016, 1, 2, 20, 18, 15, 0, time.UTC)}
	data := binaryio.Serialize(&a)
	fmt.Printf("Encode() = %x\n", data)

	var b Foo
	err := binaryio.Deserialize(&b, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode(data) = %v\n", b)

	old, _ := hex.DecodeString("01400a48656c6c6fe807")

	var c Foo
	err = binaryio.Deserialize(&c, old)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode(old) = %v\n", c)

	// Output: Encode() = 02400a48656c6c6fe80787308856
	// Decode(data) = {32 Hello 1000 2016-01-02 20:18:15 +0000 UTC}
	// Decode(old) = {32 Hello 1000 0001-01-01 00:00:00 +0000 UTC}
}
