package msg

import (
	"encoding/binary"
)

// M0200Extra JT/T808消息类型0x0200的附加信息
type M0200Extra map[byte][]byte

// Mileage 0x01键 里程值
func (e M0200Extra) Mileage() float32 {
	return float32(binary.BigEndian.Uint32(e[0x01])) / 10
}

// Oil 0x02键 油量值
func (e M0200Extra) Oil() float32 {
	return float32(binary.BigEndian.Uint16(e[0x02])) / 10
}

// Speed 0x03键 速度值
func (e M0200Extra) Speed() float32 {
	return float32(binary.BigEndian.Uint16(e[0x03])) / 10
}

// ConformWarnID 0x04键 报警事件ID
func (e M0200Extra) ConformWarnID() uint16 {
	return binary.BigEndian.Uint16(e[0x04])
}

// OverSpeedWarn 0x11键 超速报警信息
func (e M0200Extra) OverSpeedWarn() OverSpeedWarn {
	return e[0x11]
}

// IOPositionWarn 0x12键 进出区域/路线报警信息
func (e M0200Extra) IOPositionWarn() IOPositionWarn {
	return e[0x12]
}

// DriveTooShortWarn 0x13键 路段行驶时间不足/过长报警信息
func (e M0200Extra) DriveTooShortWarn() DriveTooShortWarn {
	return e[0x13]
}

// VehSignalStatus 0x25键 扩展车辆信号状态位
func (e M0200Extra) VehSignalStatus() VehSignalStatus {
	return VehSignalStatus(binary.BigEndian.Uint32(e[0x25]))
}

// IOStatus 0x2A键 I/O状态位
func (e M0200Extra) IOStatus() IOStatus {
	return IOStatus(binary.BigEndian.Uint16(e[0x2A]))
}

// Analog 0x2B键 模拟量值
func (e M0200Extra) Analog() uint32 {
	return binary.BigEndian.Uint32(e[0x2B])
}

// SignalStrength 0x30键 无线通信网络信号强度。
func (e M0200Extra) SignalStrength() byte {
	return e[0x30][0]
}

// GNSSQty 0x31键 GNSS定位卫星数
func (e M0200Extra) GNSSQty() byte {
	return e[0x31][0]
}
