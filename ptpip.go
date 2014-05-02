// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// refs :
//
// - http://sourceforge.net/p/gphoto/code/HEAD/tree/trunk/libgphoto2/camlibs/ptp2/PTPIP.TXT
// - http://www.cipa.jp/ptp-ip/documents_j/CIPA_DC-005_Whitepaper_JPN.pdf
package ptpip

import (
	"fmt"
	"net"

	"github.com/atotto/ptpip/packet"
	"github.com/atotto/ptpip/ptp"
)

type Client struct {
	commandConn net.Conn
	eventConn   net.Conn
	config      *Config
	responder   Responder
}

type ErrInvalidArgs struct {
	message string
}

//
func (e *ErrInvalidArgs) Error() string {
	return fmt.Sprintf("invalid argument: %s", e.message)
}

type Config struct {
	GUID         string
	FriendlyName string
}

type Responder struct {
	Addr         string
	GUID         string
	FriendlyName string
}

// NewConfig creates a new PTP-IP config for client connection.
// If guid argument less than 16, return ErrInvalidArgs.
// If friendlyName argument less than 20, return ErrInvalidArgs.
func NewConfig(guid, friendlyName string) (config *Config, err error) {
	config = new(Config)

	if len(guid) < 16 {
		return nil, &ErrInvalidArgs{message: fmt.Sprintf("guid length must 16, got %d", len(guid))}
	}
	config.GUID = guid

	if len(friendlyName) < 20 {
		return nil, &ErrInvalidArgs{message: fmt.Sprintf("name length should be less than 20, got %d", len(friendlyName))}
	}
	config.FriendlyName = friendlyName

	return
}

// Dial dial PTP-IP supported device.
// Default PTP-IP port number is 15740.
func DialConfig(addr string, config *Config) (c *Client, err error) {
	c = new(Client)
	c.config = config

	err = c.initConnection(addr)

	return
}

//
func (c *Client) initConnection(addr string) (err error) {
	// CommandDataConnection
	commandConn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	if err = packet.SendInitCommand(commandConn, c.config.GUID, c.config.FriendlyName); err != nil {
		return
	}
	sessionID, guid, friendlyName, err := packet.RecvInitCommand(commandConn)
	if err != nil {
		return
	}

	c.commandConn = commandConn
	c.responder.GUID = guid
	c.responder.FriendlyName = friendlyName

	// EventConnection
	eventConn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	if err = packet.SendInitEvent(eventConn, sessionID); err != nil {
		return
	}
	if err = packet.RecvInitEvent(eventConn); err != nil {
		return
	}
	c.eventConn = eventConn

	return
}

const (
	kUNKNOWN_DATA_PHASE       uint32 = 0x00000000
	kNO_DATA_OR_DATA_IN_PHASE uint32 = 0x00000001
	kDATA_OUT_PHASE           uint32 = 0x00000002
)

func (c *Client) OperationSimple(operationCode ptp.OperationCode, transactionID uint32, parameters []uint32) (err error) {
	dataPheseInfo := kNO_DATA_OR_DATA_IN_PHASE
	if err = packet.SendOperationRequest(c.commandConn, dataPheseInfo, operationCode, transactionID, parameters); err != nil {
		return
	}
	_, _, _, err = packet.RecvOperationResponse(c.commandConn)
	return
}
