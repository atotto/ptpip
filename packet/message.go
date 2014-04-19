package packet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"unsafe"
)

type InitCommand struct {
	BaseLayout
	GUID         [16]byte
	FriendlyName [20]byte
}

func NewInitCommand(guid, friendlyName string) *InitCommand {
	var p InitCommand
	copy(p.GUID[:], guid)
	copy(p.FriendlyName[:], friendlyName)
	p.BaseLayout = *NewBaseLayout(InitCommandRequestPacket, int(unsafe.Sizeof(p)-8))
	return &p
}

func (p *InitCommand) Send(w io.Writer) (err error) {
	return binary.Write(w, binary.LittleEndian, BaseLayout{0, InitCommandRequestPacket})
}

func (p *InitCommand) Recv(r io.Reader) (sessionID uint32, guid, friendlyName string, err error) {
	base, payload, err := Recv(r)
	if err != nil {
		return
	}
	if base.Len < 21 {
		err = fmt.Errorf("Invalid packet size: %d", base.Len)
		return
	}

	switch base.Typ {
	case InitCommandAck:
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
	return
}

type InitEvent struct {
	SessionID uint32
}

func (p *InitEvent) Send(w io.Writer) (err error) {
	return binary.Write(w, binary.LittleEndian, BaseLayout{0, InitEventRequestPacket})
}

type Operation struct {
	BaseLayout
}
