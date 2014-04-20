package packet

import (
	"errors"
	"fmt"
	"io"
)

func SendInitCommand(w io.Writer, guid, friendlyName string) (err error) {
	p := struct {
		GUID         [16]byte
		FriendlyName []byte
	}{}
	copy(p.GUID[:], guid)
	p.FriendlyName = []byte(friendlyName)

	b := append(p.GUID[:], p.FriendlyName[:]...)
	return Send(w, InitCommandRequestPacket, b)
}

func RecvInitCommand(r io.Reader) (sessionID uint32, guid, friendlyName string, err error) {
	base, payload, err := Recv(r)
	if err != nil {
		return
	}

	switch base.Typ {
	case InitCommandAck:
		if base.Len < 21 {
			err = fmt.Errorf("Invalid packet size: %d", base.Len)
			return
		}
		sessionID = Uint32(payload[0:4])
		guid = String(payload[4:20])
		friendlyName = String(payload[20:base.Len])
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

type Operation struct {
	BaseLayout
}
