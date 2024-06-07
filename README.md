# go-vcnl40xx

go-vcnl40xx is an I2C driver for the Vishay VCNL40xx series of integrated
proximity and ambient light sensors.  


## Support Models

Currently it supports models [VCNL4040](https://www.vishay.com/docs/84274/vcnl4040.pdf).


## Usage

To use in your Go project, get the library.
```
go get github.com/swdee/go-vcnl40xx
```

## Example

To read proximity value.

```
// initialise sensor
sensor, _ := vcnl40xx.NewVCNL4040("/dev/i2c-0", vcnl40xx.VCNL4040Address)

proximity, _ := sensor.GetProximity()

fmt.Printf("Proximity = %d\n", proximity)
```
Note: Error handling has been skipped for brevity.


For reading Proximity, Ambient Light, White Light, and  setting Interrupts see 
the more [complete example here](example/). 



## Background

This code is ported from the [C library](https://github.com/sparkfun/SparkFun_VCNL4040_Arduino_Library).