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
	led.Off()

	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			led.On()
		} else {
			led.Off()
		}

		time.Sleep(time.Second)

	}
	led.Off()
}
