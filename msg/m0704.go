package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
)

type M0704 struct {
	count uint16
	type_ uint8
	items []byte
}

func (m M0704) Count() uint16 {
	return m.count
}

func (m M0704) Type() M0104Type {
	return M0104Type(m.type_)
}

func (m M0704) Items() ([]M0200, error) {
	var res []M0200
	r := binary.NewReader(m.items)

	for i := uint16(0); i < m.Count(); i++ {
		var m0200 M0200

		// 获取数据长度
		length, err := r.ReadUint16()
		if err != nil {
			return nil, err
		}

		// 读取一定长度的定位数据
		m0200Buf, err := r.ReadBytes(int(length))
		if err != nil {
			return nil, err
		}

		// 解析定位数据
		err = m0200.Decode(m0200Buf)
		if err != nil {
			return nil, err
		}
		res = append(res, m0200)
	}
	return res, nil
}

func (m *M0704) Decode(b []byte) (err error) {
	r := binary.NewReader(b)

	if m.count, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("批量位置汇报消息数量解码失败")
	}
	if m.type_, err = r.ReadByte(); err != nil {
		return fmt.Errorf("批量位置汇报消息类型解码失败")
	}
	if m.items, err = r.ReadAllBytes(); err != nil {
		return fmt.Errorf("批量位置汇报消息内容解码失败")
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
