package binaryio

type Encodable interface {
	Encode(w *Writer)
}

type Decodable interface {
	Decode(r *Reader)
}

type Serializable interface {
	Serialize(c Coder)
}

func Encode(e Encodable) []byte {
	var w Writer
	e.Encode(&w)
	return w.Bytes()
}

func Decode(d Decodable, data []byte) error {
	var r Reader
	r.Reset(data)
	d.Decode(&r)
	r.ExpectEOF()
	return r.Err()
}

func Serialize(c Serializable) []byte {
	var w Writer
	c.Serialize(&w)
	return w.Bytes()
}

func Deserialize(c Serializable, data []byte) error {
	var r Reader
	r.Reset(data)
	c.Serialize(&r)
	r.ExpectEOF()
	return r.Err()
}
