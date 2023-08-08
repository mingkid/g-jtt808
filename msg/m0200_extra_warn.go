package msg

import "encoding/binary"

// OverSpeedWarn 超速报警附加信息
type OverSpeedWarn []byte

// PositionType 位置类型
func (w OverSpeedWarn) PositionType() PositionType {
	if len(w) == 0 {
		return PositionType(0)
	}
	return PositionType(w[0])
}

// PositionID 区域或路段ID
func (w OverSpeedWarn) PositionID() uint32 {
	if len(w) < 2 {
		return 0
	}
	return binary.BigEndian.Uint32(w[1:])
}

// IOPositionWarn 进出区域/路线报警附加信息
type IOPositionWarn []byte

// PositionType 位置类型
func (w IOPositionWarn) PositionType() PositionType {
	if len(w) == 0 {
		return PositionType(0)
	}
	return PositionType(w[0])
}

// PositionID 区域或线路ID
func (w IOPositionWarn) PositionID() uint32 {
	if len(w) < 5 {
		return 0
	}
	return binary.BigEndian.Uint32(w[1:4])
}

// Direction 方向
func (w IOPositionWarn) Direction() IOType {
	if len(w) < 6 {
		return IOType(0)
	}
	return IOType(w[5])
}

// DriveTooShortWarn 路线行驶时间不足/过长报警附加信息
type DriveTooShortWarn []byte

// RoadID 路段ID
func (w DriveTooShortWarn) RoadID() uint32 {
	if len(w) < 5 {
		return 0
	}
	return binary.BigEndian.Uint32(w[0:4])
}

// DriveTime 路段行驶时间
func (w DriveTooShortWarn) DriveTime() uint16 {
	if len(w) < 6 {
		return 0
	}
	return binary.BigEndian.Uint16(w[4:5])
}

// Result 结果
func (w DriveTooShortWarn) Result() byte {
	if len(w) < 7 {
		return 0
	}
	return w[6]
}

// PositionType 位置类型
type PositionType byte

const (
	PositionTypeForNone PositionType = iota
	// PositionTypeForRound 圆形区域
	PositionTypeForRound
	// PositionTypeForRectangle 矩形区域
	PositionTypeForRectangle
	// PositionTypeForPolygon 多边形区域
	PositionTypeForPolygon
	// PositionTypeForRoad 线路
	PositionTypeForRoad
)

// IOType 进出类型
type IOType byte

const (
	// M0200IOTypeForIn 进
	M0200IOTypeForIn IOType = iota
	// M0200IOTypeForOut 出
	M0200IOTypeForOut
)
