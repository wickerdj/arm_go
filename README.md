# arm_go

This is a set of experiments writing go programs to run on a raspberry pi

## build command

GOARM=6 GOARCH=arm GOOS=linux go build {filename.go}

## GO Packages for working with the Raspberry Pi GPIO

gobot - https://gobot.io/documentation/platforms/raspi/

go-rpio - https://github.com/stianeikeland/go-rpio

## Notes on pin

When using the gobot library refer to https://pinout.xyz

## General Reference

RPi Low-level peripherals - https://elinux.org/RPi_Low-level_peripherals

Raspberry Pi on Wikipedia - https://en.wikipedia.org/wiki/Raspberry_Pi
