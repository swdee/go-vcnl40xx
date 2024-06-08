package vcnl40xx

import (
	"fmt"
	"github.com/swdee/go-i2c"
)

const (
	LOWER = true
	UPPER = false
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

// Sensor defines the sensor device
type Sensor struct {
	// model defines the sensor model initialized
	model Model
	// cc are the command codes for the sensor model
	cc CommandCodes
	// reg are the register values for the sensor model
	reg Registers
	// i2c bus connection
	i2c *i2c.Options
}

// NewSensor returns a driver instance for the given sensor Model
func NewSensor(m Model) (*Sensor, error) {

	s := &Sensor{
		model: m,
	}

	switch m {
	case VCNL4040:
		s.cc = CommandCodes4040()
		s.reg = Registers4040()

	case VCNL4030:
		s.cc = CommandCodes4030()
		s.reg = Registers4030()

	case VCNL4035:
		s.cc = CommandCodes4035()
		s.reg = Registers4035()

	default:
		return nil, fmt.Errorf("Unknown sensor model")
	}

	return s, nil
}

// Connect to sensor device on the given I2C bus and address
func (s *Sensor) Connect(dev string, addr uint8) error {

	i2c, err := i2c.New(addr, dev)

	if err != nil {
		return fmt.Errorf("i2c bus error: %w", err)
	}

	s.i2c = i2c

	check := s.i2c.GetAddr()

	if check == 0 {
		return fmt.Errorf("I2C device is not initiated")
	}

	id, err := s.GetID()

	if err != nil {
		return fmt.Errorf("error getting sensor ID: %w", err)
	}

	if id != s.model.ID() {
		return fmt.Errorf("unexpected sensor ID value: 0x%04X, wanted 0x%04X", id, s.model.ID())
	}

	return nil
}

// Init initialises the sensor and puts it in default state with proximity
// and ambient light sensors activated
func (s *Sensor) Init() error {

	if err := s.SetLEDCurrent(200); err != nil {
		return fmt.Errorf("error setting LED current: %w", err)
	}

	if err := s.SetIRDutyCycle(40); err != nil {
		return fmt.Errorf("error setting IR duty cycle: %w", err)
	}

	if err := s.SetProximityIntegrationTime(8); err != nil {
		return fmt.Errorf("error setting proximity integration time: %w", err)
	}

	if err := s.SetProximityResolution(16); err != nil {
		return fmt.Errorf("error setting proximity resolution: %w", err)
	}

	if err := s.EnableSmartPersistance(); err != nil {
		return fmt.Errorf("error enabling smart persistance: %w", err)
	}

	if err := s.PowerOnProximity(); err != nil {
		return fmt.Errorf("error powering on proximity function: %w", err)
	}

	if err := s.SetAmbientIntegrationTime(80); err != nil {
		return fmt.Errorf("error setting ambient integration time: %w", err)
	}

	if err := s.PowerOnAmbient(); err != nil {
		return fmt.Errorf("error powering on ambient lighting function: %w", err)
	}

	if s.model == VCNL4030 || s.model == VCNL4035 {
		if err := s.PowerOnWhite(); err != nil {
			return fmt.Errorf("error powering on white channel: %w", err)
		}
	}

	return nil
}

// PowerOnWhite turns on the white channel sensor of the device
func (s *Sensor) PowerOnWhite() error {
	if s.model != VCNL4030 && s.model != VCNL4035 {
		return fmt.Errorf("command not suport for given sensor model")
	}
	return s.bitMask(s.cc.ALS_CONF2, UPPER, s.reg.WHITE_SD_MASK, s.reg.WHITE_SD_POWER_ON)
}

// PowerOffWhite turns off the white channel sensor of the device
func (s *Sensor) PowerOffWhite() error {
	if s.model != VCNL4030 && s.model != VCNL4035 {
		return fmt.Errorf("command not suport for given sensor model")
	}
	return s.bitMask(s.cc.ALS_CONF2, UPPER, s.reg.WHITE_SD_MASK, s.reg.WHITE_SD_POWER_OFF)
}

// PowerOnAmbient turns on the ambient lighting sensor of the device
func (s *Sensor) PowerOnAmbient() error {
	return s.bitMask(s.cc.ALS_CONF, LOWER, s.reg.ALS_SD_MASK, s.reg.ALS_SD_POWER_ON)
}

// PowerOffAmbient turns off the ambient lighting sensor of the device
func (s *Sensor) PowerOffAmbient() error {
	return s.bitMask(s.cc.ALS_CONF, LOWER, s.reg.ALS_SD_MASK, s.reg.ALS_SD_POWER_OFF)
}

// SetAmbientIntegrationTime sets the integration time for the ambient light
// sensor in the number of milliseconds.
// valid values for VCNL4040 are 80, 160, 320, or 640. for VCNL4030 are 50, 100,
// 200, 400, or 800
func (s *Sensor) SetAmbientIntegrationTime(timeValue uint16) error {

	if s.model == VCNL4040 {
		if timeValue >= 640 {
			timeValue = uint16(s.reg.ALS_IT_640MS)
		} else if timeValue >= 320 {
			timeValue = uint16(s.reg.ALS_IT_320MS)
		} else if timeValue >= 160 {
			timeValue = uint16(s.reg.ALS_IT_160MS)
		} else {
			timeValue = uint16(s.reg.ALS_IT_80MS)
		}

	} else if s.model == VCNL4030 {
		if timeValue >= 800 {
			timeValue = uint16(s.reg.ALS_IT_800MS)
		} else if timeValue >= 400 {
			timeValue = uint16(s.reg.ALS_IT_400MS)
		} else if timeValue >= 200 {
			timeValue = uint16(s.reg.ALS_IT_200MS)
		} else if timeValue >= 100 {
			timeValue = uint16(s.reg.ALS_IT_100MS)
		} else {
			timeValue = uint16(s.reg.ALS_IT_50MS)
		}
	}

	return s.bitMask(s.cc.ALS_CONF, LOWER, s.reg.ALS_IT_MASK, byte(timeValue))
}

// PowerOnProximity turns on the proximity sensor of the device
func (s *Sensor) PowerOnProximity() error {
	return s.bitMask(s.cc.PS_CONF1, LOWER, s.reg.PS_SD_MASK, s.reg.PS_SD_POWER_ON)
}

// PowerOffProximity turns off the proximity sensor of the device
func (s *Sensor) PowerOffProximity() error {
	return s.bitMask(s.cc.PS_CONF1, LOWER, s.reg.PS_SD_MASK, s.reg.PS_SD_POWER_OFF)
}

// EnableSmartPersistance to accelerate the PS response time, smart
// persistence prevents the misjudgment of proximity sensing but also keeps
// a fast response time.
func (s *Sensor) EnableSmartPersistance() error {
	return s.bitMask(s.cc.PS_CONF3, LOWER, s.reg.PS_SMART_PERS_MASK, s.reg.PS_SMART_PERS_ENABLE)
}

// DisableSmartPersistence disable smart persistence
func (s *Sensor) DisableSmartPersistence() error {
	return s.bitMask(s.cc.PS_CONF3, LOWER, s.reg.PS_SMART_PERS_MASK, s.reg.PS_SMART_PERS_DISABLE)
}

// SetProximityResolution sets the proximity resolution to either 16 or 12 bit.
// valid values are 12 or 16.
func (s *Sensor) SetProximityResolution(resolutionValue uint8) error {

	if resolutionValue >= 16 {
		resolutionValue = s.reg.PS_HD_16_BIT
	} else {
		resolutionValue = s.reg.PS_HD_12_BIT
	}

	return s.bitMask(s.cc.PS_CONF2, UPPER, s.reg.PS_HD_MASK, resolutionValue)
}

// SetProximityIntegrationTime sets the integration time for the proximity sensor
// which represents the duration of the energy being received. valid values
// are 1, 2, 3, 4, or 8.
func (s *Sensor) SetProximityIntegrationTime(timeValue uint8) error {

	if timeValue >= 8 {
		timeValue = s.reg.PS_IT_8T
	} else if timeValue >= 4 {
		timeValue = s.reg.PS_IT_4T
	} else if timeValue >= 3 {
		timeValue = s.reg.PS_IT_3T
	} else if timeValue >= 2 {
		timeValue = s.reg.PS_IT_2T
	} else {
		timeValue = s.reg.PS_IT_1T
	}

	return s.bitMask(s.cc.PS_CONF1, LOWER, s.reg.PS_IT_MASK, timeValue)
}

// SetIRDutyCycle sets the duty cycle of the IR LED. The higher the duty
// ratio, the faster the response time achieved with higher power
// consumption. For example, PS_Duty = 1/320, peak IRED current = 100 mA,
// averaged current consumption is 100 mA/320 = 0.3125 mA.
// valid values are 40, 80, 160, or 320.
func (s *Sensor) SetIRDutyCycle(dutyValue uint16) error {

	if dutyValue >= 320 {
		dutyValue = uint16(s.reg.PS_DUTY_320)
	} else if dutyValue >= 160 {
		dutyValue = uint16(s.reg.PS_DUTY_160)
	} else if dutyValue >= 80 {
		dutyValue = uint16(s.reg.PS_DUTY_80)
	} else {
		dutyValue = uint16(s.reg.PS_DUTY_40)
	}

	return s.bitMask(s.cc.PS_CONF1, LOWER, s.reg.PS_DUTY_MASK, byte(dutyValue))
}

// SetLEDCurrent sets the IR LED sink current to one of 8 settings. valid values
// are 50, 75, 100, 120, 140, 160, 180, or 200 (maximum)
func (s *Sensor) SetLEDCurrent(current uint8) error {

	if current >= 200 {
		current = s.reg.LED_200MA
	} else if current >= 180 {
		current = s.reg.LED_180MA
	} else if current >= 160 {
		current = s.reg.LED_160MA
	} else if current >= 140 {
		current = s.reg.LED_140MA
	} else if current >= 120 {
		current = s.reg.LED_120MA
	} else if current >= 100 {
		current = s.reg.LED_100MA
	} else if current >= 75 {
		current = s.reg.LED_75MA
	} else {
		current = s.reg.LED_50MA
	}

	return s.bitMask(s.cc.PS_MS, UPPER, s.reg.LED_I_MASK, current)
}

// readCommand writes command to sensor and reads the response
func (s *Sensor) readCommand(commandCode byte) (uint16, error) {

	readBuf := make([]byte, 2)

	if _, _, err := s.i2c.WriteThenReadBytes([]byte{commandCode}, readBuf); err != nil {
		return 0, err
	}

	// combine the two bytes into a 16-bit value
	return uint16(readBuf[1])<<8 | uint16(readBuf[0]), nil
}

// GetID gets the sensors ID
func (s *Sensor) GetID() (uint8, error) {
	return s.readCommandLower(s.cc.ID)
}

// writeCommand writes a 16-bit value to the given command code location
func (s *Sensor) writeCommand(commandCode byte, value uint16) error {

	buf := []byte{commandCode, byte(value & 0xFF), byte(value >> 8)}

	if _, err := s.i2c.WriteBytes(buf); err != nil {
		return err
	}

	return nil
}

// bitMask reads a value from a register, masks it, then writes it back
// commandHeight is used to select between the upper or lower byte of command register
// Example: Write dutyValue into PS_CONF1, lower byte, using the Duty_Mask
// bitMask( s.cc.PS_CONF1, LOWER,  s.reg.PS_DUTY_MASK, dutyValue)
func (s *Sensor) bitMask(commandAddress byte, commandHeight bool,
	mask byte, thing byte) error {

	var registerContents byte
	var err error

	if commandHeight == LOWER {
		registerContents, err = s.readCommandLower(commandAddress)
	} else {
		registerContents, err = s.readCommandUpper(commandAddress)
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
		err = s.writeCommandLower(commandAddress, registerContents)
	} else {
		err = s.writeCommandUpper(commandAddress, registerContents)
	}

	return err
}

// readCommandLower reads the lower byte for the given command code address
func (s *Sensor) readCommandLower(commandCode byte) (byte, error) {

	commandValue, err := s.readCommand(commandCode)

	if err != nil {
		return 0, err
	}

	return byte(commandValue & 0xFF), nil
}

// readCommandUpper reads the upper byte for the given command code address
func (s *Sensor) readCommandUpper(commandCode byte) (byte, error) {

	commandValue, err := s.readCommand(commandCode)

	if err != nil {
		return 0, err
	}

	return byte(commandValue >> 8), nil
}

// writeCommandLower writes to the lower byte without affecting the upper byte
// for the given command code address
func (s *Sensor) writeCommandLower(commandCode byte, newValue byte) error {

	commandValue, err := s.readCommand(commandCode)

	if err != nil {
		return err
	}

	commandValue &= 0xFF00           // Remove lower 8 bits
	commandValue |= uint16(newValue) // Mask in

	return s.writeCommand(commandCode, commandValue)
}

// writeCommandUpper writew to the upper byte without affecting the lower byte
// for the given command code address
func (s *Sensor) writeCommandUpper(commandCode byte, newValue byte) error {

	commandValue, err := s.readCommand(commandCode)

	if err != nil {
		return err
	}

	commandValue &= 0x00FF                // Remove upper 8 bits
	commandValue |= uint16(newValue) << 8 // Mask in

	return s.writeCommand(commandCode, commandValue)
}

// GetProximity reads the proximity value.  Values range from 0 to 65535
// where 0 is furthest away and 65535 is closet to sensor.
func (s *Sensor) GetProximity() (uint16, error) {
	return s.readCommand(s.cc.PS_DATA)
}

// GetAmbient reads the ambient light value. Values range from 0 to 65535
// where 0 is dark and 65535 is a bright light source.
func (s *Sensor) GetAmbient() (uint16, error) {
	return s.readCommand(s.cc.ALS_DATA)
}

// SetProximityInterruptPersistance sets the proximity interrupt persistance value
// The PS persistence function (PS_PERS, 1, 2, 3, 4) helps to avoid
// false trigger of the PS INT. It defines the amount of consecutive hits
// needed in order for a PS interrupt event to be triggered.
func (s *Sensor) SetProximityInterruptPersistance(val ProximityPersistance) error {

	var persValue uint8

	if val == ProximityPersistance1 {
		persValue = s.reg.PS_PERS_1
	} else if val == ProximityPersistance2 {
		persValue = s.reg.PS_PERS_2
	} else if val == ProximityPersistance3 {
		persValue = s.reg.PS_PERS_3
	} else {
		// ProximityPersistance4
		persValue = s.reg.PS_PERS_4
	}

	return s.bitMask(s.cc.PS_CONF1, LOWER, s.reg.PS_PERS_MASK, persValue)
}

// SetAmbientInterruptPersistance sets the Ambient interrupt persistance value
// The ALS persistence function (ALS_PERS, 1, 2, 4, 8) helps to avoid
// false trigger of the ALS INT. It defines the amount of consecutive hits
// needed in order for a ALS interrupt event to be triggered.
// valid values are  s.reg.ALS_PERS_[1,2,4,8]
func (s *Sensor) SetAmbientInterruptPersistance(val AmbientPersistance) error {

	var persValue uint8

	if val == AmbientPersistance1 {
		persValue = s.reg.ALS_PERS_1
	} else if val == AmbientPersistance2 {
		persValue = s.reg.ALS_PERS_2
	} else if val == AmbientPersistance4 {
		persValue = s.reg.ALS_PERS_4
	} else {
		// AmbientPersistance8
		persValue = s.reg.ALS_PERS_8
	}

	return s.bitMask(s.cc.ALS_CONF, LOWER, s.reg.ALS_PERS_MASK, persValue)
}

// EnableAmbientInterrupts turns on ambient light interrupts
func (s *Sensor) EnableAmbientInterrupts() error {
	return s.bitMask(s.cc.ALS_CONF, LOWER, s.reg.ALS_INT_EN_MASK, s.reg.ALS_INT_ENABLE)
}

// DisableAmbientInterrupts turns off ambient light interrupts
func (s *Sensor) DisableAmbientInterrupts() error {
	return s.bitMask(s.cc.ALS_CONF, LOWER, s.reg.ALS_INT_EN_MASK, s.reg.ALS_INT_DISABLE)
}

// SetProximityInterruptType sets the proximity interrupt type
func (s *Sensor) SetProximityInterruptType(val InterruptType) error {

	var interruptValue uint8

	switch val {
	case InterruptDisable:
		interruptValue = s.reg.PS_INT_DISABLE
	case InterruptClose:
		interruptValue = s.reg.PS_INT_CLOSE
	case InterruptAway:
		interruptValue = s.reg.PS_INT_AWAY
	case InterruptBoth:
		interruptValue = s.reg.PS_INT_BOTH
	default:
		return fmt.Errorf("unknown interrupt type")
	}

	return s.bitMask(s.cc.PS_CONF2, UPPER, s.reg.PS_INT_MASK, interruptValue)
}

// EnableActiveForceMode is an extreme power saving way to use PS is to apply
// PS active force mode. Anytime host would like to request one proximity
// measurement, enable the active force mode. This triggers a single PS
// measurement, which can be read from the PS result registers. The sensor stays
// in standby mode constantly.
func (s *Sensor) EnableActiveForceMode() error {
	return s.bitMask(s.cc.PS_CONF3, LOWER, s.reg.PS_AF_MASK, s.reg.PS_AF_ENABLE)
}

// DisableActiveForceMode disable active force mode
func (s *Sensor) DisableActiveForceMode() error {
	return s.bitMask(s.cc.PS_CONF3, LOWER, s.reg.PS_AF_MASK, s.reg.PS_AF_DISABLE)
}

// TakeSingleProximityMeasurement set trigger bit so sensor takes a force mode
// measurement and returns to standby
func (s *Sensor) TakeSingleProximityMeasurement() error {
	return s.bitMask(s.cc.PS_CONF3, LOWER, s.reg.PS_TRIG_MASK, s.reg.PS_TRIG_TRIGGER)
}

// EnableWhiteChannel enable the white measurement channel
func (s *Sensor) EnableWhiteChannel() error {
	return s.bitMask(s.cc.PS_MS, UPPER, s.reg.WHITE_EN_MASK, s.reg.WHITE_ENABLE)
}

// DisableWhiteChannel disable the white measurement channel
func (s *Sensor) DisableWhiteChannel() error {
	return s.bitMask(s.cc.PS_MS, UPPER, s.reg.WHITE_EN_MASK, s.reg.WHITE_DISABLE)
}

// EnableProximityLogicMode enables the proximity detection logic output mode
// When this mode is selected, the INT pin is pulled low when an object is
// close to the sensor (value is above high threshold) and is reset to high
// when the object moves away (value is below low threshold).
// Register: PS_THDH / PS_THDL define where these threshold levels are set.
func (s *Sensor) EnableProximityLogicMode() error {
	return s.bitMask(s.cc.PS_MS, UPPER, s.reg.PS_MS_MASK, s.reg.PS_MS_ENABLE)
}

// DisableProximityLogicMode disable the proximity detection logic output mode
func (s *Sensor) DisableProximityLogicMode() error {
	return s.bitMask(s.cc.PS_MS, UPPER, s.reg.PS_MS_MASK, s.reg.PS_MS_DISABLE)
}

// SetProximityCancellation sets the proximity sensing cancelation value which
// helps reduce cross talk with ambient light.  This value will be subtracted
// from the output value read by the sensor before being returned by
// GetProximity()
func (s *Sensor) SetProximityCancellation(cancelValue uint16) error {
	return s.writeCommand(s.cc.PS_CANC, cancelValue)
}

// SetALSHighThreshold is the value the ambient light sensor (ALS) must go
// above to trigger an interrupt
func (s *Sensor) SetALSHighThreshold(threshold uint16) error {
	return s.writeCommand(s.cc.ALS_THDH, threshold)
}

// SetALSLowThreshold is the value the ambient light sensor (ALS) must go
// below to trigger an interrupt
func (s *Sensor) SetALSLowThreshold(threshold uint16) error {
	return s.writeCommand(s.cc.ALS_THDL, threshold)
}

// SetProximityHighThreshold is the value the Proximity Sensor must go
// above to trigger an interrupt
func (s *Sensor) SetProximityHighThreshold(threshold uint16) error {
	return s.writeCommand(s.cc.PS_THDH, threshold)
}

// SetProximityLowThreshold is the value the Proximity Sensor must go
// below to trigger an interrupt
func (s *Sensor) SetProximityLowThreshold(threshold uint16) error {
	return s.writeCommand(s.cc.PS_THDL, threshold)
}

// GetWhite reads the White light value
func (s *Sensor) GetWhite() (uint16, error) {
	return s.readCommand(s.cc.WHITE_DATA)
}

// IsClose returns true if the proximity value rises above the upper threshold
func (s *Sensor) IsClose() (bool, error) {

	interruptFlags, err := s.readCommandUpper(s.cc.INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & s.reg.INT_FLAG_CLOSE) != 0, nil
}

// IsAway returns true if the proximity value drops below the lower threshold
func (s *Sensor) IsAway() (bool, error) {

	interruptFlags, err := s.readCommandUpper(s.cc.INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & s.reg.INT_FLAG_AWAY) != 0, nil
}

// IsLight returns true if the ambient light (ALS) value rises above the upper
// threshold
func (s *Sensor) IsLight() (bool, error) {

	interruptFlags, err := s.readCommandUpper(s.cc.INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & s.reg.INT_FLAG_ALS_HIGH) != 0, nil
}

// IsDark returns true if the ambient light (ALS) value drops below the lower
// threshold
func (s *Sensor) IsDark() (bool, error) {

	interruptFlags, err := s.readCommandUpper(s.cc.INT_FLAG)

	if err != nil {
		return false, err
	}

	return (interruptFlags & s.reg.INT_FLAG_ALS_LOW) != 0, nil
}
