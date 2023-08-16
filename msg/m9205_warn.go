package msg

import "encoding/binary"

// M9205Warn 表示JTT1078协议中的9205报警信息
type M9205Warn map[byte][]byte

// VideoRelatedAlarm 获取视频相关报警信息
func (w M9205Warn) VideoRelatedAlarm() VideoRelatedAlarm {
	val, ok := w[0x14]
	if !ok {
		return make([]byte, 0)
	}
	return val
}

// VideoSignalLossStatus 获取视频信号丢失状态
func (w M9205Warn) VideoSignalLossStatus() uint32 {
	val, ok := w[0x15]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint32(val)
}

// VideoSignalMaskStatus 获取视频信号遮挡状态
func (w M9205Warn) VideoSignalMaskStatus() uint32 {
	val, ok := w[0x16]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint32(val)
}

// MemoryFailureAlarm 获取存储器故障报警信息
func (w M9205Warn) MemoryFailureAlarm() MemoryFailureAlarm {
	val, ok := w[0x16]
	if !ok {
		return make([]byte, 0)
	}
	return val
}

// AbnormalDrivingBehavior 获取异常驾驶行为报警信息
func (w M9205Warn) AbnormalDrivingBehavior() AbnormalDrivingBehavior {
	val, ok := w[0x16]
	if !ok {
		return make([]byte, 0)
	}
	return val
}
