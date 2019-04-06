package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func threeFast(r *raspi.Adaptor, pin string) {
	led := gpio.NewLedDriver(r, pin)

	for i := 0; i < 3; i++ {
		led.On()
		time.Sleep(500 * time.Millisecond)
		led.Off()
		time.Sleep(200 * time.Millisecond)
	}

}

func blink(r *raspi.Adaptor, pin string) {
	led := gpio.NewLedDriver(r, pin)

	gobot.Every(1*time.Second, func() {
		led.Toggle()
	})

}

func main() {
	r := raspi.NewAdaptor()

	work := func() {
		go threeFast(r, "32")
		go blink(r, "32")
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{},
		work,
	)

	robot.Start()
}
