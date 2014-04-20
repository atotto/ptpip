package packet

import (
	"bytes"
	"encoding/binary"
)

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

func Pack(args ...interface{}) (b []byte, err error) {
	buf := new(bytes.Buffer)
	for _, v := range args {
		if err := binary.Write(buf, endian, v); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func Unpack(b []byte, args ...interface{}) (err error) {
	buf := bytes.NewReader(b)
	for _, v := range args {
		if err := binary.Read(buf, endian, &v); err != nil {
			return err
		}
	}
	return nil
}

func Length(args ...interface{}) (n int) {
	for _, v := range args {
		n += binary.Size(v)
	}
	return n
}

func ToWChar(str string) []byte {
	buf := []byte(str)
	wc := make([]byte, len(buf)*2)
	for i, c := range buf {
		wc[i*2] = c
		wc[i*2+1] = 0
	}
	return wc
}

func FromWChar(wc []byte) string {
	b := make([]byte, len(wc)/2)
	var i int
	for i = 0; i < len(b); i++ {
		if wc[i*2] == 0 {
			break
		}
		b[i] = wc[i*2]
	}
	return string(b[0:i])
}
