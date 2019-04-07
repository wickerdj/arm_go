package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func off(l *gpio.LedDriver) {
	l.Brightness(0)
}

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")

	led.Start()
	off(led)

	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			for x := 0; x <= 255; x += 10 {
				time.Sleep(50 * time.Millisecond)

				setVal := byte(x)
				bright := led.Brightness(setVal)

				fmt.Printf("loop: %v | setVal: %v | bright: %v | \n", i, setVal, bright)
			}
		} else {

			off(led)
		}

		time.Sleep(time.Second)

	}

	off(led)
}
