package jtt808

import (
	"jtt808/binary"
	"jtt808/msg"
)

// MsgHead 消息编码解码器
type MsgHead interface {
	MsgDecoder
	Encode(len uint16) (res []byte, err error)
}

// MsgEncoder 消息编码器
type MsgEncoder interface {
	Encode() (b []byte, err error)
}

// MsgDecoder 消息解码器
type MsgDecoder interface {
	Decode(b []byte) error
}

type PlantFormMsg struct {
	head MsgHead
	body MsgEncoder
}

func (m PlantFormMsg) Encode() (b []byte, err error) {
	w := binary.NewWriter()

	// 添加标识头
	if err = w.WriteUint8(msg.IdentityBit); err != nil {
		return
	}

	// 构建消息内容
	var headBuf, bodyBuf []byte
	if bodyBuf, err = m.body.Encode(); err != nil {
		return
	}
	if headBuf, err = m.head.Encode(uint16(len(bodyBuf))); err != nil {
		return
	}
	if err = w.Write(headBuf); err != nil {
		return
	}
	if err = w.Write(bodyBuf); err != nil {
		return
	}

	// 计算检验和
	if err = w.WriteByte(m.calcChecksum(w.Bytes()[1:])); err != nil {
		return
	}

	// 添加标识尾
	if err = w.WriteUint8(msg.IdentityBit); err != nil {
		return
	}

	b = w.Bytes()
	return
}

func (m PlantFormMsg) calcChecksum(b []byte) (result byte) {
	result = b[0] // 取第0个，从第1个开始异或
	for i := 1; i < len(b); i++ {
		result = result ^ b[i]
	}
	return
}

type TermMsg struct {
	head MsgHead
	body MsgDecoder
}

//func (m TermMsg) Decode(b []byte) (err error) {
//	// 转义
//	//
//}
