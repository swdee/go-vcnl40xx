package vcnl40xx

// Model defines the sensor model number
type Model int

const (
	// model numbers
	VCNL4040 Model = 1
	VCNL4030 Model = 2
	VCNL4035 Model = 3

	// I2C address based on model code
	VCNL4040Address  = 0x60
	VCNL4040SensorID = 0x86

	VCNL4030XAddress  = 0x60
	VCNL40301XAddress = 0x51
	VCNL40302XAddress = 0x40
	VCNL40303XAddress = 0x41
	VCNL4030SensorID  = 0x80

	VCNL4035XAddress  = 0x60
	VCNL40351XAddress = 0x51
	VCNL40352XAddress = 0x40
	VCNL40353XAddress = 0x41
	VCNL4035SensorID  = 0x80
)

// ID returns the model ID
func (m Model) ID() uint8 {

	switch m {
	case VCNL4040:
		return VCNL4040SensorID

	case VCNL4030:
		return VCNL4030SensorID

	case VCNL4035:
		return VCNL4035SensorID

	default:
		// return invalid value to cause error
		return 0xAF
	}
}
