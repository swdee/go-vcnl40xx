package main

import (
	"flag"
	"fmt"
	"github.com/swdee/go-vcnl40xx"
	"log"
	"strconv"
	"time"
)

func main() {

	// read in cli flags
	model := flag.String("m", "4040", "Sensor model number [4040|4030|3035]")
	i2cbus := flag.String("b", "/dev/i2c-0", "Path to I2C bus to use")
	addr := flag.String("a", "", "Hex address of sensor on I2C bus")
	flag.Parse()

	var useModel vcnl40xx.Model
	var useAddr uint8

	switch *model {
	case "4030":
		useModel = vcnl40xx.VCNL4030
		useAddr = vcnl40xx.VCNL40301XAddress

	case "4035":
		useModel = vcnl40xx.VCNL4035
		useAddr = vcnl40xx.VCNL4035XAddress

	case "4040":
		useModel = vcnl40xx.VCNL4040
		useAddr = vcnl40xx.VCNL4040Address

	default:
		log.Fatalf("Unknown sensor model: %s\n", *model)
	}

	if *addr != "" {
		intValue, err := strconv.ParseInt(*addr, 0, 64)

		if err != nil || intValue > 255 {
			log.Fatalf("Error casting sensor hex address: %v", err)
		}

		useAddr = uint8(intValue)
	}

	sensor, err := vcnl40xx.NewSensor(useModel)

	if err != nil {
		log.Fatalf("Error creating sensor: %v\n", err)
	}

	err = sensor.Connect(*i2cbus, useAddr)

	if err != nil {
		log.Fatalf("Error connecting to sensor: %v\n", err)
	}

	err = sensor.Init()

	if err != nil {
		log.Fatalf("Error initializing sensor: %v\n", err)
	}

	configureInterrupt(sensor)
	readValues(sensor)
}

func readValues(sensor *vcnl40xx.Sensor) {

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

func configureInterrupt(sensor *vcnl40xx.Sensor) {

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
