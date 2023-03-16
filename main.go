package jtt808

import (
	"encoding/binary"
	msgCom "github.com/mingkid/g-jtt808/msg/common"
)

func ExtractMsgID(b []byte) msgCom.MsgID {
	return msgCom.MsgID(binary.BigEndian.Uint16(b[1:3]))
}
