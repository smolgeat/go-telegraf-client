package telegrafClient

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"testing"
)

func TestDialUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := UDPServer(ctx, "127.0.0.1:8000")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	client, err := net.Dial("udp", serverAddr.String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

}

func UDPServer(ctx context.Context, addr string) (net.Addr, error) {
	server, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("udp server failed %s %w", addr, err)
	}
	go func() {
		go func() {
			<-ctx.Done()
			_ = server.Close

		}()

		buf := make([]byte, 1024)

		for {
			n, clientAddr, err := server.ReadFrom(buf)
			if err != nil {
				return

			}

			_, err = server.WriteTo(buf[:n], clientAddr)
			if err != nil {
				return
			}
		}
	}()

	return server.LocalAddr(), nil
}

func TestWrite(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := UDPServer(ctx, "127.0.0.1:8000")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	var UDPClient client
	var metrics metric

	UDPClient.tags = make(map[string]string)
	UDPClient.tags["my_tag_1"] = "foo"
	UDPClient.tags["my_tag_2"] = "bar"
	UDPClient.server = serverAddr.String()
	metrics.measurement = make(map[string]string)
	metrics.measurement["a"] = "5"
	metrics.measurement["b"] = "6"
	expected := []byte(`{"a":"5","b":"6","my_tag_1":"foo","my_tag_2":"bar"}`)
	result := UDPClient.Write(metrics)
	resultLessNL := bytes.TrimRight(result, "\n")

	if !bytes.Equal(expected, resultLessNL) {
		t.Fatal()
	}

}
