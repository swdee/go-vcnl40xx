# go-vcnl40xx

go-vcnl40xx is an I2C driver for the Vishay VCNL40xx series of integrated
proximity and ambient light sensors.  

Currently it supports models [VCNL4040](https://www.vishay.com/docs/84274/vcnl4040.pdf), 
[VCNL3030](https://www.vishay.com/docs/84250/vcnl4030x01.pdf), and 
[VCNL4035](https://www.vishay.com/docs/84251/vcnl4035x01.pdf).


## Usage

To use in your Go project, get the library.
```
go get github.com/swdee/go-vcnl40xx
```

## Example

To read proximity value on VCNL4040 sensor.

```
// initialise sensor
sensor, _ := vcnl40xx.NewSensor(vcnl40xx.VCNL4040)

// connect to sensor and initialize
_ = sensor.Connect("/dev/i2c-0", vcnl40xx.VCNL4040Address)
_ = sensor.Init()

proximity, _ := sensor.GetProximity()

fmt.Printf("Proximity = %d\n", proximity)
```
Note: Error handling has been skipped for brevity.


For reading Proximity, Ambient Light, White Light, and setting Interrupts see 
the more [complete example here](example/main.go). 



## Background

This code is based on the [C library](https://github.com/sparkfun/SparkFun_VCNL4040_Arduino_Library).