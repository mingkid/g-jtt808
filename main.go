package jtt808

import (
	"encoding/binary"
	msgCom "github.com/mingkid/g-jtt808/msg/common"
)

// ExtractMsgID 提取消息 ID
func ExtractMsgID(b []byte) msgCom.MsgID {
	return ExtractMsgIDFrom(b, 1)
}

// ExtractMsgIDFrom 从指定位置提取消息 ID
func ExtractMsgIDFrom(b []byte, start int) msgCom.MsgID {
	return msgCom.MsgID(binary.BigEndian.Uint16(b[start : start+2]))
}
