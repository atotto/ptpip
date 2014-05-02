// +build theta

// for RICOH THETA
// You can run `go test -run Theta -tags theta`
package ptpip_test

import (
	"log"
	"testing"

	"github.com/atotto/ptpip"
	"github.com/atotto/ptpip/ptp"
)

func TestThetaShutter(t *testing.T) {

	name := "Golang_Theta_Shutter"
	guid := "2adf2521-375c-406a-bc08-3c662406090e" // http://www.guidgenerator.com/online-guid-generator.aspx

	addr := "192.168.1.1:15740"

	config, err := ptpip.NewConfig(guid, name)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("connecting")
	c, err := ptpip.DialConfig(addr, config)
	if err != nil {
		t.Fatal(err)
	}

	sessionID := uint32(1)

	log.Println("open session")
	if err := c.OperationSimple(ptp.OC_OpenSession, 1, []uint32{sessionID}); err != nil {
		t.Fatal(err)
	}

	log.Println("initiate capture")
	if err := c.OperationSimple(ptp.OC_InitiateCapture, 2, []uint32{0, 0}); err != nil {
		t.Fatal(err)
	}

	log.Println("close session")
	if err := c.OperationSimple(ptp.OC_CloseSession, 3, []uint32{}); err != nil {
		t.Fatal(err)
	}
}
