package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
	"time"
)

type M0200 struct {
	warn      uint32
	status    uint32
	latitude  uint32
	longitude uint32
	altitude  uint16
	speed     uint16
	direction uint16
	time      []byte
	extras    []byte
}

func (m M0200) Warn() M0200Warn {
	return M0200Warn(m.warn)
}

func (m M0200) Status() M0200Status {
	return M0200Status(m.status)
}

func (m M0200) Latitude() uint32 {
	return m.latitude
}

func (m M0200) Longitude() uint32 {
	return m.longitude
}

func (m M0200) Altitude() uint16 {
	return m.altitude
}

func (m M0200) Speed() uint16 {
	return m.speed
}

func (m M0200) Direction() uint16 {
	return m.direction
}

func (m M0200) Time() (res time.Time, err error) {
	r := binary.NewReader(m.time)

	timeStr, err := r.ReadBCD(6)
	if err != nil {
		return res, err
	}
	return time.ParseInLocation("20060102150405", "20"+timeStr, time.Local)
}

func (m M0200) Extras() (result M0200Extra, err error) {
	if len(m.extras) == 0 {
		return
	}

	result = make(M0200Extra, 72)

	r := binary.NewReader(m.extras)
	for {
		// 解析附加信息ID
		id, err := r.ReadByte()
		if err != nil {
			break
		}

		// 解析附加信息长度
		length, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		// 解析附加信息数据
		v, err := r.ReadBytes(int(length))
		if err != nil {
			return nil, err
		}
		result[id] = v
	}
	return
}

func (m *M0200) Decode(b []byte) (err error) {
	r := binary.NewReader(b)

	if m.warn, err = r.ReadUint32(); err != nil {
		return fmt.Errorf("报警标志解码失败")
	}
	if m.status, err = r.ReadUint32(); err != nil {
		return fmt.Errorf("状态解码失败")
	}
	if m.latitude, err = r.ReadUint32(); err != nil {
		return fmt.Errorf("纬度解码失败")
	}
	if m.longitude, err = r.ReadUint32(); err != nil {
		return fmt.Errorf("经度解码失败")
	}
	if m.altitude, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("高程解码失败")
	}
	if m.speed, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("速度解码失败")
	}
	if m.direction, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("方向解码失败")
	}
	if m.time, err = r.ReadBytes(6); err != nil {
		return fmt.Errorf("定位时间解码失败")
	}
	if m.extras, err = r.ReadAllBytes(); err != nil {
		return fmt.Errorf("附加信息解码失败")
	}

	return
}
