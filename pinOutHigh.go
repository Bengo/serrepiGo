package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	rpio "github.com/go-rpio"
)

func main() {
	gpioPort := flag.Int("n", 0, "BCM Gpio port number")
	d := flag.Duration("d", 0, "duration")
	flag.Parse()

	if *gpioPort == 0 {
		fmt.Println("Veuillez specifier le port BCM gpio -n")
		os.Exit(1)
	}

	if *d == 0 {
		fmt.Println("Veuillez specifier une duree d'activation -d")
		os.Exit(2)
	}

	fmt.Printf("Activation du port gpio BCM %d pour %s \n", *gpioPort, *d)
	err := rpio.Open()
	if err != nil {
		fmt.Print(err)
	}
	pin := rpio.Pin(*gpioPort)
	pin.Output() // Output mode
	pin.High()   // Set pin High
	time.Sleep(*d)
	pin.Low()
	rpio.Close()
	fmt.Printf("Le port gpio BCM %d a ete active %s \n", *gpioPort, *d)
	os.Exit(0)
}