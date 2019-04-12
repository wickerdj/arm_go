package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func threeFast(led *gpio.LedDriver, delay time.Duration) {
	for i := 0; i < 3; i++ {
		led.On()
		time.Sleep(delay)
		led.Off()
		time.Sleep(delay)
	}
	time.Sleep(2 * delay)
}

func blink(led *gpio.LedDriver) {
	gobot.Every(time.Second, func() {
		led.Toggle()
	})
}

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "12")

	work := func() {
		delay, _ := time.ParseDuration("200ms")
		threeFast(led, delay)
		blink(led)
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{},
		work,
	)

	robot.Start()
}
