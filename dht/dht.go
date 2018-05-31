package serrepi

import (
	"fmt"
	"log"

	dht "github.com/d2r2/go-dht"
)

//Data  dht2302 data
type Data struct {
	temperature float32
	humidity    float32
}

//read humidity and temperature from DHT22 onport 26
func readDht() *Data {

	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT22, 26, true, 10)
	if err != nil {
		log.Fatal(err)
	}

	// Print temperature and humidity
	fmt.Printf("Temperature = %v°C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
	data := new(Data)
	data.humidity = humidity
	data.temperature = temperature
	return data
}

func main() {
	readDht()
}
