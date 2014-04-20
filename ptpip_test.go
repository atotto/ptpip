package ptpip_test

import (
	"io"
	"log"
	"net"
	"testing"
)

func testEchoServer(t *testing.T) {
	l, err := net.Listen("tcp", ":62000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}

func TestDialPTPIP(t *testing.T) {
	// _, err := ptpip.DialPTPIP("localhost:15740", "AAA", "golang ptpip client")
	// if err != nil {
	// 	t.Fatal(err)
	// }
}
