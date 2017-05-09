package binaryio

import (
	"bytes"
)

type Writer struct {
	buf bytes.Buffer
}

func (w *Writer) Clear() {
	w.buf.Truncate(0)
}

func (w *Writer) Bytes() []byte {
	return w.buf.Bytes()
}

func (w *Writer) WriteZeroBytes(n int) {
	w.buf.Grow(n)
	for i := 0; i < n; i++ {
		_ = w.buf.WriteByte(0)
	}
}
