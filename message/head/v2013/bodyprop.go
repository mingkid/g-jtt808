package v2013

import (
	"errors"

	"github.com/mingkid/g-jtt808/message/head"
	"github.com/mingkid/g-jtt808/message/msgcomm"
)

type MsgBodyProperty uint16

// BodyLength 消息体长度
func (p MsgBodyProperty) BodyLength() uint16 {
	return uint16(p & 0x03ff)
}

// SetBodyLength 设置消息体长度
func (p MsgBodyProperty) SetBodyLength(length uint16) (head.BodyProperty, error) {
	if length > 0x3ff {
		return p, BodyLengthTooLarge
	}
	p |= MsgBodyProperty(length)
	return p, nil
}

// EncryptType 加密类型
func (p MsgBodyProperty) EncryptType() msgcomm.EncryptType {
	return msgcomm.EncryptType((p >> 10) & 0x01)
}

// SetEncrypt 设置加密
func (p MsgBodyProperty) SetEncrypt() head.BodyProperty {
	p |= 1 << 10
	return p
}

// IsSub 分包
func (p MsgBodyProperty) IsSub() bool {
	return ((p >> 13) & 0x01) == 1
}

// SetIsSub 设置当前为分包
func (p MsgBodyProperty) SetIsSub() head.BodyProperty {
	p |= 1 << 13
	return p
}

var BodyLengthTooLarge = errors.New("数据过大，请分包")
