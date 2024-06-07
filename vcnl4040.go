package vcnl40xx

// converted C code from https://github.com/sparkfun/SparkFun_VCNL4040_Arduino_Library/blob/master/src/SparkFun_VCNL4040_Arduino_Library.cpp

import (
	"fmt"
	"github.com/swdee/go-i2c"
)

const (
	VCNL4040Address  = 0x60
	VCNL4040SensorID = 0x0186

	LOWER = true
	UPPER = false

	// VCNL4040 Command Codes
	VCNL4040_ALS_CONF   = 0x00
	VCNL4040_ALS_THDH   = 0x01
	VCNL4040_ALS_THDL   = 0x02
	VCNL4040_PS_CONF1   = 0x03 //Lower
	VCNL4040_PS_CONF2   = 0x03 //Upper
	VCNL4040_PS_CONF3   = 0x04 //Lower
	VCNL4040_PS_MS      = 0x04 //Upper
	VCNL4040_PS_CANC    = 0x05
	VCNL4040_PS_THDL    = 0x06
	VCNL4040_PS_THDH    = 0x07
	VCNL4040_PS_DATA    = 0x08
	VCNL4040_ALS_DATA   = 0x09
	VCNL4040_WHITE_DATA = 0x0A
	VCNL4040_INT_FLAG   = 0x0B //Upper
	VCNL4040_ID         = 0x0C

	// .h file below
	VCNL4040_ALS_IT_MASK  uint8 = ^uint8((1 << 7) | (1 << 6))
	VCNL4040_ALS_IT_80MS  uint8 = 0
	VCNL4040_ALS_IT_160MS uint8 = 1 << 7
	VCNL4040_ALS_IT_320MS uint8 = 1 << 6
	VCNL4040_ALS_IT_640MS uint8 = (1 << 7) | (1 << 6)

	VCNL4040_ALS_PERS_MASK uint8 = ^uint8((1 << 3) | (1 << 2))
	VCNL4040_ALS_PERS_1    uint8 = 0
	VCNL4040_ALS_PERS_2    uint8 = 1 << 2
	VCNL4040_ALS_PERS_4    uint8 = 1 << 3
	VCNL4040_ALS_PERS_8    uint8 = (1 << 3) | (1 << 2)

	VCNL4040_ALS_INT_EN_MASK uint8 = ^uint8(1 << 1)
	VCNL4040_ALS_INT_DISABLE uint8 = 0
	VCNL4040_ALS_INT_ENABLE  uint8 = 1 << 1

	VCNL4040_ALS_SD_MASK      uint8 = ^uint8(1 << 0)
	VCNL4040_ALS_SD_POWER_ON  uint8 = 0
	VCNL4040_ALS_SD_POWER_OFF uint8 = 1 << 0

	VCNL4040_PS_DUTY_MASK uint8 = ^uint8((1 << 7) | (1 << 6))
	VCNL4040_PS_DUTY_40   uint8 = 0
	VCNL4040_PS_DUTY_80   uint8 = (1 << 6)
	VCNL4040_PS_DUTY_160  uint8 = (1 << 7)
	VCNL4040_PS_DUTY_320  uint8 = (1 << 7) | (1 << 6)

	VCNL4040_PS_PERS_MASK uint8 = ^uint8((1 << 5) | (1 << 4))
	VCNL4040_PS_PERS_1    uint8 = 0
	VCNL4040_PS_PERS_2    uint8 = 1 << 4
	VCNL4040_PS_PERS_3    uint8 = 1 << 5
	VCNL4040_PS_PERS_4    uint8 = (1 << 5) | (1 << 4)

	VCNL4040_PS_IT_MASK uint8 = ^uint8((1 << 3) | (1 << 2) | (1 << 1))
	VCNL4040_PS_IT_1T   uint8 = 0
	VCNL4040_PS_IT_15T  uint8 = (1 << 1)
	VCNL4040_PS_IT_2T   uint8 = (1 << 2)
	VCNL4040_PS_IT_25T  uint8 = (1 << 2) | (1 << 1)
	VCNL4040_PS_IT_3T   uint8 = (1 << 3)
	VCNL4040_PS_IT_35T  uint8 = (1 << 3) | (1 << 1)
	VCNL4040_PS_IT_4T   uint8 = (1 << 3) | (1 << 2)
	VCNL4040_PS_IT_8T   uint8 = (1 << 3) | (1 << 2) | (1 << 1)

	VCNL4040_PS_SD_MASK      uint8 = ^uint8(1 << 0)
	VCNL4040_PS_SD_POWER_ON  uint8 = 0
	VCNL4040_PS_SD_POWER_OFF uint8 = 1 << 0

	VCNL4040_PS_HD_MASK   uint8 = ^uint8(1 << 3)
	VCNL4040_PS_HD_12_BIT uint8 = 0
	VCNL4040_PS_HD_16_BIT uint8 = 1 << 3

	VCNL4040_PS_INT_MASK    uint8 = ^uint8((1 << 1) | (1 << 0))
	VCNL4040_PS_INT_DISABLE uint8 = 0
	VCNL4040_PS_INT_CLOSE   uint8 = 1 << 0
	VCNL4040_PS_INT_AWAY    uint8 = 1 << 1
	VCNL4040_PS_INT_BOTH    uint8 = (1 << 1) | (1 << 0)

	VCNL4040_PS_SMART_PERS_MASK    uint8 = ^uint8(1 << 4)
	VCNL4040_PS_SMART_PERS_DISABLE uint8 = 0
	VCNL4040_PS_SMART_PERS_ENABLE  uint8 = 1 << 4

	VCNL4040_PS_AF_MASK    uint8 = ^uint8(1 << 3)
	VCNL4040_PS_AF_DISABLE uint8 = 0
	VCNL4040_PS_AF_ENABLE  uint8 = 1 << 3

	VCNL4040_PS_TRIG_MASK    uint8 = ^uint8(1 << 2)
	VCNL4040_PS_TRIG_TRIGGER uint8 = 1 << 2

	VCNL4040_WHITE_EN_MASK uint8 = ^uint8(1 << 7)
	VCNL4040_WHITE_ENABLE  uint8 = 0
	VCNL4040_WHITE_DISABLE uint8 = 1 << 7

	VCNL4040_PS_MS_MASK    uint8 = ^uint8(1 << 6)
	VCNL4040_PS_MS_DISABLE uint8 = 0
	VCNL4040_PS_MS_ENABLE  uint8 = (1 << 6)

	VCNL4040_LED_I_MASK uint8 = ^uint8((1 << 2) | (1 << 1) | (1 << 0))
	VCNL4040_LED_50MA   uint8 = 0
	VCNL4040_LED_75MA   uint8 = (1 << 0)
	VCNL4040_LED_100MA  uint8 = (1 << 1)
	VCNL4040_LED_120MA  uint8 = (1 << 1) | (1 << 0)
	VCNL4040_LED_140MA  uint8 = (1 << 2)
	VCNL4040_LED_160MA  uint8 = (1 << 2) | (1 << 0)
	VCNL4040_LED_180MA  uint8 = (1 << 2) | (1 << 1)
	VCNL4040_LED_200MA  uint8 = (1 << 2) | (1 << 1) | (1 << 0)

	VCNL4040_INT_FLAG_ALS_LOW  uint8 = 1 << 5
	VCNL4040_INT_FLAG_ALS_HIGH uint8 = 1 << 4
	VCNL4040_INT_FLAG_CLOSE    uint8 = 1 << 1
	VCNL4040_INT_FLAG_AWAY     uint8 = 1 << 0
)

// ProximityPersistance defines the proximity types
type ProximityPersistance uint8

const (
	ProximityPersistance1 ProximityPersistance = 1
	ProximityPersistance2 ProximityPersistance = 2
	ProximityPersistance3 ProximityPersistance = 3
	ProximityPersistance4 ProximityPersistance = 4
)

// ProximityPersistance defines the ambient persistance types
type AmbientPersistance uint8

const (
	AmbientPersistance1 AmbientPersistance = 1
	AmbientPersistance2 AmbientPersistance = 2
	AmbientPersistance4 AmbientPersistance = 4
	AmbientPersistance8 AmbientPersistance = 8
)

// InterruptType defines the interrupt modes the sensor can be set too
type InterruptType uint8

const (
	InterruptDisable InterruptType = 1
	InterruptClose   InterruptType = 2
	InterruptAway    InterruptType = 3
	InterruptBoth    InterruptType = 4
)

// VCNL4040 defines the VCNL4040 sensor device
type VCNL4040 struct {
	i2c *i2c.Options
}

// NewVCNL4040 returns a driver instance for the VCNL4040 on the given I2C bus
// and address of the sensor.
// The proximity and ambient light sensors are initialized and powered on.  If
// either function is not required then power off those sensors
func NewVCNL4040(dev string, addr uint8) (*VCNL4040, error) {

	i2c, err := i2c.New(addr, dev)

	if err != nil {
		return nil, fmt.Errorf("i2c bus error: %w", err)
	}

	check := i2c.GetAddr()

	if check == 0 {
		return nil, fmt.Errorf("I2C device is not initiated")
	}

	v := &VCNL4040{
		i2c: i2c,
	}

	id, err := v.GetID()

	if err != nil {
		return nil, fmt.Errorf("error getting sensor ID: %w", err)
	}

	if id != VCNL4040SensorID {
		return nil, fmt.Errorf("unexpected sensor ID value: 0x%04X, wanted 0x%04X", id, VCNL4040SensorID)

	}

	err = v.Init()

	if err != nil {
		return nil, fmt.Errorf("error initialising sensor: %w", err)
	}

	return v, nil
}

// Init initialises the sensor and puts it in default state with proximity
// and ambient light sensors activated
func (v *VCNL4040) Init() error {

	if err := v.SetLEDCurrent(200); err != nil {
		return fmt.Errorf("error setting LED current: %w", err)
	}

	if err := v.SetIRDutyCycle(40); err != nil {
		return fmt.Errorf("error setting IR duty cycle: %w", err)
	}

	if err := v.SetProximityIntegrationTime(8); err != nil {
		return fmt.Errorf("error setting proximity integration time: %w", err)
	}

	if err := v.SetProximityResolution(16); err != nil {
		return fmt.Errorf("error setting proximity resolution: %w", err)
	}

	if err := v.EnableSmartPersistance(); err != nil {
		return fmt.Errorf("error enabling smart persistance: %w", err)
	}

	if err := v.PowerOnProximity(); err != nil {
		return fmt.Errorf("error powering on proximity function: %w", err)
	}

	if err := v.SetAmbientIntegrationTime(80); err != nil {
		return fmt.Errorf("error setting ambient integration time: %w", err)
	}

	if err := v.PowerOnAmbient(); err != nil {
		return fmt.Errorf("error powering on ambient lighting function: %w", err)
	}

	return nil
}

// PowerOnAmbient turns on the ambient lighting sensor of the device
func (v *VCNL4040) PowerOnAmbient() error {
	return v.bitMask(VCNL4040_ALS_CONF, LOWER, VCNL4040_ALS_SD_MASK, VCNL4040_ALS_SD_POWER_ON)
}

// PowerOffAmbient turns off the ambient lighting sensor of the device
func (v *VCNL4040) PowerOffAmbient() error {
	return v.bitMask(VCNL4040_ALS_CONF, LOWER, VCNL4040_ALS_SD_MASK, VCNL4040_ALS_SD_POWER_OFF)
}

// SetAmbientIntegrationTime sets the integration time for the ambient light
// sensor. valid values are 80, 160, 320, or 640 which represents the number
// of milliseconds
func (v *VCNL4040) SetAmbientIntegrationTime(timeValue uint16) error {

	if timeValue >= 640 {
		timeValue = uint16(VCNL4040_ALS_IT_640MS)
	} else if timeValue >= 320 {
		timeValue = uint16(VCNL4040_ALS_IT_320MS)
	} else if timeValue >= 160 {
		timeValue = uint16(VCNL4040_ALS_IT_160MS)
	} else {
		timeValue = uint16(VCNL4040_ALS_IT_80MS)
	}

	return v.bitMask(VCNL4040_ALS_CONF, LOWER, VCNL4040_ALS_IT_MASK, byte(timeValue))
}

// PowerOnProximity turns on the proximity sensor of the device
func (v *VCNL4040) PowerOnProximity() error {
	return v.bitMask(VCNL4040_PS_CONF1, LOWER, VCNL4040_PS_SD_MASK, VCNL4040_PS_SD_POWER_ON)
}

// PowerOffProximity turns off the proximity sensor of the device
func (v *VCNL4040) PowerOffProximity() error {
	return v.bitMask(VCNL4040_PS_CONF1, LOWER, VCNL4040_PS_SD_MASK, VCNL4040_PS_SD_POWER_OFF)
}

// EnableSmartPersistance to accelerate the PS response time, smart
// persistence prevents the misjudgment of proximity sensing but also keeps
// a fast response time.
func (v *VCNL4040) EnableSmartPersistance() error {
	return v.bitMask(VCNL4040_PS_CONF3, LOWER, VCNL4040_PS_SMART_PERS_MASK, VCNL4040_PS_SMART_PERS_ENABLE)
}

// DisableSmartPersistence disable smart persistence
func (v *VCNL4040) DisableSmartPersistence() error {
	return v.bitMask(VCNL4040_PS_CONF3, LOWER, VCNL4040_PS_SMART_PERS_MASK, VCNL4040_PS_SMART_PERS_DISABLE)
}

// SetProximityResolution sets the proximity resolution to either 16 or 12 bit.
// valid values are 12 or 16.
func (v *VCNL4040) SetProximityResolution(resolutionValue uint8) error {

	if resolutionValue >= 16 {
		resolutionValue = VCNL4040_PS_HD_16_BIT
	} else {
		resolutionValue = VCNL4040_PS_HD_12_BIT
	}

	return v.bitMask(VCNL4040_PS_CONF2, UPPER, VCNL4040_PS_HD_MASK, resolutionValue)
}

// SetProximityIntegrationTime sets the integration time for the proximity sensor
// which represents the duration of the energy being received. valid values
// are 1, 2, 3, 4, or 8.
func (v *VCNL4040) SetProximityIntegrationTime(timeValue uint8) error {

	if timeValue >= 8 {
		timeValue = VCNL4040_PS_IT_8T
	} else if timeValue >= 4 {
		timeValue = VCNL4040_PS_IT_4T
	} else if timeValue >= 3 {
		timeValue = VCNL4040_PS_IT_3T
	} else if timeValue >= 2 {
		timeValue = VCNL4040_PS_IT_2T
	} else {
		timeValue = VCNL4040_PS_IT_1T
	}

	return v.bitMask(VCNL4040_PS_CONF1, LOWER, VCNL4040_PS_IT_MASK, timeValue)
}

// SetIRDutyCycle sets the duty cycle of the IR LED. The higher the duty
// ratio, the faster the response time achieved with higher power
// consumption. For example, PS_Duty = 1/320, peak IRED current = 100 mA,
// averaged current consumption is 100 mA/320 = 0.3125 mA.
// valid values are 40, 80, 160, or 320.
func (v *VCNL4040) SetIRDutyCycle(dutyValue uint16) error {

	if dutyValue >= 320 {
		dutyValue = uint16(VCNL4040_PS_DUTY_320)
	} else if dutyValue >= 160 {
		dutyValue = uint16(VCNL4040_PS_DUTY_160)
	} else if dutyValue >= 80 {
		dutyValue = uint16(VCNL4040_PS_DUTY_80)
	} else {
		dutyValue = uint16(VCNL4040_PS_DUTY_40)
	}

	return v.bitMask(VCNL4040_PS_CONF1, LOWER, VCNL4040_PS_DUTY_MASK, byte(dutyValue))
}

// SetLEDCurrent sets the IR LED sink current to one of 8 settings. valid values
// are 50, 75, 100, 120, 140, 160, 180, or 200 (maximum)
func (v *VCNL4040) SetLEDCurrent(current uint8) error {

	if current >= 200 {
		current = VCNL4040_LED_200MA
	} else if current >= 180 {
		current = VCNL4040_LED_180MA
	} else if current >= 160 {
		current = VCNL4040_LED_160MA
	} else if current >= 140 {
		current = VCNL4040_LED_140MA
	} else if current >= 120 {
		current = VCNL4040_LED_120MA
	} else if current >= 100 {
		current = VCNL4040_LED_100MA
	} else if current >= 75 {
		current = VCNL4040_LED_75MA
	} else {
		current = VCNL4040_LED_50MA
	}

	return v.bitMask(VCNL4040_PS_MS, UPPER, VCNL4040_LED_I_MASK, current)
}

// readCommand writes command to sensor and reads the response
func (v *VCNL4040) readCommand(commandCode byte) (uint16, error) {

	readBuf := make([]byte, 2)

	if _, _, err := v.i2c.WriteThenReadBytes([]byte{commandCode}, readBuf); err != nil {
		return 0, err
	}

	// combine the two bytes into a 16-bit value
	return uint16(readBuf[1])<<8 | uint16(readBuf[0]), nil
}

// GetID gets the sensors ID
func (v *VCNL4040) GetID() (uint16, error) {
	return v.readCommand(VCNL4040_ID)
}

// writeCommand writes a 16-bit value to the given command code location
func (v *VCNL4040) writeCommand(commandCode byte, value uint16) error {

	buf := []byte{commandCode, byte(value & 0xFF), byte(value >> 8)}

	if _, err := v.i2c.WriteBytes(buf); err != nil {
		return err
	}

	return nil
}

// bitMask reads a value from a register, masks it, then writes it back
// commandHeight is used to select between the upper or lower byte of command register
// Example: Write dutyValue into PS_CONF1, lower byte, using the Duty_Mask
// bitMask(VCNL4040_PS_CONF1, LOWER, VCNL4040_PS_DUTY_MASK, dutyValue)
func (v *VCNL4040) bitMask(commandAddress byte, commandHeight bool,
	mask byte, thing byte) error {

	var registerContents byte
	var err error

	if commandHeight == LOWER {
		registerContents, err = v.readCommandLower(commandAddress)
	} else {
		registerContents, err = v.readCommandUpper(commandAddress)
	}

	if err != nil {
		return err
	}

	// zero-out the portions of the register we're interested in
	registerContents &= mask

	// mask in new thing
	registerContents |= thing

	// change contents
	if commandHeight == LOWER {
		err = v.writeCommandLower(commandAddress, registerContents)
	} else {
		err = v.writeCommandUpper(commandAddress, registerContents)
	}

	return err
}

// readCommandLower reads the lower byte for the given command code address
func (v *VCNL4040) readCommandLower(commandCode byte) (byte, error) {

	commandValue, err := v.readCommand(commandCode)

	if err != nil {
		return 0, err
	}

	return byte(commandValue & 0xFF), nil
}

// readCommandUpper reads the upper byte for the given command code address
func (v *VCNL4040) readCommandUpper(commandCode byte) (byte, error) {

	commandValue, err := v.readCommand(commandCode)

	if err != nil {
		return 0, err
	}

	return byte(commandValue >> 8), nil
}

// writeCommandLower writes to the lower byte without affecting the upper byte
// for the given command code address
func (v *VCNL4040) writeCommandLower(commandCode byte, newValue byte) error {

	commandValue, err := v.readCommand(commandCode)

	if err != nil {
		return err
	}

	commandValue &= 0xFF00           // Remove lower 8 bits
	commandValue |= uint16(newValue) // Mask in

	return v.writeCommand(commandCode, commandValue)
}

// writeCommandUpper writew to the upper byte without affecting the lower byte
// for the given command code address
func (v *VCNL4040) writeCommandUpper(commandCode byte, newValue byte) error {

	commandValue, err := v.readCommand(commandCode)

	if err != nil {
		return err
	}

	commandValue &= 0x00FF                // Remove upper 8 bits
	commandValue |= uint16(newValue) << 8 // Mask in

	return v.writeCommand(commandCode, commandValue)
}

// GetProximity reads the proximity value.  Values range from 0 to 65535
// where 0 is furthest away and 65535 is closet to sensor.
func (v *VCNL4040) GetProximity() (uint16, error) {
	return v.readCommand(VCNL4040_PS_DATA)
}

// GetAmbient reads the ambient light value. Values range from 0 to 65535
// where 0 is dark and 65535 is a bright light source.
func (v *VCNL4040) GetAmbient() (uint16, error) {
	return v.readCommand(VCNL4040_ALS_DATA)
}

// SetProximityInterruptPersistance sets the proximity interrupt persistance value
// The PS persistence function (PS_PERS, 1, 2, 3, 4) helps to avoid
// false trigger of the PS INT. It defines the amount of consecutive hits
// needed in order for a PS interrupt event to be triggered.
func (v *VCNL4040) SetProximityInterruptPersistance(val ProximityPersistance) error {

	var persValue uint8

	if val == ProximityPersistance1 {
		persValue = VCNL4040_PS_PERS_1
	} else if val == ProximityPersistance2 {
		persValue = VCNL4040_PS_PERS_2
	} else if val == ProximityPersistance3 {
		persValue = VCNL4040_PS_PERS_3
	} else {
		// ProximityPersistance4
		persValue = VCNL4040_PS_PERS_4
	}

	return v.bitMask(VCNL4040_PS_CONF1, LOWER, VCNL4040_PS_PERS_MASK, persValue)
}

// SetAmbientInterruptPersistance sets the Ambient interrupt persistance value
// The ALS persistence function (ALS_PERS, 1, 2, 4, 8) helps to avoid
// false trigger of the ALS INT. It defines the amount of consecutive hits
// needed in order for a ALS interrupt event to be triggered.
// valid values are VCNL4040_ALS_PERS_[1,2,4,8]
func (v *VCNL4040) SetAmbientInterruptPersistance(val AmbientPersistance) error {

	var persValue uint8

	if val == AmbientPersistance1 {
		persValue = VCNL4040_ALS_PERS_1
	} else if val == AmbientPersistance2 {
		persValue = VCNL4040_ALS_PERS_2
	} else if val == AmbientPersistance4 {
		persValue = VCNL4040_ALS_PERS_4
	} else {
		// AmbientPersistance8
		persValue = VCNL4040_ALS_PERS_8
	}

	return v.bitMask(VCNL4040_ALS_CONF, LOWER, VCNL4040_ALS_PERS_MASK, persValue)
}

// EnableAmbientInterrupts turns on ambient light interrupts
func (v *VCNL4040) EnableAmbientInterrupts() error {
	return v.bitMask(VCNL4040_ALS_CONF, LOWER, VCNL4040_ALS_INT_EN_MASK, VCNL4040_ALS_INT_ENABLE)
}

// DisableAmbientInterrupts turns off ambient light interrupts
func (v *VCNL4040) DisableAmbientInterrupts() error {
	return v.bitMask(VCNL4040_ALS_CONF, LOWER, VCNL4040_ALS_INT_EN_MASK, VCNL4040_ALS_INT_DISABLE)
}

// SetProximityInterruptType sets the proximity interrupt type
func (v *VCNL4040) SetProximityInterruptType(val InterruptType) error {

	var interruptValue uint8

	switch val {
	case InterruptDisable:
		interruptValue = VCNL4040_PS_INT_DISABLE
	case InterruptClose:
		interruptValue = VCNL4040_PS_INT_CLOSE
	case InterruptAway:
		interruptValue = VCNL4040_PS_INT_AWAY
	case InterruptBoth:
		interruptValue = VCNL4040_PS_INT_BOTH
	default:
		return fmt.Errorf("unknown interrupt type")
	}

	return v.bitMask(VCNL4040_PS_CONF2, UPPER, VCNL4040_PS_INT_MASK, interruptValue)
}

// EnableActiveForceMode is an extreme power saving way to use PS is to apply
// PS active force mode. Anytime host would like to request one proximity
// measurement, enable the active force mode. This triggers a single PS
// measurement, which can be read from the PS result registers. The sensor stays
// in standby mode constantly.
func (v *VCNL4040) EnableActiveForceMode() error {
	return v.bitMask(VCNL4040_PS_CONF3, LOWER, VCNL4040_PS_AF_MASK, VCNL4040_PS_AF_ENABLE)
}

// DisableActiveForceMode disable active force mode
func (v *VCNL4040) DisableActiveForceMode() error {
	return v.bitMask(VCNL4040_PS_CONF3, LOWER, VCNL4040_PS_AF_MASK, VCNL4040_PS_AF_DISABLE)
}

// TakeSingleProximityMeasurement set trigger bit so sensor takes a force mode
// measurement and returns to standby
func (v *VCNL4040) TakeSingleProximityMeasurement() error {
	return v.bitMask(VCNL4040_PS_CONF3, LOWER, VCNL4040_PS_TRIG_MASK, VCNL4040_PS_TRIG_TRIGGER)
}

// EnableWhiteChannel enable the white measurement channel
func (v *VCNL4040) EnableWhiteChannel() error {
	return v.bitMask(VCNL4040_PS_MS, UPPER, VCNL4040_WHITE_EN_MASK, VCNL4040_WHITE_ENABLE)
}

// DisableWhiteChannel disable the white measurement channel
func (v *VCNL4040) DisableWhiteChannel() error {
	return v.bitMask(VCNL4040_PS_MS, UPPER, VCNL4040_WHITE_EN_MASK, VCNL4040_WHITE_DISABLE)
}

// EnableProximityLogicMode enables the proximity detection logic output mode
// When this mode is selected, the INT pin is pulled low when an object is
// close to the sensor (value is above high threshold) and is reset to high
// when the object moves away (value is below low threshold).
// Register: PS_THDH / PS_THDL define where these threshold levels are set.
func (v *VCNL4040) EnableProximityLogicMode() error {
	return v.bitMask(VCNL4040_PS_MS, UPPER, VCNL4040_PS_MS_MASK, VCNL4040_PS_MS_ENABLE)
}

// DisableProximityLogicMode disable the proximity detection logic output mode
func (v *VCNL4040) DisableProximityLogicMode() error {
	return v.bitMask(VCNL4040_PS_MS, UPPER, VCNL4040_PS_MS_MASK, VCNL4040_PS_MS_DISABLE)
}

// SetProximityCancellation sets the proximity sensing cancelation value which
// helps reduce cross talk with ambient light.  This value will be subtracted
// from the output value read by the sensor before being returned by
// GetProximity()
func (v *VCNL4040) SetProximityCancellation(cancelValue uint16) error {
	return v.writeCommand(VCNL4040_PS_CANC, cancelValue)
}

// SetALSHighThreshold is the value the ambient light sensor (ALS) must go
// above to trigger an interrupt
func (v *VCNL4040) SetALSHighThreshold(threshold uint16) error {
	return v.writeCommand(VCNL4040_ALS_THDH, threshold)
}

// SetALSLowThreshold is the value the ambient light sensor (ALS) must go
// below to trigger an interrupt
func (v *VCNL4040) SetALSLowThreshold(threshold uint16) error {
	return v.writeCommand(VCNL4040_ALS_THDL, threshold)
}

// SetProximityHighThreshold is the value the Proximity Sensor must go
// above to trigger an interrupt
func (v *VCNL4040) SetProximityHighThreshold(threshold uint16) error {
	return v.writeCommand(VCNL4040_PS_THDH, threshold)
}

// SetProximityLowThreshold is the value the Proximity Sensor must go
// below to trigger an interrupt
func (v *VCNL4040) SetProximityLowThreshold(threshold uint16) error {
	return v.writeCommand(VCNL4040_PS_THDL, threshold)
}

// GetWhite reads the White light value
func (v *VCNL4040) GetWhite() (uint16, error) {
	return v.readCommand(VCNL4040_WHITE_DATA)
}

// IsClose returns true if the proximity value rises above the upper threshold
func (v *VCNL4040) IsClose() (bool, error) {

	interruptFlags, err := v.readCommandUpper(VCNL4040_INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & VCNL4040_INT_FLAG_CLOSE) != 0, nil
}

// IsAway returns true if the proximity value drops below the lower threshold
func (v *VCNL4040) IsAway() (bool, error) {

	interruptFlags, err := v.readCommandUpper(VCNL4040_INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & VCNL4040_INT_FLAG_AWAY) != 0, nil
}

// IsLight returns true if the ambient light (ALS) value rises above the upper
// threshold
func (v *VCNL4040) IsLight() (bool, error) {

	interruptFlags, err := v.readCommandUpper(VCNL4040_INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & VCNL4040_INT_FLAG_ALS_HIGH) != 0, nil
}

// IsDark returns true if the ambient light (ALS) value drops below the lower
// threshold
func (v *VCNL4040) IsDark() (bool, error) {

	interruptFlags, err := v.readCommandUpper(VCNL4040_INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & VCNL4040_INT_FLAG_ALS_LOW) != 0, nil
}
