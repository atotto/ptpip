package packet_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
	"unsafe"

	"github.com/atotto/ptpip/packet"
)

func pack(t *testing.T, w io.Writer, typ packet.PacketType, args ...interface{}) {
	n := 0
	for _, v := range args {
		n += binary.Size(v)
	}
	base := packet.NewBaseLayout(typ, n)
	binaryWrite(t, w, base)

	for _, v := range args {
		binaryWrite(t, w, v)
	}
}

func binaryWrite(t *testing.T, w io.Writer, v interface{}) {
	if err := binary.Write(w, binary.LittleEndian, v); err != nil {
		t.Fatalf("binary.Write failed: %v", err)
	}
}

func TestInitCommandRequest(t *testing.T) {
	p := packet.NewInitCommand(
		"2adf2521-375c-406a-bc08-3c662406090e",
		"golang_ptpip_client")

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, p); err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// fmt.Printf("%# x\n", buf.Bytes())
	// fmt.Println(unsafe.Sizeof(*p) - 8)

	if expect := uint32(unsafe.Sizeof(*p)) - 8; p.Len != expect {
		t.Errorf("want packet length %d, got %d", expect, p.Len)
	}
}

func TestInitCommandAck(t *testing.T) {
	// Setup
	buf := new(bytes.Buffer)
	expect_sessionID := uint32(123)
	expect_guid := "1234567890123456"
	expect_friendlyName := "golang_ptpip_client"
	{
		sessionID := expect_sessionID
		guid := make([]byte, 16)
		copy(guid[:], expect_guid)
		str := expect_friendlyName
		friendlyName := make([]byte, len(str))
		copy(friendlyName[:], str)

		pack(t, buf, packet.InitCommandAck, sessionID, guid, friendlyName)
	}
	p := packet.NewInitCommand("", "")

	// Test
	sessionID, guid, friendlyName, err := p.Recv(buf)

	// Verify
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sessionID != expect_sessionID {
		t.Errorf("want sessionID %d, got %d", expect_sessionID, sessionID)
	}
	if guid != expect_guid {
		t.Errorf("want GUID %s, got %s", expect_guid, guid)
	}
	if friendlyName != expect_friendlyName {
		t.Errorf("want frendlyname %s, got %s", expect_friendlyName, friendlyName)
	}
}
