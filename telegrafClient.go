package telegrafClient

//Package telegrafClient sends metric to telegraf socketlistener over udp in JSON
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

type client struct {
	host     string
	port     string
	tags     map[string]string
	protocol string
}

type metric struct {
	measurement map[string]string
	tags        map[string]string
}

func (c client) Write(metrics metric) {

	message := metrics.measurement

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
	udpClient, err := net.ListenPacket("udp", c.host+":"+c.port)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		n, err := udpClient.WriteTo(msgInBytes.Bytes(), udpClient.LocalAddr())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("success: %d\n", n)
		}
	}

}
