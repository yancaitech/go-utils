package utils

import "encoding/binary"

// Amount represents the base bitcoin monetary unit (colloquially referred
// to as a `Satoshi').  A single Amount is equal to 1e-8 of a bitcoin.
type Amount int64

// Round converts a floating point number, which may or may not be representable
// as an integer, to the Amount integer type by rounding to the nearest integer.
// This is performed by adding or subtracting 0.5 depending on the sign, and
// relying on integer truncation to round the value to the nearest Amount.
func Round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}
	return Amount(f + 0.5)
}

// Int64ToBytes func
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// BytesToInt64 func
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// Uint64ToBytes func
func Uint64ToBytes(i uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], i)
	return buf[:]
}

// BytesToUint64 func
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Add function to add two uint64 numbers
func Add(existing, new []byte) []byte {
	return Uint64ToBytes(BytesToUint64(existing) + BytesToUint64(new))
}
