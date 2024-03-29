package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
)

type M8100 struct {
	serialNumber uint16
	result       uint8
	token        string
}

func (m M8100) SerialNumber() uint16 {
	return m.serialNumber
}

func (m *M8100) SetSerialNumber(num uint16) *M8100 {
	m.serialNumber = num
	return m
}

func (m M8100) Result() M8100Result {
	return M8100Result(m.result)
}

func (m *M8100) SetResult(num M8100Result) *M8100 {
	m.result = uint8(num)
	return m
}

func (m M8100) Token() string {
	return m.token
}

func (m *M8100) SetToken(token string) *M8100 {
	m.token = token
	return m
}

func (m M8100) Encode() (res []byte, err error) {
	w := binary.NewWriter()

	if err = w.WriteUint16(m.serialNumber); err != nil {
		return nil, fmt.Errorf("终端注册响应消息流水号编码失败")
	}
	if err = w.WriteUint8(m.result); err != nil {
		return nil, fmt.Errorf("终端注册响应结果编码失败")
	}
	if err = w.WriteString(m.token); err != nil {
		return nil, fmt.Errorf("终端注册响应鉴权码编码失败")
	}

	res = w.Bytes()
	return
}
