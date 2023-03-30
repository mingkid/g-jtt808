package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
	"github.com/mingkid/g-jtt808/msg/common"
)

// M0001 终端通用应答
type M0001 struct {
	serialNo  uint16 // 应答流水号
	msgID     uint16 // 应答消息ID
	result    byte   // 处理结果
	errorCode uint16 // 错误代码
}

// SerialNo 应答流水号
func (m *M0001) SerialNo() uint16 {
	return m.serialNo
}

// MsgID 应答消息ID
func (m *M0001) MsgID() common.MsgID {
	return common.MsgID(m.msgID)
}

// Result 处理结果
func (m *M0001) Result() M0001Result {
	return M0001Result(m.result)
}

// ErrorCode 错误代码
func (m *M0001) ErrorCode() uint16 {
	return m.errorCode
}

// Decode decodes the binary data in b and stores the result in the
// M0001 struct.
func (m *M0001) Decode(b []byte) (err error) {
	r := binary.NewReader(b)

	// Read the serial number from the buffer
	if m.serialNo, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("解码应答流水号失败")
	}
	// Read the message ID from the buffer
	if m.msgID, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("解码应答消息ID失败")
	}
	// Read the result from the buffer
	if m.result, err = r.ReadByte(); err != nil {
		return fmt.Errorf("解码处理结果失败")
	}
	// Read the error code from the buffer
	if m.errorCode, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("解码错误代码失败")
	}

	return
}
