package v2013

import (
	"encoding/binary"
)

type M0200Extra map[byte][]byte

func (e M0200Extra) Mileage() float64 {
	return float64(binary.BigEndian.Uint32(e[0x01])) / 10
}

func (e M0200Extra) Oil() float64 {
	return float64(binary.BigEndian.Uint32(e[0x02])) / 10
}

func (e M0200Extra) Speed() float64 {
	return float64(binary.BigEndian.Uint16(e[0x03])) / 10
}

func (e M0200Extra) ConformWarnID() uint16 {
	return binary.BigEndian.Uint16(e[0x04])
}

func (e M0200Extra) OverSpeedWarn() OverSpeedWarn {
	return e[0x11]
}

func (e M0200Extra) IOPositionWarn() IOPositionWarn {
	return e[0x12]
}

func (e M0200Extra) DriveTooShortWarn() DriveTooShortWarn {
	return e[0x13]
}

func (e M0200Extra) VehSignalStatus() VehSignalStatus {
	return VehSignalStatus(binary.BigEndian.Uint32(e[0x25]))
}

func (e M0200Extra) IOStatus() IOStatus {
	return IOStatus(binary.BigEndian.Uint16(e[0x2A]))
}

func (e M0200Extra) Analog() uint32 {
	return binary.BigEndian.Uint32(e[0x2B])
}

func (e M0200Extra) SignalStrength() byte {
	return e[0x30][0]
}

func (e M0200Extra) GNSSQty() byte {
	return e[0x31][0]
}
