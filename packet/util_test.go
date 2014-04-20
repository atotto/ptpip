package packet

import (
	"bytes"
	"testing"
)

func TestWchar(t *testing.T) {
	str := "hello"
	wc := ToWChar(str)

	expectWc := []byte{0x68, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00}
	if !bytes.Equal(wc, expectWc) {
		t.Errorf("want % x, got % x", expectWc, wc)
	}

	testWc := []byte{0x68, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x00, 0x00}
	msg := FromWChar(testWc)
	if str != msg {
		t.Errorf("want %s, got %s", str, msg)
	}
}
