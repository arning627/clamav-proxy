package clamd

import (
	"fmt"
	"io"
	"net"
)

type clamd struct {
	host string
	port uint
	// _timeout int
}

func NewClient(host string, port uint, timeout int) *clamd {
	if port < 0 || port > 65535 {
		panic("Port exceeds limit")
	}
	clamd := clamd{host: host, port: port}
	return &clamd
}

// clamd PING command
func (this *clamd) Ping() (string, error) {
	//ping
	conn, netErr := this.tcpConnection()
	defer conn.Close()
	if netErr != nil {
		return "", netErr
	}
	_, writeErr := conn.Write([]byte("PING"))
	if writeErr != nil {
		return "", writeErr
	}
	return "PANG", nil
}

func (this *clamd) SacnStream(r io.Reader) (bool, error) {
	// scan
	conn, netErr := this.tcpConnection()
	defer conn.Close()
	if netErr != nil {
		return false, netErr
	}
	for {
		//TODO
		buffer := make([]byte, 2048)
		length, readErr := r.Read(buffer)
		if length > 0 {
			conn.Write(buffer)
		}
		if readErr != nil {
			break
		}
	}
	//TODO get response
	return false, nil
}

func (this *clamd) tcpConnection() (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprint(this.host, ":", this.port))
}
