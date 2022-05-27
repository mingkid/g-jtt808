package v2013

import (
	"testing"

	"github.com/mingkid/g-jtt808/message/msgcomm"
)

func TestMsgHead_ID(t *testing.T) {
	var h MsgHead
	h.SetID(msgcomm.PlatformCommResp)
	if h.ID() != msgcomm.PlatformCommResp {
		t.Fatalf("设置消息ID错误，应为%d，实际为%d", msgcomm.PlatformCommResp, h.ID())
	}
}

func TestMsgHead_SerialNum(t *testing.T) {
	var h MsgHead
	h.SetSerialNum(123)
	if h.SerialNum() != 123 {
		t.Fatalf("设置消息流水号错误，应为%d，实际为%d", 123, h.SerialNum())
	}
}

func TestMsgHead_Phone(t *testing.T) {
	var h MsgHead
	h.SetPhone("12345678901")
	if h.Phone() != "12345678901" {
		t.Fatalf("设置终端手机号错误，应为%s，实际为%s", "12345678901", h.Phone())
	}
}
