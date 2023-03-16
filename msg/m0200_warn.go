package msg

type M0200Warn uint32

func (w M0200Warn) Emergency() bool {
	return ((w >> 0) & 0x01) == 1
}

func (w M0200Warn) OverSpeed() bool {
	return ((w >> 1) & 0x01) == 1
}

func (w M0200Warn) Tired() bool {
	return ((w >> 2) & 0x01) == 1
}

func (w M0200Warn) Danger() bool {
	return ((w >> 3) & 0x01) == 1
}

func (w M0200Warn) TermFault() bool {
	return ((w >> 4) & 0x01) == 1
}

func (w M0200Warn) AerialUnConn() bool {
	return ((w >> 5) & 0x01) == 1
}

func (w M0200Warn) AerialShortCircuit() bool {
	return ((w >> 6) & 0x01) == 1
}

func (w M0200Warn) TermUndervoltage() bool {
	return ((w >> 7) & 0x01) == 1
}

func (w M0200Warn) TermPowerFail() bool {
	return ((w >> 8) & 0x01) == 1
}

func (w M0200Warn) LCDFault() bool {
	return ((w >> 9) & 0x01) == 1
}

func (w M0200Warn) TTSFault() bool {
	return ((w >> 10) & 0x01) == 1
}

// CameraFault 摄像头故障
func (w M0200Warn) CameraFault() bool {
	return ((w >> 11) & 0x01) == 1
}

func (w M0200Warn) ICCardFault() bool {
	return ((w >> 12) & 0x01) == 1
}

func (w M0200Warn) PreOverSpeed() bool {
	return ((w >> 13) & 0x01) == 1
}

func (w M0200Warn) PreTired() bool {
	return ((w >> 14) & 0x01) == 1
}

// TotalDriveTimeout 累计驾驶时长超时
func (w M0200Warn) TotalDriveTimeout() bool {
	return ((w >> 18) & 0x01) == 1
}

// StopTimeout 停车超时
func (w M0200Warn) StopTimeout() bool {
	return ((w >> 19) & 0x01) == 1
}

// AreaIO 区域进出
func (w M0200Warn) AreaIO() bool {
	return ((w >> 20) & 0x01) == 1
}

// RoadIO 路线进出
func (w M0200Warn) RoadIO() bool {
	return ((w >> 21) & 0x01) == 1
}

// RoadDriveNotEnoughTime 路线驾驶时长不足
func (w M0200Warn) RoadDriveNotEnoughTime() bool {
	return ((w >> 22) & 0x01) == 1
}

// RoadDeparture 路线偏离
func (w M0200Warn) RoadDeparture() bool {
	return ((w >> 23) & 0x01) == 1
}

// VSSFault 车辆VSS故障
func (w M0200Warn) VSSFault() bool {
	return ((w >> 24) & 0x01) == 1
}

// OilError 油量异常
func (w M0200Warn) OilError() bool {
	return ((w >> 25) & 0x01) == 1
}

// VehStolen 车辆被盗
func (w M0200Warn) VehStolen() bool {
	return ((w >> 26) & 0x01) == 1
}

// IllegalIgnition 非法点火
func (w M0200Warn) IllegalIgnition() bool {
	return ((w >> 27) & 0x01) == 1
}

// IllegalMove 非法位移
func (w M0200Warn) IllegalMove() bool {
	return ((w >> 28) & 0x01) == 1
}

// PreCrash 碰撞预警
func (w M0200Warn) PreCrash() bool {
	return ((w >> 29) & 0x01) == 1
}

// PreRollOver 侧翻预警
func (w M0200Warn) PreRollOver() bool {
	return ((w >> 30) & 0x01) == 1
}

// IllegalOpenDoor 非法开门
func (w M0200Warn) IllegalOpenDoor() bool {
	return ((w >> 31) & 0x01) == 1
}
