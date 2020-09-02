你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# binaryio

Buffered Reader and Writer for serializing structured binary data.

See [docs on godoc.org](https://godoc.org/github.com/andreyvit/binaryio).


## Installation

    go get -u github.com/andreyvit/binaryio


## Example (Reader)

```go
func main() {
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
    // ver=2 id=32 name=Hello score=1000 time=2016-01-02 20:18:15 +0000 UTC
}
```

## Example (Serialize)

```go
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
    data := binaryio.Serialize(&a)
    fmt.Printf("%x\n", data)
    // 02400a48656c6c6fe80787308856

    var b Highscore
    err := binaryio.Deserialize(&b, data)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v\n", b)
    // {32 Hello 1000 2016-01-02 20:18:15 +0000 UTC}
}
```

