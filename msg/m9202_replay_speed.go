package msg

// ReplaySpeed 快进或快退倍数类型枚举
type ReplaySpeed uint8

const (
	// ReplaySpeedInvalid 表示无效的回放倍数
	ReplaySpeedInvalid ReplaySpeed = iota

	// ReplaySpeed1x 表示1倍回放速度，即正常速度
	ReplaySpeed1x

	// ReplaySpeed2x 表示2倍回放速度
	ReplaySpeed2x

	// ReplaySpeed4x 表示4倍回放速度
	ReplaySpeed4x

	// ReplaySpeed8x 表示8倍回放速度
	ReplaySpeed8x

	// ReplaySpeed16x 表示16倍回放速度
	ReplaySpeed16x
)
