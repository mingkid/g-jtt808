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

type BodyProperty interface {
	BodyLength() uint16
	SetBodyLength(uint16) (BodyProperty, error)
	EncryptType() msgcomm.EncryptType
	SetEncrypt() BodyProperty
	IsSub() bool
	SetIsSub() BodyProperty
}
