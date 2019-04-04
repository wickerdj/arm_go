package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func blink(r *raspi.Adaptor, pin string) {
	led := gpio.NewLedDriver(r, pin)

	go gobot.Every(1*time.Second, func() {
		led.Toggle()
		fmt.Println("switch")
	})

}

func main() {
	r := raspi.NewAdaptor()

	work := func() {
		go blink(r, "12")
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{},
		work,
	)

	robot.Start()
}
