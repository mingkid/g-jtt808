package msg

// M9201ReplaySpeed 回放速度枚举
type M9201ReplaySpeed uint8

const (
	// M9201ReplaySpeedInvalid 0无效
	M9201ReplaySpeedInvalid M9201ReplaySpeed = iota

	// M9201ReplaySpeed1x 1倍速
	M9201ReplaySpeed1x

	// M9201ReplaySpeed2x 2倍速
	M9201ReplaySpeed2x

	// M9201ReplaySpeed4x 4倍速
	M9201ReplaySpeed4x

	// M9201ReplaySpeed8x 8倍速
	M9201ReplaySpeed8x

	// M9201ReplaySpeed16x 16倍速
	M9201ReplaySpeed16x
)
