package msg

import "encoding/binary"

// AbnormalDrivingBehavior 表示异常驾驶行为信息
type AbnormalDrivingBehavior []byte

// Fatigue 返回是否存在疲劳驾驶行为
func (a AbnormalDrivingBehavior) Fatigue() bool {
	return ((a[0x00]) & 0x01) == 1
}

// PhoneUsage 返回是否存在手机使用行为
func (a AbnormalDrivingBehavior) PhoneUsage() bool {
	return ((a[0x01]) & 0x01) == 1
}

// Smoking 返回是否存在吸烟行为
func (a AbnormalDrivingBehavior) Smoking() bool {
	return ((a[0x02]) & 0x01) == 1
}

// UserDefined 获取用户自定义的异常驾驶行为标志
func (a AbnormalDrivingBehavior) UserDefined() uint16 {
	return binary.BigEndian.Uint16(a[0x11:0x15])
}
