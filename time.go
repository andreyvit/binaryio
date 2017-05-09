package binaryio

import (
	"time"
)

func (r *Reader) ReadTimeSec32() time.Time {
	u, ok := r.TryReadUint32()
	if !ok {
		return time.Time{}
	}
	return time.Unix(int64(u), 0).UTC()
}

func (r *Reader) ReadDuration() time.Duration {
	return time.Duration(r.ReadVarint64())
}

func (w *Writer) WriteTimeSec32(tm time.Time) {
	w.WriteUint32(uint32(tm.Unix()))
}
func (w *Writer) WriteDuration(v time.Duration) {
	w.WriteVarint64(int64(v))
}

func (r *Reader) TimeSec32(tm *time.Time) {
	*tm = r.ReadTimeSec32()
}
func (w *Writer) TimeSec32(tm *time.Time) {
	w.WriteTimeSec32(*tm)
}

func (r *Reader) Duration(d *time.Duration) {
	*d = r.ReadDuration()
}
func (w *Writer) Duration(d *time.Duration) {
	w.WriteDuration(*d)
}
