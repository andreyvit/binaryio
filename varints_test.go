package binaryio_test

import (
	"math"
	"testing"

	"github.com/andreyvit/binaryio"
)

func TestVarintsRoundtrip(t *testing.T) {
	var w binaryio.Writer

	w.WriteVarint(1000)
	w.WriteVarint32(1000000000)
	w.WriteVarint64(1000000000000000)
	w.WriteUvarint32(math.MaxUint32)
	w.WriteUvarint64(math.MaxUint64)
	w.WriteUvarint(1000000000000000)

	r := binaryio.NewReader(w.Bytes(), binaryio.LittleEndian)
	if a, e := r.ReadVarint(), 1000; a != e {
		t.Fatalf("ReadVarint() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadVarint32(), int32(1000000000); a != e {
		t.Fatalf("ReadVarint32() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadVarint64(), int64(1000000000000000); a != e {
		t.Fatalf("ReadVarint64() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadUvarint32(), uint32(math.MaxUint32); a != e {
		t.Fatalf("ReadUvarint32() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadUvarint64(), uint64(math.MaxUint64); a != e {
		t.Fatalf("ReadUvarint64() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadUvarint(), uint(1000000000000000); a != e {
		t.Fatalf("ReadUvarint() == %v, wanted %v", a, e)
	}
}
