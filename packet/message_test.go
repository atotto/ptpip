package packet_test

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/atotto/ptpip/packet"
)

func pack(t testing.TB, typ packet.PacketType, args ...interface{}) (b []byte) {
	buf := new(bytes.Buffer)
	n := 0
	for _, v := range args {
		n += binary.Size(v)
	}
	base := packet.NewBaseLayout(typ, n)
	binaryWrite(t, buf, base)

	for _, v := range args {
		binaryWrite(t, buf, v)
	}
	return buf.Bytes()
}

func binaryWrite(t testing.TB, w io.Writer, v interface{}) {
	if err := binary.Write(w, binary.LittleEndian, v); err != nil {
		t.Fatalf("binary.Write failed: %v", err)
	}
}

func TestInitCommandRequest(t *testing.T) {
	testInitCommandRequest(t, 1)
}

func BenchmarkInitCommandRequest(b *testing.B) {
	testInitCommandRequest(b, b.N)
}

func testInitCommandRequest(t testing.TB, N int) {
	// Setup
	var expect_packetLayout []byte
	{
		guid := make([]byte, 16)
		copy(guid[:], "1234567890123456")
		friendlyName := []byte("golang_ptpip_client")
		expect_packetLayout = pack(t, packet.InitCommandRequestPacket, guid, friendlyName)
	}

	for i := 0; i < N; i++ {
		buf := new(bytes.Buffer)
		// Test
		err := packet.SendInitCommand(buf, "1234567890123456", "golang_ptpip_client")

		// Verify
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !bytes.Equal(expect_packetLayout, buf.Bytes()) {
			t.Errorf("packet fail")
		}
	}
}

func TestInitCommandAck(t *testing.T) {
	testInitCommandAck(t, 1)
}

func BenchmarkInitCommandAck(b *testing.B) {
	testInitCommandAck(b, b.N)
}

func testInitCommandAck(t testing.TB, N int) {
	// Setup
	expect_sessionID := uint32(123)
	expect_guid := "1234567890123456"
	expect_friendlyName := "golang_ptpip_client"
	var test_packet []byte
	{
		sessionID := expect_sessionID
		guid := make([]byte, 16)
		copy(guid[:], expect_guid)
		str := expect_friendlyName
		friendlyName := make([]byte, len(str))
		copy(friendlyName[:], str)

		test_packet = pack(t, packet.InitCommandAckPacket, sessionID, guid, friendlyName)
	}

	for i := 0; i < N; i++ {
		// Test
		sessionID, guid, friendlyName, err := packet.RecvInitCommand(bytes.NewReader(test_packet))

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
}

func TestInitEventRequest(t *testing.T) {
	// Setup
	sessionID := uint32(1234)
	expect_packetLayout := pack(t, packet.InitEventRequestPacket, sessionID)

	// Test
	buf := new(bytes.Buffer)
	err := packet.SendInitEvent(buf, sessionID)

	// Verify
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !bytes.Equal(expect_packetLayout, buf.Bytes()) {
		t.Errorf("packet fail")
	}
}

func TestInitEventAck(t *testing.T) {
	// Setup
	test_packet := pack(t, packet.InitEventAckPacket)

	// Test
	err := packet.RecvInitEvent(bytes.NewReader(test_packet))

	// Verify
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
