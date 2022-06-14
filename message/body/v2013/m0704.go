package v2013

import (
	jt808b "github.com/mingkid/g-jtt808/binary"
)

type M0704 struct {
	RawCount uint16
	RawType  uint8
	RawItems []byte `jtt808:"0,raw"`
}

func (m M0704) Count() uint16 {
	return m.RawCount
}

func (m M0704) Type() M0104Type {
	return M0104Type(m.RawType)
}

func (m M0704) Items() (res []M0200, err error) {
	er := jt808b.ErrReader{R: jt808b.NewReader(m.RawItems)}
	for i := uint16(0); i < m.Count(); i++ {
		length := er.ReadUint16()
		res = append(res, M0200{
			Warn:         M0200Warn(er.ReadUint32()),
			Status:       M0200Status(er.ReadUint32()),
			RawLatitude:  er.ReadUint32(),
			RawLongitude: er.ReadUint32(),
			RawAltitude:  er.ReadUint16(),
			RawSpeed:     er.ReadUint16(),
			RawDirection: er.ReadUint16(),
			RawTime:      er.ReadBCD(6),
			RawExtras:    er.ReadBytes(int(length - 28)), // 位置附加信息项长度 = 0704位置汇报数据体长度 - 0200位置基本信息长度（28）
		})
	}
	if er.Err != nil {
		return res, er.Err
	}
	return
}

type M0104Type uint8

const (
	// M0704TypeForNormal 正常位置批量汇报
	M0704TypeForNormal M0104Type = iota

	// M0704TypeForBlindSpot 盲区补传
	M0704TypeForBlindSpot
)
