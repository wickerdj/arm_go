package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	w1root   = "/sys/bus/w1/devices"
	w1master = "/w1_bus_master1/w1_master_slaves"
	dataFile = "/w1_slave"
)

var errReadingSensor = errors.New("failed to read sensor")

func getSensors() ([]string, error) {
	data, err := ioutil.ReadFile(filepath.Join(w1root, w1master))

	if err != nil {
		return nil, err
	}

	sensors := strings.Split(string(data), "\n")
	if len(sensors) > 0 {
		sensors = sensors[:len(sensors)-1]
	}

	return sensors, nil
}

func reading(sensor string) (float64, error) {
	f := filepath.Join(w1root, "/", sensor, dataFile)

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return 0.0, fmt.Errorf("Could not read sensor %s", f)
	}

	raw := string(data)

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		log.Println(raw)
		return 0.0, fmt.Errorf("Could not find tempature value")
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return 0.0, fmt.Errorf("Coulnd not convert response to float")
	}

	return c / 1000.0, nil
}

func main() {
	mySensors, err := getSensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sensor Ids: %v\n", mySensors)

	for _, sensor := range mySensors {
		t, err := reading(sensor)
		if err == nil {
			fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
		}
	}
}
