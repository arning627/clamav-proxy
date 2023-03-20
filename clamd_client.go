package clamd

type clamd struct {
	host    string
	port    uint
	timeout int
}

func (c clamd) NewClient(host string, port uint, timeout int) *clamd {
	if port < 0 || port > 65535 {
		panic("Port exceeds limit")
	}
	clamd := clamd{host: host, port: port, timeout: timeout}
	return &clamd
}

// clamd PING command
func (c *clamd) Ping() string {

	return "pang"
}

func (c *clamd) Sacn() bool {
	return false
}
