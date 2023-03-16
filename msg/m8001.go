package msg

import (
	"github.com/mingkid/g-jtt808/binary"
	"github.com/mingkid/g-jtt808/msg/common"
)

type M8001 struct {
	serialNumber uint16
	msgID        uint16
	result       uint8
}

func (m M8001) SerialNumber() uint16 {
	return m.serialNumber
}

func (m M8001) MsgID() common.MsgID {
	return common.MsgID(m.msgID)
}

func (m M8001) Result() M8001Result {
	return M8001Result(m.result)
}

func (m *M8001) SetSerialNumber(number uint16) *M8001 {
	m.serialNumber = number
	return m
}

func (m *M8001) SetMsgID(id common.MsgID) *M8001 {
	m.msgID = uint16(id)
	return m
}

func (m *M8001) SetResult(res M8001Result) *M8001 {
	m.result = uint8(res)
	return m
}

func (m M8001) Encode() (res []byte, err error) {
	w := binary.NewWriter()

	if err = w.WriteUint16(m.serialNumber); err != nil {
		return
	}
	if err = w.WriteUint16(m.msgID); err != nil {
		return
	}
	if err = w.WriteUint8(m.result); err != nil {
		return
	}

	res = w.Bytes()
	return
}
