package v2013

import "github.com/mingkid/g-jtt808/message/msgcomm"

type MsgHead struct {
	RawID          uint16 `jtt808:""`
	BodyProperty   MsgBodyProperty
	RawPhone       string `jtt808:"6,bcd"`
	RawSerialNum   uint16 `jtt808:""`
	PackagePacking *MsgPackagePacking
}

// ID 消息ID
func (h MsgHead) ID() msgcomm.MsgID {
	return msgcomm.MsgID(h.RawID)
}

// SetID 设置消息ID
func (h *MsgHead) SetID(id msgcomm.MsgID) {
	h.RawID = uint16(id)
}

// SerialNum 消息流水号
func (h MsgHead) SerialNum() uint16 {
	return h.RawSerialNum
}

// SetSerialNum 设置消息流水号
func (h *MsgHead) SetSerialNum(num uint16) {
	h.RawSerialNum = num
}

// Phone 终端手机号
func (h MsgHead) Phone() string {
	return h.RawPhone
}

// SetPhone 设置终端手机号
func (h *MsgHead) SetPhone(p string) {
	h.RawPhone = p
}
