package head

import "github.com/mingkid/g-jtt808/message/msgcomm"

type Head interface {
	ID() msgcomm.MsgID
	SetID(id msgcomm.MsgID)
	SerialNum() uint16
	SetSerialNum(uint16)
	Phone() string
	SetPhone(p string)
}

// ProtocolVerRecorder 协议版本记录者
type ProtocolVerRecorder interface {
	// ProtocolVer 协议版本
	ProtocolVer() byte
	// SetProtocolVer 设置协议版本
	SetProtocolVer(v byte)
}

type BodyLengthRecorder interface {
	BodyLength() uint16
	SetBodyLength(uint16) error
}

type BodyProperty interface {
	BodyLengthRecorder
	EncryptType() msgcomm.EncryptType
	SetEncrypt()
	IsSub() bool
	SetIsSub()
}
