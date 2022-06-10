package v2013

import (
	"errors"

	"github.com/mingkid/g-jtt808/message/msgcomm"
)

type MsgBodyProperty struct {
	Raw uint16
}

// BodyLength 消息体长度
func (p *MsgBodyProperty) BodyLength() uint16 {
	return p.Raw & 0x03ff
}

// SetBodyLength 设置消息体长度
func (p *MsgBodyProperty) SetBodyLength(length uint16) error {
	if length > 0x3ff {
		return BodyLengthTooLarge
	}
	p.Raw |= length
	return nil
}

// EncryptType 加密类型
func (p *MsgBodyProperty) EncryptType() msgcomm.EncryptType {
	return msgcomm.EncryptType((p.Raw >> 10) & 0x01)
}

// SetEncrypt 设置加密
func (p *MsgBodyProperty) SetEncrypt() {
	p.Raw |= 1 << 10
}

// IsSub 分包
func (p *MsgBodyProperty) IsSub() bool {
	return ((p.Raw >> 13) & 0x01) == 1
}

// SetIsSub 设置当前为分包
func (p *MsgBodyProperty) SetIsSub() {
	p.Raw |= 1 << 13
}

// VerFlag 版本标识
func (p *MsgBodyProperty) VerFlag() msgcomm.VersionFlag {
	return msgcomm.VersionFlag((p.Raw >> 14) & 0x01)
}

// SetVerFlag 设置版本标识
func (p *MsgBodyProperty) SetVerFlag(v msgcomm.VersionFlag) {
	p.Raw |= uint16(v) << 14
}

var BodyLengthTooLarge = errors.New("数据过大，请分包")
