package v2013

type M0200Status uint32

// ACCOn ACC开
func (s M0200Status) ACCOn() bool {
	return ((s >> 0) & 0x01) == 1
}

// PositionOn 定位开
func (s M0200Status) PositionOn() bool {
	return ((s >> 1) & 0x01) == 1
}

// IsSLat 南纬
func (s M0200Status) IsSLat() bool {
	return ((s >> 2) & 0x01) == 1
}

// IsWLong 西经
func (s M0200Status) IsWLong() bool {
	return ((s >> 3) & 0x01) == 1
}

// IsOperation 运营
func (s M0200Status) IsOperation() bool {
	return ((s >> 4) & 0x01) != 1
}

// IsEncrypt 加密
func (s M0200Status) IsEncrypt() bool {
	return ((s >> 5) & 0x01) == 1
}

// IsOilChannelNormal 油路正常
func (s M0200Status) IsOilChannelNormal() bool {
	return ((s >> 10) & 0x01) != 1
}

// IsCircuitNormal 电路正常
func (s M0200Status) IsCircuitNormal() bool {
	return ((s >> 11) & 0x01) != 1
}

// IsDoorLock 车门加锁
func (s M0200Status) IsDoorLock() bool {
	return ((s >> 12) & 0x01) == 1
}

// IsFrontDoorOpen 前门打开
func (s M0200Status) IsFrontDoorOpen() bool {
	return ((s >> 13) & 0x01) == 1
}

// IsMidDoorOpen 中门打开
func (s M0200Status) IsMidDoorOpen() bool {
	return ((s >> 14) & 0x01) == 1
}

// IsBackDoorOpen 后门打开
func (s M0200Status) IsBackDoorOpen() bool {
	return ((s >> 15) & 0x01) == 1
}

// IsDriveRoomDoorOpen 驾驶席门打开
func (s M0200Status) IsDriveRoomDoorOpen() bool {
	return ((s >> 16) & 0x01) == 1
}

// IsElseRoomDoorOpen 其它门打开
func (s M0200Status) IsElseRoomDoorOpen() bool {
	return ((s >> 17) & 0x01) == 1
}

// UsedGPS 使用卫星定位
func (s M0200Status) UsedGPS() bool {
	return ((s >> 18) & 0x01) == 1
}

// UsedBPS 使用北斗卫星定位
func (s M0200Status) UsedBPS() bool {
	return ((s >> 19) & 0x01) == 1
}

// UsedGLONASS 使用GLONASS卫星定位
func (s M0200Status) UsedGLONASS() bool {
	return ((s >> 20) & 0x01) == 1
}

// UsedGalileo 使用Galileo卫星定位
func (s M0200Status) UsedGalileo() bool {
	return ((s >> 21) & 0x01) == 1
}
