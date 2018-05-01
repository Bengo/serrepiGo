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
	state := flag.Int("s", 0, "GPIO state:  - 0 : low  - 1 : high")

	flag.Parse()

	if *gpioPort == 0 {
		fmt.Println("Veuillez specifier le port BCM gpio -n")
		os.Exit(1)
	}

	if *d == 0 {
		fmt.Println("Veuillez specifier une duree d'activation -d")
		os.Exit(2)
	}

	if *state != 0 && *state != 1 {
		fmt.Println("Veuillez specifier un état valide  0 ou 1 -s")
		os.Exit(3)
	}

	fmt.Printf("Activation du port gpio BCM %d pour %s a létat %d \n", *gpioPort, *d, *state)
	err := rpio.Open()
	if err != nil {
		fmt.Print(err)
	}
	pin := rpio.Pin(*gpioPort)
	pin.Output() // Output mode

	if *state == 0 {
		pin.Low()
	} else {
		pin.High()
	}
	time.Sleep(*d)
	pin.Toggle()
	pin.Input()
	rpio.Close()
	fmt.Printf("Le port gpio BCM %d a ete active %s a l'état %d \n", *gpioPort, *d, *state)
	os.Exit(0)
}
