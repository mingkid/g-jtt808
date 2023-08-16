package msg

// VideoRelatedAlarm 表示视频相关报警信息
type VideoRelatedAlarm []byte

// VideoSignalLossAlarm 检查视频信号丢失报警
func (w VideoRelatedAlarm) VideoSignalLossAlarm() bool {
	return ((w[0x00]) & 0x01) == 1
}

// VideoSignalMaskAlarm 检查视频信号遮挡报警
func (w VideoRelatedAlarm) VideoSignalMaskAlarm() bool {
	return ((w[0x01]) & 0x01) == 1
}

// StorageUnitFailureAlarm 检查存储单元故障报警
func (w VideoRelatedAlarm) StorageUnitFailureAlarm() bool {
	return ((w[0x02]) & 0x01) == 1
}

// OtherVideoDeviceFailure 检查其他视频设备故障报警
func (w VideoRelatedAlarm) OtherVideoDeviceFailure() bool {
	return ((w[0x03]) & 0x01) == 1
}

// OvercrowdingAlarm 检查拥挤报警
func (w VideoRelatedAlarm) OvercrowdingAlarm() bool {
	return ((w[0x04]) & 0x01) == 1
}

// AbnormalDrivingBehavior 检查异常驾驶行为报警
func (w VideoRelatedAlarm) AbnormalDrivingBehavior() bool {
	return ((w[0x05]) & 0x01) == 1
}

// SpecialAlarmStorageReached 检查特殊报警存储达到报警
func (w VideoRelatedAlarm) SpecialAlarmStorageReached() bool {
	return ((w[0x06]) & 0x01) == 1
}
