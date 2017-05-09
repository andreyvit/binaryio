package binaryio

import (
	"math/big"
	"time"
)

type Coder interface {
	IsReading() bool

	Ver(curver int) int

	Byte(b *byte)
	FixedBuf(buf []byte)
	FixedBufPtr(buf *[]byte, n int)

	Varint(v *int)
	VarintUnsigned(v *int)
	Uvarint(v *uint)
	Uvarint32(v *uint32)
	Uvarint64(v *uint64)

	Int32(v *int32)
	Int64(v *int64)
	Uint32(v *uint32)
	Uint64(v *uint64)
	IntFixed32(v *int)
	IntFixed64(v *int)
	UintFixed32(v *uint)
	UintFixed64(v *uint)

	Bool(v *bool)

	Blob(v *[]byte)
	String(v *string)
	BigInt(v **big.Int)

	TimeSec32(tm *time.Time)
	Duration(d *time.Duration)
}

func (r *Reader) IsReading() bool {
	return true
}

func (w *Writer) IsReading() bool {
	return false
}
