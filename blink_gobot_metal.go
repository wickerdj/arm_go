package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "32")

	led.Start()
	led.Off()

	for i := 0; i < 3; i++ {
		// led.Toggle()
		// fmt.Println("Led State:", led.State())

		if i%2 == 0 {
			// led.On()
			for x := 0; x <= 255; x++ {
				time.Sleep(50 * time.Millisecond)

				setVal := byte(x)
				bright := led.Brightness(setVal)

				fmt.Printf("loop: %v | setVal: %v | bright: %v | \n", i, setVal, bright)
			}
		} else {
			led.Off()
		}

		time.Sleep(time.Second)

	}
	led.Off()
}
