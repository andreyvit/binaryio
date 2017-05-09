package binaryio

func (r *Reader) Ver(curver int) int {
	return r.ReadVarintUnsigned()
}

func (w *Writer) Ver(curver int) int {
	w.WriteVarintUnsigned(curver)
	return curver
}
