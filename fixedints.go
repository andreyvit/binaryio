package binaryio

func (r *Reader) TryReadUint16() (uint16, bool) {
	if !r.need(2) {
		return 0, false
	}

	b := r.rem
	var v uint16
	if r.ByteOrder == LittleEndian {
		v = uint16(b[0])
		v |= uint16(b[1]) << (8 * 1)
	} else {
		v = uint16(b[1])
		v |= uint16(b[0]) << (8 * 1)
	}
	r.skipInternal(3)
	return v, true
}

func (r *Reader) TryReadUint24() (uint32, bool) {
	if !r.need(3) {
		return 0, false
	}

	b := r.rem
	var v uint32
	if r.ByteOrder == LittleEndian {
		v = uint32(b[0])
		v |= uint32(b[1]) << (8 * 1)
		v |= uint32(b[2]) << (8 * 2)
	} else {
		v = uint32(b[2])
		v |= uint32(b[1]) << (8 * 1)
		v |= uint32(b[0]) << (8 * 2)
	}
	r.skipInternal(3)
	return v, true
}

func (r *Reader) TryReadUint32() (uint32, bool) {
	if !r.need(4) {
		return 0, false
	}

	b := r.rem
	var v uint32
	if r.ByteOrder == LittleEndian {
		v = uint32(b[0])
		v |= uint32(b[1]) << (8 * 1)
		v |= uint32(b[2]) << (8 * 2)
		v |= uint32(b[3]) << (8 * 3)
	} else {
		v = uint32(b[3])
		v |= uint32(b[2]) << (8 * 1)
		v |= uint32(b[1]) << (8 * 2)
		v |= uint32(b[0]) << (8 * 3)
	}
	r.skipInternal(4)
	return v, true
}
func (r *Reader) ReadUint32() uint32 {
	v, _ := r.TryReadUint32()
	return v
}
func (r *Reader) ReadInt32() int32 {
	return int32(r.ReadUint32())
}

func (r *Reader) TryReadUint64() (uint64, bool) {
	if !r.need(8) {
		return 0, false
	}

	b := r.rem
	var v uint64
	if r.ByteOrder == LittleEndian {
		v = uint64(b[0])
		v |= uint64(b[1]) << (8 * 1)
		v |= uint64(b[2]) << (8 * 2)
		v |= uint64(b[3]) << (8 * 3)
		v |= uint64(b[4]) << (8 * 4)
		v |= uint64(b[5]) << (8 * 5)
		v |= uint64(b[6]) << (8 * 6)
		v |= uint64(b[7]) << (8 * 7)
	} else {
		v = uint64(b[7])
		v |= uint64(b[6]) << (8 * 1)
		v |= uint64(b[5]) << (8 * 2)
		v |= uint64(b[4]) << (8 * 3)
		v |= uint64(b[3]) << (8 * 4)
		v |= uint64(b[2]) << (8 * 5)
		v |= uint64(b[1]) << (8 * 6)
		v |= uint64(b[0]) << (8 * 7)
	}
	r.skipInternal(8)
	return v, true
}
func (r *Reader) ReadUint64() uint64 {
	v, _ := r.TryReadUint64()
	return v
}
func (r *Reader) ReadInt64() int64 {
	return int64(r.ReadUint64())
}

func (w *Writer) WriteUint24(v uint32) {
	var b [3]byte
	b[0] = byte(v)
	b[1] = byte(v >> (8 * 1))
	b[2] = byte(v >> (8 * 2))
	_, _ = w.buf.Write(b[:])
}

func (w *Writer) WriteUint32(v uint32) {
	var b [4]byte
	b[0] = byte(v)
	b[1] = byte(v >> (8 * 1))
	b[2] = byte(v >> (8 * 2))
	b[3] = byte(v >> (8 * 3))
	_, _ = w.buf.Write(b[:])
}

func (w *Writer) WriteUint64(v uint64) {
	var b [8]byte
	b[0] = byte(v)
	b[1] = byte(v >> (8 * 1))
	b[2] = byte(v >> (8 * 2))
	b[3] = byte(v >> (8 * 3))
	b[4] = byte(v >> (8 * 4))
	b[5] = byte(v >> (8 * 5))
	b[6] = byte(v >> (8 * 6))
	b[7] = byte(v >> (8 * 7))
	_, _ = w.buf.Write(b[:])
}

func (w *Writer) WriteInt32(v int32) {
	w.WriteUint32(uint32(v))
}

func (w *Writer) WriteInt64(v int64) {
	w.WriteUint64(uint64(v))
}

func (r *Reader) Int32(v *int32) {
	*v = r.ReadInt32()
}
func (r *Reader) Int64(v *int64) {
	*v = r.ReadInt64()
}
func (r *Reader) Uint32(v *uint32) {
	*v = r.ReadUint32()
}
func (r *Reader) Uint64(v *uint64) {
	*v = r.ReadUint64()
}
func (r *Reader) IntFixed32(v *int) {
	*v = int(r.ReadInt32())
}
func (r *Reader) IntFixed64(v *int) {
	*v = int(r.ReadInt64())
}
func (r *Reader) UintFixed32(v *uint) {
	*v = uint(r.ReadUint32())
}
func (r *Reader) UintFixed64(v *uint) {
	*v = uint(r.ReadUint64())
}

func (w *Writer) Int32(v *int32) {
	w.WriteInt32(*v)
}
func (w *Writer) Int64(v *int64) {
	w.WriteInt64(*v)
}
func (w *Writer) Uint32(v *uint32) {
	w.WriteUint32(*v)
}
func (w *Writer) Uint64(v *uint64) {
	w.WriteUint64(*v)
}
func (w *Writer) IntFixed32(v *int) {
	w.WriteInt32(int32(*v))
}
func (w *Writer) IntFixed64(v *int) {
	w.WriteInt64(int64(*v))
}
func (w *Writer) UintFixed32(v *uint) {
	w.WriteUint32(uint32(*v))
}
func (w *Writer) UintFixed64(v *uint) {
	w.WriteUint64(uint64(*v))
}
