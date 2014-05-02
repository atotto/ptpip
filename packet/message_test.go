package packet_test

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/atotto/ptpip/packet"
	"github.com/atotto/ptpip/packet/util"
	"github.com/atotto/ptpip/ptp"
)

func pack(t testing.TB, typ packet.PacketType, args ...interface{}) (b []byte) {
	buf := new(bytes.Buffer)
	n := 8
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
		friendlyName := util.ToWChar("golang_ptpip_client")
		n := []byte{0, 0}
		expect_packetLayout = pack(t, packet.InitCommandRequestPacket, guid, friendlyName, n)
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
		friendlyName := util.ToWChar(expect_friendlyName)
		n := []byte{0, 0}

		test_packet = pack(t, packet.InitCommandAckPacket, sessionID, guid, friendlyName, n)
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

func TestOperationRequest(t *testing.T) {
	// Setup
	dataPheseInfo := uint32(123)
	operationCode := ptp.OC_InitiateCapture
	transactionID := uint32(1)
	parameters := []uint32{1, 2, 3}

	expect_packetLayout := pack(t, packet.OperationRequestPacket, dataPheseInfo, operationCode, transactionID, parameters)

	// Test
	buf := new(bytes.Buffer)
	err := packet.SendOperationRequest(buf, dataPheseInfo, operationCode, transactionID, parameters)

	// Verify
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !bytes.Equal(expect_packetLayout, buf.Bytes()) {
		t.Errorf("packet fail")
	}
}

func TestOperationResponse(t *testing.T) {
	// Setup
	expect_responseCode := ptp.RC_OK
	expect_transactionID := uint32(1)
	expect_parameters := []uint32{1, 2, 3}

	test_packet := pack(t, packet.OperationResponsePacket, expect_responseCode, expect_transactionID, expect_parameters)

	// Test
	responseCode, transactionID, parameters, err := packet.RecvOperationResponse(bytes.NewReader(test_packet))

	// Verify
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if responseCode != expect_responseCode {
		t.Errorf("want responseCode %d, got %d", expect_responseCode, responseCode)
	}
	if transactionID != expect_transactionID {
		t.Errorf("want transactionID %d, got %d", expect_transactionID, transactionID)
	}
	if len(parameters) != len(expect_parameters) {
		t.Errorf("want parameters %+v, got %+v", expect_parameters, parameters)
	}
	for n, v := range parameters {
		if v != expect_parameters[n] {
			t.Fatalf("want parameters %+v, got %+v", expect_parameters, parameters)
		}
	}
}
