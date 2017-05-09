package binaryio

import (
	"io"
)

func (r *Reader) TryReadByte() (byte, bool) {
	if !r.need(1) {
		return 0, false
	}
	v := r.rem[0]
	r.skipInternal(1)
	return v, true
}

func (r *Reader) ReadByte() byte {
	v, _ := r.TryReadByte()
	return v
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.err != nil {
		return 0, r.err
	}
	if len(r.rem) == 0 {
		err = io.EOF
	}

	n = copy(p, r.rem)
	r.skipInternal(n)
	return
}

func (r *Reader) ReadFull(buf []byte) {
	n := len(buf)
	if !r.need(n) {
		return
	}
	copy(buf, r.rem)
	r.skipInternal(n)
}

func (r *Reader) ReadN(n int) []byte {
	if !r.need(n) {
		return nil
	}
	v := r.rem[:n]
	r.skipInternal(n)
	return v
}

// WriteByte writes the given byte to the buffer. Always returns nil.
func (w *Writer) WriteByte(v byte) error {
	_ = w.buf.WriteByte(v)
	return nil
}

// Write writes the given raw bytes to the output stream. Always returns a nil error.
func (w *Writer) Write(v []byte) (n int, err error) {
	_, _ = w.buf.Write(v)
	return len(v), nil
}

func (r *Reader) Byte(b *byte) {
	*b = r.ReadByte()
}
func (r *Reader) FixedBuf(buf []byte) {
	r.ReadFull(buf)
}
func (r *Reader) FixedBufPtr(buf *[]byte, n int) {
	*buf = r.ReadN(n)
}

func (w *Writer) Byte(b *byte) {
	w.WriteByte(*b)
}
func (w *Writer) FixedBuf(buf []byte) {
	w.Write(buf)
}
func (w *Writer) FixedBufPtr(buf *[]byte, n int) {
	if len(*buf) != n {
		panic("len(buf) != n")
	}
	w.Write(*buf)
}
