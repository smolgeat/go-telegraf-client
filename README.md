# Telegraf Client
- Go client for sending metrics to telegraf

- Designed to work with telegraf [socket listener plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/socket_listener) 

- Sends data in JSON format


Example using global tags
```
import (
	telegrafClient "github.com/smolgeat/go-telegraf-client"
)

func main() {

	var UDPClient telegrafClient.Client
	var metrics telegrafClient.Metric

	UDPClient.Tags = make(map[string]string)
	UDPClient.Tags["my_tag_1"] = "tag1"
	UDPClient.Tags["my_tag_2"] = "tag2"
	UDPClient.Server = "127.0.0.1:8094"
	metrics.Measurement = make(map[string]string)
	metrics.Measurement["a"] = "5"
	metrics.Measurement["b"] = "6"
	for {
		UDPClient.Write(metrics)
	}
}
```

Example using metric tag instead of global tag 
**NB** tag3 is sent with metrics instead of tag1 and tag1
```
import (
	telegrafClient "github.com/smolgeat/go-telegraf-client"
)

func main() {

	var UDPClient telegrafClient.Client
	var metrics telegrafClient.Metric

	UDPClient.Tags = make(map[string]string)
	UDPClient.Tags["my_tag_1"] = "tag1"
	UDPClient.Tags["my_tag_2"] = "tag2"
	UDPClient.Server = "127.0.0.1:8094"
  metrics.Tags = make(map[string]string)
  metrics.Tags["my_tag_3"] = "tag3"
	metrics.Measurement = make(map[string]string)
	metrics.Measurement["a"] = "5"
	metrics.Measurement["b"] = "6"
	for {
		UDPClient.Write(metrics)
	}
}
```
