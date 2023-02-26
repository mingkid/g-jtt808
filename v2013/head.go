package v2013

import (
	"errors"
	"jtt808/binary"
	"jtt808/msg"
)

type Head struct {
	msgID          uint16
	property       uint16
	phone          string
	serialNum      uint16
	packagePacking *MsgPackagePacking
}

func (h Head) MsgID() msg.MsgID {
	return msg.MsgID(h.msgID)
}

// SetID 设置消息ID
func (h *Head) SetID(id msg.MsgID) {
	h.msgID = uint16(id)
}

func (h Head) Phone() string {
	return h.phone
}

// SetPhone 设置终端手机号
func (h *Head) SetPhone(p string) {
	h.phone = p
}

func (h Head) SerialNum() uint16 {
	return h.serialNum
}

// SetSerialNum 设置消息流水号
func (h *Head) SetSerialNum(num uint16) {
	h.serialNum = num
}

func (h Head) BodyProperty() BodyProperty {
	return BodyProperty(h.property)
}

func (h *Head) Decode(b []byte) (err error) {
	r := binary.NewReader(b)
	if h.msgID, err = r.ReadUint16(); err != nil {
		return
	}
	if h.property, err = r.ReadUint16(); err != nil {
		return
	}
	if h.phone, err = r.ReadBCD(6); err != nil {
		return
	}
	if h.serialNum, err = r.ReadUint16(); err != nil {
		return
	}

	return nil
}

func (h Head) Encode(len uint16) (res []byte, err error) {
	w := binary.NewWriter()

	// 设置消息体长度，BodyProperty为值对象，这里的修改不会对消息头的消息体属性产生影响
	prop, err := h.BodyProperty().SetBodyLength(len)
	if err != nil {
		return
	}

	// 编码
	if err = w.WriteUint16(h.msgID); err != nil {
		return
	}
	if err = w.WriteUint16(uint16(prop)); err != nil {
		return
	}
	if err = w.WriteBCD(h.phone, 6); err != nil {
		return
	}
	if err = w.WriteUint16(h.serialNum); err != nil {
		return
	}
	if h.packagePacking != nil {
		var ppBuf []byte
		ppBuf, err = h.packagePacking.Encode()
		if err = w.Write(ppBuf); err != nil {
			return
		}
	}

	res = w.Bytes()
	return
}

type BodyProperty uint16

func (p BodyProperty) BodyLength() uint16 {
	return uint16(p & 0x03ff)
}

func (p BodyProperty) SetBodyLength(length uint16) (BodyProperty, error) {
	if length > 0x3ff {
		return p, BodyLengthTooLarge
	}
	p |= BodyProperty(length)
	return p, nil
}

func (p BodyProperty) EncryptType() msg.EncryptType {
	return msg.EncryptType((p >> 10) & 0x01)
}

func (p BodyProperty) SetEncrypt() BodyProperty {
	p |= 1 << 10
	return p
}

func (p BodyProperty) IsSub() bool {
	return ((p >> 13) & 0x01) == 1
}

func (p BodyProperty) SetIsSub() BodyProperty {
	p |= 1 << 13
	return p
}

var BodyLengthTooLarge = errors.New("数据过大，请分包")
