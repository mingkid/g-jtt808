package msg

import (
	"encoding/hex"
	"github.com/mingkid/g-jtt808/msg/common"
	"testing"
)

func TestHead_Decode(t *testing.T) {
	h := Head{}
	_ = h.Decode([]byte{1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128})
	if h.MsgID() != common.TermRegister {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", common.TermRegister, h.MsgID())
	}
	if h.SerialNum() != 128 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 128, h.SerialNum())
	}
	if h.Phone() != "013051192675" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051192675", h.Phone())
	}
	if h.BodyProperty().IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, h.BodyProperty().IsSub())
	}
	if h.BodyProperty().BodyLength() != 39 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 39, h.BodyProperty().BodyLength())
	}
	if h.BodyProperty().EncryptType() != common.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, h.BodyProperty().EncryptType())
	}
}

func TestHead_Encode(t *testing.T) {
	var h Head
	h.SetID(common.TermRegResp)
	h.SetPhone("13680179679")
	h.SetSerialNum(123)
	b, _ := h.Encode(13)
	if res := hex.EncodeToString(b); res != "8100000d013680179679007b" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "8100000d013680179679007b", res)
	}
}
