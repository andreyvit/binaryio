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

func Encode(e Encodable, bo ByteOrder) []byte {
	var w Writer
	e.Encode(&w)
	return w.Bytes()
}

func Decode(d Decodable, data []byte, bo ByteOrder) error {
	var r Reader
	r.ResetBO(data, bo)
	d.Decode(&r)
	r.ExpectEOF()
	return r.Err()
}

func Serialize(c Serializable, bo ByteOrder) []byte {
	var w Writer
	c.Serialize(&w)
	return w.Bytes()
}

func Deserialize(c Serializable, data []byte, bo ByteOrder) error {
	var r Reader
	r.ResetBO(data, bo)
	c.Serialize(&r)
	r.ExpectEOF()
	return r.Err()
}
