package vcnl40xx

// Registers defines the registers from the sensor datasheet
type Registers struct {
	ALS_IT_MASK uint8
	// 4040
	ALS_IT_80MS  uint8
	ALS_IT_160MS uint8
	ALS_IT_320MS uint8
	ALS_IT_640MS uint8
	// 4030
	ALS_IT_50MS  uint8
	ALS_IT_100MS uint8
	ALS_IT_200MS uint8
	ALS_IT_400MS uint8
	ALS_IT_800MS uint8

	// 4030
	ALS_HD_MASK uint8
	ALS_HD_1    uint8
	ALS_HD_2    uint8

	ALS_PERS_MASK uint8
	ALS_PERS_1    uint8
	ALS_PERS_2    uint8
	ALS_PERS_4    uint8
	ALS_PERS_8    uint8

	ALS_INT_EN_MASK uint8
	ALS_INT_DISABLE uint8
	ALS_INT_ENABLE  uint8

	ALS_SD_MASK      uint8
	ALS_SD_POWER_ON  uint8
	ALS_SD_POWER_OFF uint8

	// 4030
	ALS_NS_MASK uint8
	ALS_NS_1    uint8
	ALS_NS_2    uint8

	// 4030
	WHITE_SD_MASK      uint8
	WHITE_SD_POWER_ON  uint8
	WHITE_SD_POWER_OFF uint8

	PS_DUTY_MASK uint8
	PS_DUTY_40   uint8
	PS_DUTY_80   uint8
	PS_DUTY_160  uint8
	PS_DUTY_320  uint8

	PS_PERS_MASK uint8
	PS_PERS_1    uint8
	PS_PERS_2    uint8
	PS_PERS_3    uint8
	PS_PERS_4    uint8

	PS_IT_MASK uint8
	PS_IT_1T   uint8
	PS_IT_15T  uint8
	PS_IT_2T   uint8
	PS_IT_25T  uint8
	PS_IT_3T   uint8
	PS_IT_35T  uint8
	PS_IT_4T   uint8
	PS_IT_8T   uint8

	PS_SD_MASK      uint8
	PS_SD_POWER_ON  uint8
	PS_SD_POWER_OFF uint8

	// 4030
	PS_GAIN_MASK     uint8
	PS_GAIN_TWO_STEP uint8
	PS_GAIN_SINGLE_8 uint8
	PS_GAIN_SINGLE_1 uint8

	PS_HD_MASK   uint8
	PS_HD_12_BIT uint8
	PS_HD_16_BIT uint8

	// 4030
	PS_NS_MASK       uint8
	PS_NS_TWO_STEP_4 uint8
	PS_NS_TWO_STEP_1 uint8

	PS_INT_MASK    uint8
	PS_INT_DISABLE uint8
	PS_INT_CLOSE   uint8
	PS_INT_AWAY    uint8
	PS_INT_BOTH    uint8

	PS_MPS_MASK uint8
	PS_MPS_1    uint8
	PS_MPS_2    uint8
	PS_MPS_4    uint8
	PS_MPS_8    uint8

	LED_I_LOW_MASK    uint8
	LED_I_LOW_DISABLE uint8
	LED_I_LOW_ENABLE  uint8

	PS_SMART_PERS_MASK    uint8
	PS_SMART_PERS_DISABLE uint8
	PS_SMART_PERS_ENABLE  uint8

	PS_AF_MASK    uint8
	PS_AF_DISABLE uint8
	PS_AF_ENABLE  uint8

	PS_TRIG_MASK    uint8
	PS_TRIG_TRIGGER uint8

	// 4030
	CONF3_PS_MS_MASK        uint8
	CONF3_PS_MS_NORMAL      uint8
	CONF3_PS_MS_OUTPUT_MODE uint8

	PS_SC_EN_MASK    uint8
	PS_SC_EN_ENABLE  uint8
	PS_SC_EN_DISABLE uint8

	WHITE_EN_MASK uint8
	WHITE_ENABLE  uint8
	WHITE_DISABLE uint8

	PS_MS_MASK    uint8
	PS_MS_DISABLE uint8
	PS_MS_ENABLE  uint8

	LED_I_MASK uint8
	LED_50MA   uint8
	LED_75MA   uint8
	LED_100MA  uint8
	LED_120MA  uint8
	LED_140MA  uint8
	LED_160MA  uint8
	LED_180MA  uint8
	LED_200MA  uint8

	// 4030
	PS_SC_CUR_MASK uint8
	PS_SC_CUR_1    uint8
	PS_SC_CUR_2    uint8
	PS_SC_CUR_4    uint8
	PS_SC_CUR_8    uint8

	// 4030
	PS_SP_MASK uint8
	PS_SP_1    uint8
	PS_SP_15   uint8

	// 4030
	PS_SPO_MASK   uint8
	PS_SPO_MODE_0 uint8
	PS_SPO_MODE_1 uint8

	INT_FLAG_ALS_LOW  uint8
	INT_FLAG_ALS_HIGH uint8
	INT_FLAG_CLOSE    uint8
	INT_FLAG_AWAY     uint8
}

// Registers4040 returns the register values for the VCNL4040 sensor
func Registers4040() Registers {
	return Registers{
		ALS_IT_MASK:  ^uint8((1 << 7) | (1 << 6)),
		ALS_IT_80MS:  0,
		ALS_IT_160MS: 1 << 7,
		ALS_IT_320MS: 1 << 6,
		ALS_IT_640MS: (1 << 7) | (1 << 6),

		ALS_PERS_MASK: ^uint8((1 << 3) | (1 << 2)),
		ALS_PERS_1:    0,
		ALS_PERS_2:    1 << 2,
		ALS_PERS_4:    1 << 3,
		ALS_PERS_8:    (1 << 3) | (1 << 2),

		ALS_INT_EN_MASK: ^uint8(1 << 1),
		ALS_INT_DISABLE: 0,
		ALS_INT_ENABLE:  1 << 1,

		ALS_SD_MASK:      ^uint8(1 << 0),
		ALS_SD_POWER_ON:  0,
		ALS_SD_POWER_OFF: 1 << 0,

		PS_DUTY_MASK: ^uint8((1 << 7) | (1 << 6)),
		PS_DUTY_40:   0,
		PS_DUTY_80:   (1 << 6),
		PS_DUTY_160:  (1 << 7),
		PS_DUTY_320:  (1 << 7) | (1 << 6),

		PS_PERS_MASK: ^uint8((1 << 5) | (1 << 4)),
		PS_PERS_1:    0,
		PS_PERS_2:    1 << 4,
		PS_PERS_3:    1 << 5,
		PS_PERS_4:    (1 << 5) | (1 << 4),

		PS_IT_MASK: ^uint8((1 << 3) | (1 << 2) | (1 << 1)),
		PS_IT_1T:   0,
		PS_IT_15T:  (1 << 1),
		PS_IT_2T:   (1 << 2),
		PS_IT_25T:  (1 << 2) | (1 << 1),
		PS_IT_3T:   (1 << 3),
		PS_IT_35T:  (1 << 3) | (1 << 1),
		PS_IT_4T:   (1 << 3) | (1 << 2),
		PS_IT_8T:   (1 << 3) | (1 << 2) | (1 << 1),

		PS_SD_MASK:      ^uint8(1 << 0),
		PS_SD_POWER_ON:  0,
		PS_SD_POWER_OFF: 1 << 0,

		PS_HD_MASK:   ^uint8(1 << 3),
		PS_HD_12_BIT: 0,
		PS_HD_16_BIT: 1 << 3,

		PS_INT_MASK:    ^uint8((1 << 1) | (1 << 0)),
		PS_INT_DISABLE: 0,
		PS_INT_CLOSE:   1 << 0,
		PS_INT_AWAY:    1 << 1,
		PS_INT_BOTH:    (1 << 1) | (1 << 0),

		PS_MPS_MASK: ^uint8((1 << 6) | (1 << 5)),
		PS_MPS_1:    0,
		PS_MPS_2:    1 << 5,
		PS_MPS_4:    1 << 6,
		PS_MPS_8:    (1 << 6) | (1 << 5),

		PS_SMART_PERS_MASK:    ^uint8(1 << 4),
		PS_SMART_PERS_DISABLE: 0,
		PS_SMART_PERS_ENABLE:  1 << 4,

		PS_AF_MASK:    ^uint8(1 << 3),
		PS_AF_DISABLE: 0,
		PS_AF_ENABLE:  1 << 3,

		PS_TRIG_MASK:    ^uint8(1 << 2),
		PS_TRIG_TRIGGER: 1 << 2,

		PS_SC_EN_MASK:    ^uint8(1 << 0),
		PS_SC_EN_ENABLE:  0,
		PS_SC_EN_DISABLE: 1 << 0,

		WHITE_EN_MASK: ^uint8(1 << 7),
		WHITE_ENABLE:  0,
		WHITE_DISABLE: 1 << 7,

		PS_MS_MASK:    ^uint8(1 << 6),
		PS_MS_DISABLE: 0,
		PS_MS_ENABLE:  (1 << 6),

		LED_I_MASK: ^uint8((1 << 2) | (1 << 1) | (1 << 0)),
		LED_50MA:   0,
		LED_75MA:   (1 << 0),
		LED_100MA:  (1 << 1),
		LED_120MA:  (1 << 1) | (1 << 0),
		LED_140MA:  (1 << 2),
		LED_160MA:  (1 << 2) | (1 << 0),
		LED_180MA:  (1 << 2) | (1 << 1),
		LED_200MA:  (1 << 2) | (1 << 1) | (1 << 0),

		INT_FLAG_ALS_LOW:  1 << 5,
		INT_FLAG_ALS_HIGH: 1 << 4,
		INT_FLAG_CLOSE:    1 << 1,
		INT_FLAG_AWAY:     1 << 0,
	}
}

// Registers4030 returns the register values for the VCNL4030 sensor
func Registers4030() Registers {
	return Registers{
		ALS_IT_MASK:  ^uint8((1 << 7) | (1 << 6) | (1 << 5)),
		ALS_IT_50MS:  0,
		ALS_IT_100MS: 1 << 5,
		ALS_IT_200MS: 1 << 6,
		ALS_IT_400MS: (1 << 6) | (1 << 5),
		ALS_IT_800MS: 1 << 7,

		ALS_HD_MASK: ^uint8(1 << 4),
		ALS_HD_1:    0,
		ALS_HD_2:    1 << 4,

		ALS_PERS_MASK: ^uint8((1 << 3) | (1 << 2)),
		ALS_PERS_1:    0,
		ALS_PERS_2:    1 << 2,
		ALS_PERS_4:    1 << 3,
		ALS_PERS_8:    (1 << 3) | (1 << 2),

		ALS_INT_EN_MASK: ^uint8(1 << 1),
		ALS_INT_DISABLE: 0,
		ALS_INT_ENABLE:  1 << 1,

		ALS_SD_MASK:      ^uint8(1 << 0),
		ALS_SD_POWER_ON:  0,
		ALS_SD_POWER_OFF: 1 << 0,

		ALS_NS_MASK: ^uint8(1 << 1),
		ALS_NS_1:    0,
		ALS_NS_2:    1 << 1,

		WHITE_SD_MASK:      ^uint8(1 << 0),
		WHITE_SD_POWER_ON:  0,
		WHITE_SD_POWER_OFF: 1 << 0,

		PS_DUTY_MASK: ^uint8((1 << 7) | (1 << 6)),
		PS_DUTY_40:   0,
		PS_DUTY_80:   (1 << 6),
		PS_DUTY_160:  (1 << 7),
		PS_DUTY_320:  (1 << 7) | (1 << 6),

		PS_PERS_MASK: ^uint8((1 << 5) | (1 << 4)),
		PS_PERS_1:    0,
		PS_PERS_2:    1 << 4,
		PS_PERS_3:    1 << 5,
		PS_PERS_4:    (1 << 5) | (1 << 4),

		PS_IT_MASK: ^uint8((1 << 3) | (1 << 2) | (1 << 1)),
		PS_IT_1T:   0,
		PS_IT_15T:  (1 << 1),
		PS_IT_2T:   (1 << 2),
		PS_IT_25T:  (1 << 2) | (1 << 1),
		PS_IT_3T:   (1 << 3),
		PS_IT_35T:  (1 << 3) | (1 << 1),
		PS_IT_4T:   (1 << 3) | (1 << 2),
		PS_IT_8T:   (1 << 3) | (1 << 2) | (1 << 1),

		PS_SD_MASK:      ^uint8(1 << 0),
		PS_SD_POWER_ON:  0,
		PS_SD_POWER_OFF: 1 << 0,

		PS_GAIN_MASK:     ^uint8((1 << 5) | (1 << 4)),
		PS_GAIN_TWO_STEP: 0,
		PS_GAIN_SINGLE_8: 1 << 4,
		PS_GAIN_SINGLE_1: (1 << 5) | (1 << 4),

		PS_HD_MASK:   ^uint8(1 << 3),
		PS_HD_12_BIT: 0,
		PS_HD_16_BIT: 1 << 3,

		PS_NS_MASK:       ^uint8(1 << 2),
		PS_NS_TWO_STEP_4: 0,
		PS_NS_TWO_STEP_1: 1 << 2,

		PS_INT_MASK:    ^uint8((1 << 1) | (1 << 0)),
		PS_INT_DISABLE: 0,
		PS_INT_CLOSE:   1 << 0,
		PS_INT_AWAY:    1 << 1,
		PS_INT_BOTH:    (1 << 1) | (1 << 0),

		LED_I_LOW_MASK:    ^uint8(1 << 7),
		LED_I_LOW_DISABLE: 0,
		LED_I_LOW_ENABLE:  1 << 7,

		PS_SMART_PERS_MASK:    ^uint8(1 << 4),
		PS_SMART_PERS_DISABLE: 0,
		PS_SMART_PERS_ENABLE:  1 << 4,

		PS_AF_MASK:    ^uint8(1 << 3),
		PS_AF_DISABLE: 0,
		PS_AF_ENABLE:  1 << 3,

		PS_TRIG_MASK:    ^uint8(1 << 2),
		PS_TRIG_TRIGGER: 1 << 2,

		CONF3_PS_MS_MASK:        ^uint8(1 << 1),
		CONF3_PS_MS_NORMAL:      0,
		CONF3_PS_MS_OUTPUT_MODE: 1 << 1,

		PS_SC_EN_MASK:    ^uint8(1 << 0),
		PS_SC_EN_ENABLE:  0,
		PS_SC_EN_DISABLE: 1 << 0,

		WHITE_EN_MASK: ^uint8(1 << 7),
		WHITE_ENABLE:  0,
		WHITE_DISABLE: 1 << 7,

		PS_MS_MASK:    ^uint8(1 << 6),
		PS_MS_DISABLE: 0,
		PS_MS_ENABLE:  (1 << 6),

		LED_I_MASK: ^uint8((1 << 2) | (1 << 1) | (1 << 0)),
		LED_50MA:   0,
		LED_75MA:   (1 << 0),
		LED_100MA:  (1 << 1),
		LED_120MA:  (1 << 1) | (1 << 0),
		LED_140MA:  (1 << 2),
		LED_160MA:  (1 << 2) | (1 << 0),
		LED_180MA:  (1 << 2) | (1 << 1),
		LED_200MA:  (1 << 2) | (1 << 1) | (1 << 0),

		PS_SC_CUR_MASK: ^uint8((1 << 6) | (1 << 5)),
		PS_SC_CUR_1:    0,
		PS_SC_CUR_2:    (1 << 5),
		PS_SC_CUR_4:    (1 << 6),
		PS_SC_CUR_8:    (1 << 6) | (1 << 5),

		PS_SP_MASK: ^uint8(1 << 4),
		PS_SP_1:    0,
		PS_SP_15:   1 << 4,

		PS_SPO_MASK:   ^uint8(1 << 3),
		PS_SPO_MODE_0: 0,
		PS_SPO_MODE_1: 1 << 3,

		INT_FLAG_ALS_LOW:  1 << 5,
		INT_FLAG_ALS_HIGH: 1 << 4,
		INT_FLAG_CLOSE:    1 << 1,
		INT_FLAG_AWAY:     1 << 0,
	}
}

// Registers4035 returns the register values for the VCNL4035 sensor
func Registers4035() Registers {
	return Registers{
		ALS_IT_MASK:  ^uint8((1 << 7) | (1 << 6) | (1 << 5)),
		ALS_IT_50MS:  0,
		ALS_IT_100MS: 1 << 5,
		ALS_IT_200MS: 1 << 6,
		ALS_IT_400MS: (1 << 6) | (1 << 5),
		ALS_IT_800MS: 1 << 7,

		ALS_HD_MASK: ^uint8(1 << 4),
		ALS_HD_1:    0,
		ALS_HD_2:    1 << 4,

		ALS_PERS_MASK: ^uint8((1 << 3) | (1 << 2)),
		ALS_PERS_1:    0,
		ALS_PERS_2:    1 << 2,
		ALS_PERS_4:    1 << 3,
		ALS_PERS_8:    (1 << 3) | (1 << 2),

		ALS_INT_EN_MASK: ^uint8(1 << 1),
		ALS_INT_DISABLE: 0,
		ALS_INT_ENABLE:  1 << 1,

		ALS_SD_MASK:      ^uint8(1 << 0),
		ALS_SD_POWER_ON:  0,
		ALS_SD_POWER_OFF: 1 << 0,

		ALS_NS_MASK: ^uint8(1 << 1),
		ALS_NS_1:    0,
		ALS_NS_2:    1 << 1,

		WHITE_SD_MASK:      ^uint8(1 << 0),
		WHITE_SD_POWER_ON:  0,
		WHITE_SD_POWER_OFF: 1 << 0,

		PS_DUTY_MASK: ^uint8((1 << 7) | (1 << 6)),
		PS_DUTY_40:   0,
		PS_DUTY_80:   (1 << 6),
		PS_DUTY_160:  (1 << 7),
		PS_DUTY_320:  (1 << 7) | (1 << 6),

		PS_PERS_MASK: ^uint8((1 << 5) | (1 << 4)),
		PS_PERS_1:    0,
		PS_PERS_2:    1 << 4,
		PS_PERS_3:    1 << 5,
		PS_PERS_4:    (1 << 5) | (1 << 4),

		PS_IT_MASK: ^uint8((1 << 3) | (1 << 2) | (1 << 1)),
		PS_IT_1T:   0,
		PS_IT_15T:  (1 << 1),
		PS_IT_2T:   (1 << 2),
		PS_IT_25T:  (1 << 2) | (1 << 1),
		PS_IT_3T:   (1 << 3),
		PS_IT_35T:  (1 << 3) | (1 << 1),
		PS_IT_4T:   (1 << 3) | (1 << 2),
		PS_IT_8T:   (1 << 3) | (1 << 2) | (1 << 1),

		PS_SD_MASK:      ^uint8(1 << 0),
		PS_SD_POWER_ON:  0,
		PS_SD_POWER_OFF: 1 << 0,

		PS_GAIN_MASK:     ^uint8((1 << 5) | (1 << 4)),
		PS_GAIN_TWO_STEP: 0,
		PS_GAIN_SINGLE_8: 1 << 4,
		PS_GAIN_SINGLE_1: (1 << 5) | (1 << 4),

		PS_HD_MASK:   ^uint8(1 << 3),
		PS_HD_12_BIT: 0,
		PS_HD_16_BIT: 1 << 3,

		PS_NS_MASK:       ^uint8(1 << 2),
		PS_NS_TWO_STEP_4: 0,
		PS_NS_TWO_STEP_1: 1 << 2,

		PS_INT_MASK:    ^uint8((1 << 1) | (1 << 0)),
		PS_INT_DISABLE: 0,
		PS_INT_CLOSE:   1 << 0,
		PS_INT_AWAY:    1 << 1,
		PS_INT_BOTH:    (1 << 1) | (1 << 0),

		LED_I_LOW_MASK:    ^uint8(1 << 7),
		LED_I_LOW_DISABLE: 0,
		LED_I_LOW_ENABLE:  1 << 7,

		PS_SMART_PERS_MASK:    ^uint8(1 << 4),
		PS_SMART_PERS_DISABLE: 0,
		PS_SMART_PERS_ENABLE:  1 << 4,

		PS_AF_MASK:    ^uint8(1 << 3),
		PS_AF_DISABLE: 0,
		PS_AF_ENABLE:  1 << 3,

		PS_TRIG_MASK:    ^uint8(1 << 2),
		PS_TRIG_TRIGGER: 1 << 2,

		CONF3_PS_MS_MASK:        ^uint8(1 << 1),
		CONF3_PS_MS_NORMAL:      0,
		CONF3_PS_MS_OUTPUT_MODE: 1 << 1,

		PS_SC_EN_MASK:    ^uint8(1 << 0),
		PS_SC_EN_ENABLE:  0,
		PS_SC_EN_DISABLE: 1 << 0,

		WHITE_EN_MASK: ^uint8(1 << 7),
		WHITE_ENABLE:  0,
		WHITE_DISABLE: 1 << 7,

		PS_MS_MASK:    ^uint8(1 << 6),
		PS_MS_DISABLE: 0,
		PS_MS_ENABLE:  (1 << 6),

		LED_I_MASK: ^uint8((1 << 2) | (1 << 1) | (1 << 0)),
		LED_50MA:   0,
		LED_75MA:   (1 << 0),
		LED_100MA:  (1 << 1),
		LED_120MA:  (1 << 1) | (1 << 0),
		LED_140MA:  (1 << 2),
		LED_160MA:  (1 << 2) | (1 << 0),
		LED_180MA:  (1 << 2) | (1 << 1),
		LED_200MA:  (1 << 2) | (1 << 1) | (1 << 0),

		PS_SC_CUR_MASK: ^uint8((1 << 6) | (1 << 5)),
		PS_SC_CUR_1:    0,
		PS_SC_CUR_2:    (1 << 5),
		PS_SC_CUR_4:    (1 << 6),
		PS_SC_CUR_8:    (1 << 6) | (1 << 5),

		PS_SP_MASK: ^uint8(1 << 4),
		PS_SP_1:    0,
		PS_SP_15:   1 << 4,

		PS_SPO_MASK:   ^uint8(1 << 3),
		PS_SPO_MODE_0: 0,
		PS_SPO_MODE_1: 1 << 3,

		INT_FLAG_ALS_LOW:  1 << 5,
		INT_FLAG_ALS_HIGH: 1 << 4,
		INT_FLAG_CLOSE:    1 << 1,
		INT_FLAG_AWAY:     1 << 0,
	}
}
