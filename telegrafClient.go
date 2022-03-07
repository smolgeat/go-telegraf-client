package telegrafClient

//Package telegrafClient sends metric to telegraf socketlistener over udp in JSON
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

type client struct {
	server   string
	tags     map[string]string
	protocol string
}

type metric struct {
	measurement map[string]string
	tags        map[string]string
}

func (c client) Write(metrics metric) []byte {

	message := metrics.measurement
	serverAddr, err := net.ResolveUDPAddr("udp", c.server)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if metrics.tags != nil {
		for key, value := range metrics.tags {
			message[key] = value
		}
	} else {
		for key, value := range c.tags {
			message[key] = value
		}
	}
	msgInBytes := new(bytes.Buffer)
	json.NewEncoder(msgInBytes).Encode(message)
	udpClient, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		_, err := udpClient.WriteTo(msgInBytes.Bytes(), serverAddr)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
	return msgInBytes.Bytes()
}
