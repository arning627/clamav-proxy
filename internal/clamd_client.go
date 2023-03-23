package clamd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

type clamd struct {
	host string
	port string
	// _timeout int
}

func NewClient(host string, port string) *clamd {
	if p, e := strconv.ParseUint(port, 0, 0); e == nil {
		if p < 0 || p > 65535 {
			panic("Port exceeds limit")
		}
	} else {
		panic("Port exceeds limit")
	}

	clamd := clamd{host: host, port: port}
	return &clamd
}

// clamd PING command
func (c *clamd) Ping() (string, error) {
	//ping
	conn, netErr := c.tcpConnection()
	defer conn.Close()
	if netErr != nil {
		return "", netErr
	}
	_, writeErr := conn.Write([]byte("PING"))
	if writeErr != nil {
		return "", writeErr
	}

	reader := bufio.NewReader(conn)

	for {
		line, e := reader.ReadString('\n')
		log.Println("line====", line)
		if e == io.EOF {
			break
		}
		if e == nil {
			break
		}
	}

	return "PONG", nil
}

func (c *clamd) SacnStream(r io.Reader) (bool, error) {
	// scan
	conn, netErr := c.tcpConnection()
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

	reader := bufio.NewReader(conn)

	for {
		line, e := reader.ReadString('\n')
		log.Println("line====", line)
		if e == io.EOF {
			break
		}
		if e == nil {
			break
		}
	}
	//TODO get response
	return false, nil
}

func (c *clamd) Execute(command string) {
	conn, netErr := c.tcpConnection()
	defer conn.Close()
	if netErr != nil {
		log.Fatalln(netErr)
	}
	_, writeErr := conn.Write([]byte(command))
	if writeErr != nil {
		log.Fatalln(writeErr)
	}

	reader := bufio.NewReader(conn)

	for {
		line, e := reader.ReadString('\n')
		log.Println("line====", line)
		if e == io.EOF {
			break
		}
		if e == nil {
			break
		}
	}

}

func (c *clamd) tcpConnection() (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprint(c.host, ":", c.port))
}
