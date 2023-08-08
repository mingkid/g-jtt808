package msg

import (
	"encoding/binary"
)

// M0200Extra JT/T808消息类型0x0200的附加信息
type M0200Extra map[byte][]byte

// Mileage 0x01键 里程值
func (e M0200Extra) Mileage() uint32 {
	val, ok := e[0x01]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint32(val)
}

// Oil 0x02键 油量值
func (e M0200Extra) Oil() uint16 {
	val, ok := e[0x02]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// Speed 0x03键 速度值
func (e M0200Extra) Speed() uint16 {
	val, ok := e[0x03]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// ConformWarnID 0x04键 报警事件ID
func (e M0200Extra) ConformWarnID() uint16 {
	val, ok := e[0x04]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// TirePressure 0x05键 胎压值
func (e M0200Extra) TirePressure() uint16 {
	val, ok := e[0x05]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// CabinTemperature 0x06键 车厢温度
func (e M0200Extra) CabinTemperature() uint16 {
	val, ok := e[0x06]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// OverSpeedWarn 0x11键 超速报警信息
func (e M0200Extra) OverSpeedWarn() OverSpeedWarn {
	val, ok := e[0x11]
	if !ok {
		return make([]byte, 0)
	}
	return val
}

// IOPositionWarn 0x12键 进出区域/路线报警信息
func (e M0200Extra) IOPositionWarn() IOPositionWarn {
	val, ok := e[0x12]
	if !ok {
		return make([]byte, 0)
	}
	return val
}

// DriveTooShortWarn 0x13键 路段行驶时间不足/过长报警信息
func (e M0200Extra) DriveTooShortWarn() DriveTooShortWarn {
	val, ok := e[0x13]
	if !ok {
		return make([]byte, 0)
	}
	return val
}

// VehSignalStatus 0x25键 扩展车辆信号状态位
func (e M0200Extra) VehSignalStatus() VehSignalStatus {
	val, ok := e[0x25]
	if !ok {
		return 0
	}
	return VehSignalStatus(binary.BigEndian.Uint32(val))
}

// IOStatus 0x2A键 I/O状态位
func (e M0200Extra) IOStatus() IOStatus {
	val, ok := e[0x2A]
	if !ok {
		return 0
	}
	return IOStatus(binary.BigEndian.Uint16(val))
}

// Analog 0x2B键 模拟量值
func (e M0200Extra) Analog() uint32 {
	val, ok := e[0x2B]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint32(val)
}

// SignalStrength 0x30键 无线通信网络信号强度。
func (e M0200Extra) SignalStrength() byte {
	val, ok := e[0x30]
	if !ok || len(val) == 0 {
		return 0
	}
	return val[0]
}

// GNSSQty 0x31键 GNSS定位卫星数
func (e M0200Extra) GNSSQty() byte {
	val, ok := e[0x31]
	if !ok || len(val) == 0 {
		return 0
	}
	return val[0]
}
