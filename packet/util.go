package packet

import "encoding/binary"

var endian = binary.LittleEndian

func Uint16(b []byte) uint16 {
	return endian.Uint16(b)
}

func Uint32(b []byte) uint32 {
	return endian.Uint32(b)
}

func Uint64(b []byte) uint64 {
	return endian.Uint64(b)
}

func String(b []byte) string {
	return string(b)
}

func PutUint16(b []byte, v uint16) {
	endian.PutUint16(b, v)
}

func PutUint32(b []byte, v uint32) {
	endian.PutUint32(b, v)
}

func PutUint64(b []byte, v uint64) {
	endian.PutUint64(b, v)
}

func PutString(b []byte, v string) (n int) {
	n = copy(b[:], v)
	return
}
