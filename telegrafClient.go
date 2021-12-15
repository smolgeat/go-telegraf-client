package telegrafClient

import (
	"fmt"
	"net"
)

type client struct {
	host     string
	port     string
	tags     map[string]string
	protocol string
}

func (c client) write(metrics) {

	udpClient, err := net.ListenPacket("udp", c.host)
	if err != nil {

	} else {
		msg := []byte(metrics)
		udpClient.WriteTo(msg, c.host)
	}

}

func main() {

	fmt.Println("Test")
}
