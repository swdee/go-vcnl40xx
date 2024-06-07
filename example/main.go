package main

import (
	"fmt"
	"github.com/swdee/go-vcnl40xx"
	"log"
	"time"
)

const (
	// The I2C bus your sensor is connect on
	I2cDevice = "/dev/i2c-0"
)

func main() {

	sensor, err := vcnl40xx.NewVCNL4040(I2cDevice, vcnl40xx.VCNL4040Address)

	if err != nil {
		log.Fatalf("Error initializing driver: %v\n", err)
	}

	configureInterrupt(sensor)
	readValues(sensor)
}

func readValues(sensor *vcnl40xx.VCNL4040) {

	for {
		time.Sleep(1 * time.Second)

		// read proximity data
		proximity, err := sensor.GetProximity()

		if err != nil {
			log.Printf("Failed to read proximity: %v", err)
			continue
		}

		// read ambient light data
		ambient, err := sensor.GetAmbient()

		if err != nil {
			log.Printf("Failed to read ambient light: %v", err)
			continue
		}

		// read white light level
		white, err := sensor.GetWhite()

		if err != nil {
			log.Printf("Failed to white light: %v", err)
			continue
		}

		fmt.Printf("Proximity: %d, Ambient Light: %d, White Light: %d\n", proximity, ambient, white)
	}
}

func configureInterrupt(sensor *vcnl40xx.VCNL4040) {

	// if sensor sees a value higher than this, interrupt pin will go LOW
	err := sensor.SetProximityHighThreshold(2000)

	if err != nil {
		log.Fatalf("Failed to set proximity high threshold: %v", err)
	}

	// the interrupt pin will stay LOW until the value goes below the LOW
	// threshold value
	err = sensor.SetProximityLowThreshold(150)

	if err != nil {
		log.Fatalf("Failed to set proximity low threshold: %v", err)
	}

	// enable both 'away' and 'close' interrupts
	err = sensor.SetProximityInterruptType(vcnl40xx.InterruptBoth)

	if err != nil {
		log.Fatalf("Failed to set proximity interrupt type: %v", err)
	}

	// enables the proximity detection logic output mode
	// When this mode is selected, the INT pin is pulled LOW when an object is
	// close to the sensor (ie: value is above high threshold) and is reset to HIGH
	// when the object moves away (ie: value is below low threshold).
	// Get a multimeter and probe the INT pin to see this feature in action
	err = sensor.EnableProximityLogicMode()

	if err != nil {
		log.Fatalf("Failed to set proximity logic mode: %v", err)
	}
}
