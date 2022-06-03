package v2013

import "github.com/mingkid/g-jtt808/message/msgcomm"

type M8001 struct {
	RawSerialNumber uint16 `jtt808:""`
	RawMsgID        uint16 `jtt808:""`
	RawResult       uint8  `jtt808:""`
}

func (m M8001) SerialNumber() uint16 {
	return m.RawSerialNumber
}

func (m M8001) MsgID() msgcomm.MsgID {
	return msgcomm.MsgID(m.RawMsgID)
}

func (m M8001) Result() M8001Result {
	return M8001Result(m.RawResult)
}

func (m *M8001) SetSerialNumber(number uint16) *M8001 {
	m.RawSerialNumber = number
	return m
}

func (m *M8001) SetMsgID(id msgcomm.MsgID) *M8001 {
	m.RawMsgID = uint16(id)
	return m
}

func (m *M8001) SetResult(result M8001Result) *M8001 {
	m.RawResult = uint8(result)
	return m
}

type M8001Result uint8

const (
	// M8001Success 成功
	M8001Success M8001Result = iota

	// M8001Fail 失败
	M8001Fail

	// M8001MsgErr 消息有误
	M8001MsgErr

	// M8001Unsupported 不支持
	M8001Unsupported

	// M8001WarnConfirm 报警处理确认
	M8001WarnConfirm
)
