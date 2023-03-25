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
	conn.Write([]byte("nINSTREAM\n"))
	for {
		buffer := make([]byte, 2048)
		length, readErr := r.Read(buffer)
		if length > 0 {
			send(buffer, conn)
		}
		if readErr != nil {
			log.Println(readErr)
			break
		}
	}
	log.Println("write end 0,0,0,0")
	_, e := conn.Write([]byte{0, 0, 0, 0})
	if e != nil {
		log.Println(e)
	}

	responseReader := bufio.NewReader(conn)
	log.Println("read response")
	for {
		log.Println("reading response...")
		line, e := responseReader.ReadString('\n')
		log.Println("line====", line)
		if e == io.EOF {
			break
		}
		if e != nil {
			break
		}
	}
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
		if e != nil {
			break
		}
	}

}

func (c *clamd) tcpConnection() (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprint(c.host, ":", c.port))
}

// https://github.com/dutchcoders/go-clamd/blob/master/conn.go 直接写入会抛出 INSTREAM size limit exceeded. ERROR
// 使用这个文件中的方法可以正常检测 具体还不清楚为何要这样写
func send(data []byte, conn net.Conn) error {
	var buf [4]byte
	lenData := len(data)
	buf[0] = byte(lenData >> 24)
	buf[1] = byte(lenData >> 16)
	buf[2] = byte(lenData >> 8)
	buf[3] = byte(lenData >> 0)

	a := buf

	b := make([]byte, len(a))
	for i := range a {
		b[i] = a[i]
	}

	conn.Write(b)

	_, err := conn.Write(data)
	return err
}
