package telegrafClient

//Package telegrafClient sends metric to telegraf socketlistener over udp in JSON
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

type Client struct {
	Server   string
	Tags     map[string]string
	Protocol string
}

type Metric struct {
	Measurement map[string]string
	Tags        map[string]string
}

func (c Client) Write(metrics Metric) []byte {
	// Writes metrics to telegraf will use global Tags if metric Tags arent specified
	//
	//
	//
	message := metrics.Measurement
	serverAddr, err := net.ResolveUDPAddr("udp", c.Server)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if metrics.Tags != nil {
		for key, value := range metrics.Tags {
			message[key] = value
		}
	} else {
		for key, value := range c.Tags {
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
