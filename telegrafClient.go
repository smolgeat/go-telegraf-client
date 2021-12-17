package telegrafClient

import (
	"fmt"
	"net"
	//"encoding/json"
	// "encoding/gob"
)

type client struct {
	host     string
	port     string
	tags     map[string]string
	protocol string
}

func (c client) Write(metrics []byte) {

	udpClient, err := net.ListenPacket("udp", c.host)
	if err != nil {

	}
	udpClient.WriteTo(metrics, udpClient.LocalAddr())
	udpClient.Close()
}

func main() {

	fmt.Println("Test")
}
