package v2013

import "jtt808/binary"

type MsgPackagePacking struct {
	total     uint16
	serialNum uint16
}

// Total 消息包总数
func (p MsgPackagePacking) Total() uint16 {
	return p.total
}

// SetTotal 设置消息包总数
func (p *MsgPackagePacking) SetTotal(total uint16) *MsgPackagePacking {
	p.total = total
	return p
}

// SerialNum 包序号
func (p MsgPackagePacking) SerialNum() uint16 {
	return p.serialNum
}

// SetSerialNum 设置包序号
func (p *MsgPackagePacking) SetSerialNum(SerialNum uint16) *MsgPackagePacking {
	p.serialNum = SerialNum
	return p
}

func (p MsgPackagePacking) Encode() (b []byte, err error) {
	w := binary.NewWriter()
	if err = w.WriteUint16(p.total); err != nil {
		return
	}
	if err = w.WriteUint16(p.serialNum); err != nil {
		return
	}
	b = w.Bytes()
	return
}
