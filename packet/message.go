package packet

import (
	"errors"
	"fmt"
	"io"

	"github.com/atotto/ptpip/ptp"
)

func SendInitCommand(w io.Writer, guid, friendlyName string) (err error) {
	p := struct {
		GUID         []byte
		FriendlyName []byte
	}{}
	p.GUID = []byte(guid)
	p.FriendlyName = ToWChar(friendlyName)
	n := []byte{0, 0}

	b := append(p.GUID[0:16], p.FriendlyName[:]...)
	b = append(b, n...)
	return Send(w, InitCommandRequestPacket, b)
}

func RecvInitCommand(r io.Reader) (sessionID uint32, guid, friendlyName string, err error) {
	base, payload, err := Recv(r)
	if err != nil {
		return
	}

	switch base.Typ {
	case InitCommandAckPacket:
		if base.Len < 21 {
			err = fmt.Errorf("Invalid packet size: %d", base.Len)
			return
		}
		sessionID = Uint32(payload[0:4])
		guid = String(payload[4:20])
		friendlyName = FromWChar(payload[20 : base.Len-8])
		return
	case InitFailPacket:
		reason := Uint32(payload[0:4])
		err = fmt.Errorf("Initialise Failed. reason code: %d", reason)
		return
	default:
		err = errors.New("Invalid State.")
		return
	}
}

func SendInitEvent(w io.Writer, sessionID uint32) (err error) {
	b := make([]byte, 4)
	PutUint32(b, sessionID)
	return Send(w, InitEventRequestPacket, b)
}

func RecvInitEvent(r io.Reader) (err error) {
	base, payload, err := Recv(r)
	if err != nil {
		return
	}
	switch base.Typ {
	case InitEventAckPacket:
		return
	case InitFailPacket:
		reason := Uint32(payload[0:4])
		err = fmt.Errorf("Initialise Failed. reason code: %d", reason)
		return
	default:
		err = errors.New("Invalid State.")
		return
	}
}

func SendOperationRequest(w io.Writer, dataPheseInfo uint32, operationCode ptp.OperationCode, transactionID uint32, parameters []uint32) (err error) {
	b, err := Pack(dataPheseInfo, operationCode, transactionID, parameters)
	if err != nil {
		return
	}
	return Send(w, OperationRequestPacket, b)
}

func RecvOperationResponse(r io.Reader) (responseCode ptp.OperationResponseCode, transactionID uint32, parameters []uint32, err error) {
	base, payload, err := Recv(r)
	if err != nil {
		return
	}
	switch base.Typ {
	case OperationResponsePacket:
		responseCode = ptp.OperationResponseCode(Uint16(payload[0:2]))
		transactionID = Uint32(payload[2:6])
		parameters = make([]uint32, (base.Len-8-6)/4)
		for i := 0; i < len(parameters); i++ {
			parameters[i] = Uint32(payload[6+i*4 : 6+(i+1)*4])
		}
		return
	default:
		err = errors.New("Invalid State.")
		return
	}
}

func RecvStartData(r io.Reader) (err error) {
	panic("NotImplementedYet")
}

func RecvData(r io.Reader) (err error) {
	panic("NotImplementedYet")
}

func RecvEndData(r io.Reader) (err error) {
	panic("NotImplementedYet")
}
