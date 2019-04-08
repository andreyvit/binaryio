package binaryio

type ByteOrder int

const (
	BigEndian    = ByteOrder(0)
	LittleEndian = ByteOrder(1)
)
