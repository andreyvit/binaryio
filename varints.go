package binaryio

import (
	"encoding/binary"
)

func (r *Reader) ReadVarint64() int64 {
	v, n := binary.Varint(r.rem)
	if n <= 0 {
		r.Fail(ErrInvalidVarint)
		return 0
	}
	r.skipInternal(n)
	return v
}

func (r *Reader) ReadVarint32() int32 {
	n := r.ReadVarint64()
	if int64(int32(n)) != n {
		r.Fail(ErrIntegerSizeMismatch)
		return 0
	}
	return int32(n)
}

func (r *Reader) ReadVarint() int {
	n := r.ReadVarint64()
	if int64(int(n)) != n {
		r.Fail(ErrIntegerSizeMismatch)
		return 0
	}
	return int(n)
}

func (r *Reader) ReadVarintUnsigned() int {
	n := r.ReadUvarint64()
	if uint64(int(n)) != n {
		r.Fail(ErrIntegerSizeMismatch)
		return 0
	}
	return int(n)
}

func (r *Reader) ReadUvarint64() uint64 {
	v, n := binary.Uvarint(r.rem)
	if n <= 0 {
		r.Fail(ErrInvalidVarint)
		return 0
	}
	r.skipInternal(n)
	return v
}

func (r *Reader) ReadUvarint32() uint32 {
	n := r.ReadUvarint64()
	if uint64(uint32(n)) != n {
		r.Fail(ErrIntegerSizeMismatch)
		return 0
	}
	return uint32(n)
}

func (r *Reader) ReadUvarint() uint {
	n := r.ReadUvarint64()
	if uint64(uint(n)) != n {
		r.Fail(ErrIntegerSizeMismatch)
		return 0
	}
	return uint(n)
}

func (r *Reader) Varint(v *int) {
	*v = r.ReadVarint()
}
func (r *Reader) VarintUnsigned(v *int) {
	*v = r.ReadVarintUnsigned()
}
func (r *Reader) Uvarint(v *uint) {
	*v = r.ReadUvarint()
}
func (r *Reader) Uvarint32(v *uint32) {
	*v = r.ReadUvarint32()
}
func (r *Reader) Uvarint64(v *uint64) {
	*v = r.ReadUvarint64()
}

func (w *Writer) WriteVarint64(v int64) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutVarint(buf[:], v)
	w.buf.Write(buf[:n])
}
func (w *Writer) WriteVarint32(v int32) {
	w.WriteVarint64(int64(v))
}
func (w *Writer) WriteVarint(v int) {
	w.WriteVarint64(int64(v))
}

func (w *Writer) WriteUvarint64(v uint64) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], v)
	w.buf.Write(buf[:n])
}
func (w *Writer) WriteUvarint32(v uint32) {
	w.WriteUvarint64(uint64(v))
}
func (w *Writer) WriteUvarint(v uint) {
	w.WriteUvarint64(uint64(v))
}
func (w *Writer) WriteVarintUnsigned(v int) {
	if v < 0 {
		panic("WriteVarintUnsigned with negative value")
	}
	w.WriteUvarint64(uint64(v))
}

func (w *Writer) Varint(v *int) {
	w.WriteVarint(*v)
}
func (w *Writer) VarintUnsigned(v *int) {
	w.WriteVarintUnsigned(*v)
}
func (w *Writer) Uvarint(v *uint) {
	w.WriteUvarint(*v)
}
func (w *Writer) Uvarint32(v *uint32) {
	w.WriteUvarint32(*v)
}
func (w *Writer) Uvarint64(v *uint64) {
	w.WriteUvarint64(*v)
}
