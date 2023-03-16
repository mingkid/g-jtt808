package msg

// VehSignalStatus 车辆信号状态
type VehSignalStatus uint32

// LowBeam 近光灯
func (s VehSignalStatus) LowBeam() bool {
	return ((s >> 0) & 0x01) == 1
}

// HighBeam 远光灯
func (s VehSignalStatus) HighBeam() bool {
	return ((s >> 1) & 0x01) == 1
}

// RightTurnSignal 右转灯
func (s VehSignalStatus) RightTurnSignal() bool {
	return ((s >> 2) & 0x01) == 1
}

// LeftTurnSignal 左转灯
func (s VehSignalStatus) LeftTurnSignal() bool {
	return ((s >> 3) & 0x01) == 1
}

// Break 刹车
func (s VehSignalStatus) Break() bool {
	return ((s >> 4) & 0x01) == 1
}

// Reverse 倒档
func (s VehSignalStatus) Reverse() bool {
	return ((s >> 5) & 0x01) == 1
}

// FogLights 雾灯
func (s VehSignalStatus) FogLights() bool {
	return ((s >> 6) & 0x01) == 1
}

// PositionLights 示廓灯
func (s VehSignalStatus) PositionLights() bool {
	return ((s >> 7) & 0x01) == 1
}

// Horn 喇叭
func (s VehSignalStatus) Horn() bool {
	return ((s >> 8) & 0x01) == 1
}

// AirConditioning 空调状态
func (s VehSignalStatus) AirConditioning() bool {
	return ((s >> 9) & 0x01) == 1
}

// Neutral 空档
func (s VehSignalStatus) Neutral() bool {
	return ((s >> 10) & 0x01) == 1
}

// Retarder 缓速器
func (s VehSignalStatus) Retarder() bool {
	return ((s >> 11) & 0x01) == 1
}

func (s VehSignalStatus) ABS() bool {
	return ((s >> 12) & 0x01) == 1
}

// Heater 加热器
func (s VehSignalStatus) Heater() bool {
	return ((s >> 13) & 0x01) == 1
}

// Clutch 离合器
func (s VehSignalStatus) Clutch() bool {
	return ((s >> 14) & 0x01) == 1
}

// IOStatus IO状态
type IOStatus uint16

// DeepSleep 深度休眠
func (s IOStatus) DeepSleep() bool {
	return ((s >> 0) & 0x01) == 1
}

// Sleep 休眠
func (s IOStatus) Sleep() bool {
	return ((s >> 1) & 0x01) == 1
}
