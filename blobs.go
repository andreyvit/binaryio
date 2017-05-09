package binaryio

import (
	"math/big"
)

func (r *Reader) ReadBlob() []byte {
	n := r.ReadVarint()
	return r.ReadN(n)
}

func (r *Reader) ReadBigInt() *big.Int {
	b := r.ReadBlob()
	if b == nil {
		return nil
	}

	n := new(big.Int)
	n.SetBytes(b)
	return n
}

func (r *Reader) ReadString() string {
	b := r.ReadBlob()
	if b == nil {
		return ""
	}
	return string(b)
}

func (w *Writer) WriteBlob(v []byte) {
	w.WriteVarint(len(v))
	w.Write(v)
}

func (w *Writer) WriteString(s string) {
	w.WriteBlob([]byte(s))
}

func (w *Writer) WriteBigInt(v *big.Int) {
	b := v.Bytes()
	if len(b) == 0 {
		b = []byte{0}
	}

	w.WriteBlob(b)
}

func (r *Reader) Blob(v *[]byte) {
	*v = r.ReadBlob()
}
func (r *Reader) String(v *string) {
	*v = r.ReadString()
}
func (r *Reader) BigInt(v **big.Int) {
	*v = r.ReadBigInt()
}

func (w *Writer) Blob(v *[]byte) {
	w.WriteBlob(*v)
}
func (w *Writer) String(v *string) {
	w.WriteString(*v)
}
func (w *Writer) BigInt(v **big.Int) {
	w.WriteBigInt(*v)
}
