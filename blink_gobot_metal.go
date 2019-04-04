package main

import (
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")

	led.Start()

	for {
		led.Toggle()
		time.Sleep(1 * time.Second)
	}
}
