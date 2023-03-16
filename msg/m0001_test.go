package msg

import (
	"github.com/mingkid/g-jtt808/msg/common"
	"testing"
)

// TestM0001_Decode 测试终端通用应答解码
func TestM0001_Decode(t *testing.T) {
	// 测试数据
	b := []byte{
		0x00, 0x01, // 应答流水号
		0x81, 0x00, // 应答消息ID
		0x00,       // 处理结果
		0x00, 0x00, // 错误代码
	}

	// 解码
	var m M0001
	if err := m.Decode(b); err != nil {
		t.Fatal(err.Error())
	}

	// 测试流水号
	if m.SerialNo() != 0x0001 {
		t.Fatalf("SerialNo() want 0x0001, got 0x%04X", m.SerialNo())
	}

	// 测试消息ID
	if m.MsgID() != common.TermRegResp {
		t.Fatalf("MsgID() want 0x%04X, got 0x%04X", common.TermRegResp, m.MsgID())
	}

	// 测试处理结果
	if m.Result() != M0001_OK {
		t.Fatalf("Result() want 0x%02X, got 0x%02X", M0001_OK, m.Result())
	}

	// 测试错误代码
	if m.ErrorCode() != 0x0000 {
		t.Fatalf("ErrorCode() want 0x0000, got 0x%04X", m.ErrorCode())
	}
}
