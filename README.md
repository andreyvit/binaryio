# binaryio

Buffered Reader and Writer for serializing structured binary data.

See [docs on godoc.org](https://godoc.org/github.com/andreyvit/binaryio).


## Installation

    go get github.com/andreyvit/binaryio


The format is positional: readers and writers must agree on field order, byte order and encoding. Store an explicit format version when evolving records; `Coder.Ver` supports conditional fields, as shown below. Changing an existing field’s encoding requires a migration or a new version.

## Example (Reader)

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/andreyvit/binaryio"
)

func main() {
	data, _ := hex.DecodeString("02400a48656c6c6fe80787308856")

	r := binaryio.NewReader(data, binaryio.LittleEndian)
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
	// ver=2 id=32 name=Hello score=1000 time=2016-01-02 20:18:15 +0000 UTC
}
```

## Example (Serialize)

```go
package main

import (
	"fmt"
	"time"

	"github.com/andreyvit/binaryio"
)

type Highscore struct {
	Player int
	Name   string
	Score  uint64
	Time   time.Time
}

func (hs *Highscore) Serialize(c binaryio.Coder) {
	ver := c.Ver(2)
	c.Varint(&hs.Player)
	c.String(&hs.Name)
	c.Uvarint64(&hs.Score)
	if ver >= 2 {
		c.TimeSec32(&hs.Time)
	}
}

func main() {
	a := Highscore{32, "Hello", 1000, time.Date(2016, 1, 2, 20, 18, 15, 0, time.UTC)}
	data := binaryio.Serialize(&a, binaryio.LittleEndian)
	fmt.Printf("%x\n", data)
	// 02400a48656c6c6fe80787308856

	var b Highscore
	err := binaryio.Deserialize(&b, data, binaryio.LittleEndian)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b)
	// {32 Hello 1000 2016-01-02 20:18:15 +0000 UTC}
}
```

