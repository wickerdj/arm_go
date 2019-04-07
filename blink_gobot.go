package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func threeFast(led *gpio.LedDriver) {
	for i := 0; i < 3; i++ {
		led.On()
		time.Sleep(500 * time.Millisecond)
		led.Off()
		time.Sleep(200 * time.Millisecond)
	}

}

func blink(led *gpio.LedDriver) {
	gobot.Every(time.Second, func() {
		led.Toggle()
	})
}

func off(led *gpio.LedDriver) {
	led.Off()
}

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")

	work := func() {
		off(led)
		threeFast(led)
		blink(led)
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{},
		work,
	)

	robot.Start()
}
