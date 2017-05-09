package binaryio_test

import (
	"github.com/andreyvit/binaryio"
	"math"
	"testing"
)

func TestFixedIntsRoundtrip(t *testing.T) {
	var w binaryio.Writer

	w.WriteInt32(1000000000)
	w.WriteInt64(1000000000000000)
	w.WriteUint32(math.MaxUint32)
	w.WriteUint64(math.MaxUint64)

	r := binaryio.NewReader(w.Bytes())
	if a, e := r.ReadInt32(), int32(1000000000); a != e {
		t.Fatalf("ReadInt32() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadInt64(), int64(1000000000000000); a != e {
		t.Fatalf("ReadInt64() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadUint32(), uint32(math.MaxUint32); a != e {
		t.Fatalf("ReadUint32() == %v, wanted %v", a, e)
	}
	if a, e := r.ReadUint64(), uint64(math.MaxUint64); a != e {
		t.Fatalf("ReadUint64() == %v, wanted %v", a, e)
	}
}
