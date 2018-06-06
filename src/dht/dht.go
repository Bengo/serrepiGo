package dht

import (
	"log"
	"time"

	r2d2 "github.com/d2r2/go-dht"
)

//Data  dht2302 data
type Data struct {
	Temperature float32
	Humidity    float32
	Time        string
}

//ReadDht read humidity and temperature from DHT22 onport 26
func ReadDht() *Data {

	temperature, humidity, _, err :=
		r2d2.ReadDHTxxWithRetry(r2d2.DHT22, 26, true, 10)
	if err != nil {
		log.Fatal(err)
	}
	data := new(Data)
	data.Humidity = humidity
	data.Temperature = temperature
	data.Time = time.Now().Format(time.RFC3339)
	return data
}
