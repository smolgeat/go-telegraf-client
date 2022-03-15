/*Package telegrafClient sends metric to telegraf socketlistener over udp in JSON



 */

package telegrafClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

type Client struct {
	Server   string            //Server IP and port
	Tags     map[string]string //Global Tags
	Protocol string            //Server Protocol
}

type Metric struct {
	Measurement map[string]string // Metric Values
	Tags        map[string]string // Metric Tags
}

//Write sends metrics to telegraf,
//will use global Tags if metric Tags arent specified
func (c Client) Write(metrics Metric) []byte {

	var result []byte

	if c.Protocol == "UDP" {

		result = c.WriteUDP(metrics)

	} else if c.Protocol == "TCP" {

		result = c.WriteTCP(metrics)

	} else {
		fmt.Printf("No protocol specified")
	}

	return result

}

//WriteUDP creates UDP client then sends metrics as bytes to UDP server
func (c Client) WriteUDP(metrics Metric) []byte {

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

func (c Client) WriteTCP(metrics Metric) []byte {

	placeholder := make(map[string]string)
	msgInBytes := new(bytes.Buffer)
	json.NewEncoder(msgInBytes).Encode(placeholder)

	return msgInBytes.Bytes()
}
