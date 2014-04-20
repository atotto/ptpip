package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type PacketType uint32

const (
	InitCommandRequestPacket PacketType = 0x00000001
	InitCommandAckPacket     PacketType = 0x00000002

	InitEventRequestPacket PacketType = 0x00000003
	InitEventAckPacket     PacketType = 0x00000004

	InitFailPacket PacketType = 0x00000005

	OperationRequestPacket  PacketType = 0x00000006
	OperationResponsePacket PacketType = 0x00000007

	EventPacket     PacketType = 0x00000008
	StartDataPacket PacketType = 0x00000009
	DataPacket      PacketType = 0x0000000A
	CancelPacket    PacketType = 0x0000000B
	EndDataPacket   PacketType = 0x0000000C

	ProbeRequestPacket  PacketType = 0x0000000D
	ProbeResponsePacket PacketType = 0x0000000E
)

type Event interface {
}

type Command interface {
}

type Data interface {
}

type BaseLayout struct {
	Len uint32
	Typ PacketType
}

func NewBaseLayout(typ PacketType, length int) *BaseLayout {
	var p BaseLayout
	p.Typ = typ
	p.Len = uint32(length)
	return &p
}

func Send(w io.Writer, typ PacketType, payload []byte) (err error) {
	base := NewBaseLayout(typ, len(payload))
	if err = binary.Write(w, binary.LittleEndian, base); err != nil {
		return
	}
	var n int
	if n, err = w.Write(payload); err != nil {
		return
	}
	if n != len(payload) {
		return fmt.Errorf("write byte size mismatch: %d != %d", len(payload), n)
	}
	return
}

func Recv(r io.Reader) (base BaseLayout, payload []byte, err error) {
	if err = binary.Read(r, binary.LittleEndian, &base); err != nil {
		return
	}
	payload = make([]byte, base.Len)
	_, err = io.ReadFull(r, payload)

	return
}
