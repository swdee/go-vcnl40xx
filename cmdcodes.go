package vcnl40xx

// CommandCodes defines the register command codes from the sensor datasheet
type CommandCodes struct {
	ALS_CONF   byte
	ALS_CONF2  byte
	ALS_THDH   byte
	ALS_THDL   byte
	PS_CONF1   byte
	PS_CONF2   byte
	PS_CONF3   byte
	PS_MS      byte
	PS_CANC    byte
	PS_THDL    byte
	PS_THDH    byte
	PS_DATA    byte
	ALS_DATA   byte
	WHITE_DATA byte
	INT_FLAG   byte
	ID         byte
	// extra fields for 4035
	PS_DATA1 byte
	PS_DATA2 byte
	PS_DATA3 byte
}

// CommandCodes4040 returns the command code values for the VCNL4040 sensor
func CommandCodes4040() CommandCodes {
	return CommandCodes{
		ALS_CONF:   0x00,
		ALS_THDH:   0x01,
		ALS_THDL:   0x02,
		PS_CONF1:   0x03, // Lower
		PS_CONF2:   0x03, // Upper
		PS_CONF3:   0x04, // Lower
		PS_MS:      0x04, // Upper
		PS_CANC:    0x05,
		PS_THDL:    0x06,
		PS_THDH:    0x07,
		PS_DATA:    0x08,
		ALS_DATA:   0x09,
		WHITE_DATA: 0x0A,
		INT_FLAG:   0x0B, // Upper
		ID:         0x0C,
	}
}

// CommandCodes4030 returns the command code values for the VCNL4030 sensor
func CommandCodes4030() CommandCodes {
	return CommandCodes{
		ALS_CONF:   0x00,
		ALS_CONF2:  0x00,
		ALS_THDH:   0x01,
		ALS_THDL:   0x02,
		PS_CONF1:   0x03, // Lower
		PS_CONF2:   0x03, // Upper
		PS_CONF3:   0x04, // Lower
		PS_MS:      0x04, // Upper
		PS_CANC:    0x05,
		PS_THDL:    0x06,
		PS_THDH:    0x07,
		PS_DATA:    0x08,
		ALS_DATA:   0x0B,
		WHITE_DATA: 0x0C,
		INT_FLAG:   0x0D, // Upper
		ID:         0x0E,
	}
}

// CommandCodes4035 returns the command code values for the VCNL4035 sensor
func CommandCodes4035() CommandCodes {
	return CommandCodes{
		ALS_CONF:  0x00,
		ALS_CONF2: 0x00,
		ALS_THDH:  0x01,
		ALS_THDL:  0x02,
		PS_CONF1:  0x03, // Lower
		PS_CONF2:  0x03, // Upper
		PS_CONF3:  0x04, // Lower
		PS_MS:     0x04, // Upper
		PS_CANC:   0x05,
		PS_THDL:   0x06,
		PS_THDH:   0x07,

		PS_DATA:  0x08, // default this to PS_DATA1
		PS_DATA1: 0x08,
		PS_DATA2: 0x09,
		PS_DATA3: 0x0A,

		ALS_DATA:   0x0B,
		WHITE_DATA: 0x0C,
		INT_FLAG:   0x0D, // Upper
		ID:         0x0E,
	}
}
