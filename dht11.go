package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	host.Init()

	pin := gpioreg.ByName("GPIO4")
	pin.Out(gpio.High)

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		gcPercent := debug.SetGCPercent(-1)

		pin.Out(gpio.Low)
		time.Sleep(time.Millisecond)

		pin.In(gpio.PullUp, gpio.NoEdge)
		value := pin.Read()

		fmt.Printf("value: %v | type: %T | base16: % x | binary: %b \n", value, value, value, value)

		debug.SetGCPercent(gcPercent)
		time.Sleep(2 * time.Second)
	}

	pin.Out(gpio.High)
}
