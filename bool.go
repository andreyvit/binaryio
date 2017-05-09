package binaryio

func (r *Reader) ReadBool() bool {
	b := r.ReadByte()
	if b == 0 {
		return false
	} else if b == 1 {
		return true
	} else {
		r.Fail(ErrInvalidBool)
		return false
	}
}

func (w *Writer) WriteBool(v bool) {
	if v {
		w.WriteByte(1)
	} else {
		w.WriteByte(0)
	}
}

func (r *Reader) Bool(v *bool) {
	*v = r.ReadBool()
}
func (w *Writer) Bool(v *bool) {
	w.WriteBool(*v)
}
