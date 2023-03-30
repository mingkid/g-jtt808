package msg

import (
	"bytes"
	"fmt"
	"io"

	"github.com/mingkid/g-jtt808/binary"
	"github.com/mingkid/g-jtt808/msg/common"
)

// MsgHead 消息编码解码器
type MsgHead interface {
	MsgDecoder
	Encode(len uint16) (res []byte, err error)
	ByteLength() int
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
	if err = w.WriteUint8(common.IdentityBit); err != nil {
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
	if err = w.WriteByte(calcChecksum(w.Bytes()[1:])); err != nil {
		return
	}

	// 添加标识尾
	if err = w.WriteUint8(common.IdentityBit); err != nil {
		return
	}

	b = w.Bytes()
	return
}

// NewPlantFormMsg 创建平台消息
func NewPlantFormMsg(head MsgHead, body MsgEncoder) *PlantFormMsg {
	return &PlantFormMsg{
		head: head,
		body: body,
	}
}

type TermMsg struct {
	head MsgHead
	body MsgDecoder
}

// Decode 解码消息。转义的消息不符合规范时会报错。
func (m TermMsg) Decode(b []byte) (err error) {
	var buf []byte

	buf, err = m.Unescape(b)
	if err = m.head.Decode(buf[1:]); err != nil {
		return
	}
	if m.body == nil {
		return
	}
	if err = m.body.Decode(buf[m.head.ByteLength()+1 : len(buf)-2]); err != nil {
		return
	}
	return
}

// Unescape 反转义。根据规则 0x7d 0x01 转换为 0x7d，0x7d 0x02 转换为 0x7e。当0x7d 后一位字节不是 0x01 或 0x02 会报错。
// 0x7d 后一位字节不存在时，会报错。
func (m TermMsg) Unescape(data []byte) (res []byte, err error) {
	var (
		writer bytes.Buffer
		b      byte
		reader = bytes.NewReader(data)
	)

	for {
		// 读取每一个字节，读取到结尾时终止循环
		b, err = reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		if b == common.Escape {
			// 处理需要转义的字节，读取 0x7d 后一位字节
			var nb byte
			nb, err = reader.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("反转义时，0x7d后面不存在字节")
			}
			switch nb {
			case 0x01:
				// 根据规则 0x7d 0x01 转换为 0x7d
				writer.WriteByte(common.Escape)
				continue
			case 0x02:
				// 根据规则 0x7d 0x02 转换为 0x7e
				writer.WriteByte(common.IdentityBit)
				continue
			default:
				return nil, fmt.Errorf("0x7d后字节不符合规则，应为0x01或0x02, 实际为0x%x", nb)
			}
		}

		// 不需要转义的字节直接写入
		writer.WriteByte(b)
	}

	return writer.Bytes(), nil
}

func NewTermMsg(head MsgHead, body MsgDecoder) *TermMsg {
	return &TermMsg{
		head: head,
		body: body,
	}
}

// calcChecksum 计算校验和
func calcChecksum(b []byte) (result byte) {
	result = b[0] // 取第0个，从第1个开始异或
	for i := 1; i < len(b); i++ {
		result = result ^ b[i]
	}
	return
}
