package binaryio

import (
	"errors"
	"fmt"
)

var ErrMessageTooShort = errors.New("message too short")
var ErrTrailingData = errors.New("unexpected trailing data")
var ErrInvalidBool = errors.New("bool not 0 or 1")
var ErrInvalidVarint = errors.New("invalid varint")
var ErrIntegerSizeMismatch = errors.New("integer size mismatch")

type Reader struct {
	rem  []byte
	orig []byte
	err  error
	offs int64
}

type Error struct {
	Err    error
	Offset int64
	Data   []byte
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v at offset %d decoding %x <!> %x (len=%d)", e.Err, e.Offset, e.Data[:e.Offset], e.Data[e.Offset:], len(e.Data))
}

func NewReader(data []byte) *Reader {
	r := new(Reader)
	r.Reset(data)
	return r
}

func (r *Reader) Reset(data []byte) {
	*r = Reader{data, data, nil, 0}
}

func (r *Reader) ResetErr(data []byte, err error) {
	*r = Reader{data, data, err, 0}
}

func (r *Reader) Fail(err error) {
	if r.err == nil {
		r.err = &Error{err, r.offs, r.orig}
	}
}

func (r *Reader) Err() error {
	return r.err
}

func (r *Reader) Failed() bool {
	return r.err != nil
}

func (r *Reader) need(cb int) bool {
	if r.err != nil {
		return false
	}
	if len(r.rem) < cb {
		r.Fail(ErrMessageTooShort)
		return false
	} else {
		return true
	}
}

func (r *Reader) skipInternal(cb int) {
	r.rem = r.rem[cb:]
	r.offs += int64(cb)
}

func (r *Reader) ExpectEOF() {
	if len(r.rem) > 0 {
		r.Fail(ErrTrailingData)
	}
}

func (r *Reader) Skip(n int) {
	if r.need(n) {
		r.skipInternal(n)
	}
}
