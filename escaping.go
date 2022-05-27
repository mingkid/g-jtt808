package jtt808

import (
	"bytes"
	"fmt"
	"io"

	"github.com/mingkid/g-jtt808/binary"
)

// Unescape 反转义
func Unescape(data []byte) ([]byte, error) {
	var writer bytes.Buffer
	// 把数据交给bytes包
	var reader = bytes.NewReader(data)

	for {
		// 读取每一个字节
		currentByte, err := reader.ReadByte()

		if err != nil {
			// 没有更多输入的时候就停止循环
			if err == io.EOF {
				break
			}

			// 读取异常，返回异常
			return nil, err
		}

		if currentByte != binary.Escape {
			// 把不需要转义的字符放到缓冲区
			err := writer.WriteByte(currentByte)

			if err != nil {
				return nil, err
			}

			continue
		}

		// 读取 0x7d 后一位字符
		nextByte, err := reader.ReadByte()

		if err != nil {
			return nil, err
		}

		// 处理需要转义的字符
		switch nextByte {
		case 0x01:
			// 根据规则 0x7d 0x01 转换为 0x7d 到缓冲区
			writer.WriteByte(binary.Escape)
		case 0x02:
			// 根据规则 0x7d 0x01 转换为 0x7e 到缓冲区
			writer.WriteByte(binary.IdentityBit)
		default:
			return nil, fmt.Errorf("0x7d后面的字符是不符合规则的")
		}

	}

	return writer.Bytes(), nil
}

// Escape 转义
func Escape(data []byte) ([]byte, error) {
	var writer bytes.Buffer

	data = binary.EliminateIdentityBit(data)

	var reader = bytes.NewReader(data)

	// 先转换转义7d的
	for {
		// 读取每一个字节
		currentByte, err := reader.ReadByte()

		if err != nil {
			// 没有更多输入的时候就停止循环
			if err == io.EOF {
				break
			}

			// 读取异常，返回异常
			return nil, err
		}

		if currentByte == binary.Escape {
			writer.WriteByte(binary.Escape)
			writer.WriteByte(0x01)
		} else if currentByte == binary.IdentityBit {
			writer.WriteByte(binary.Escape)
			writer.WriteByte(0x02)
		} else {
			writer.WriteByte(currentByte)
		}
	}
	return writer.Bytes(), nil
}
