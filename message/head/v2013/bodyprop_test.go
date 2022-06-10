package v2013

import (
	"testing"

	"github.com/mingkid/g-jtt808/message/msgcomm"
)

func TestMsgBodyProperty_IsSub(t *testing.T) {
	var p MsgBodyProperty
	if p.SetIsSub(); p.Raw != 8192 || p.IsSub() == false {
		t.Fatalf("设置分包错误，应为%d，实际为%d", 8192, p.Raw)
	}
}

func TestMsgBodyProperty_BodyLength(t *testing.T) {
	var p MsgBodyProperty

	// 长度设置正常范围
	_ = p.SetBodyLength(8)
	if p.Raw != 8 || p.BodyLength() != 8 {
		t.Fatalf("设置数据包长度错误，应为%d，实际为%d", 8, p)
	}

	// 长度设置超过允许范围
	if err := p.SetBodyLength(1024); err == nil || err != BodyLengthTooLarge {
		t.Fatalf("设置数据包长度过长没抛出异常")
	}
}

func TestMsgBodyProperty_SetEncrypt(t *testing.T) {
	var p MsgBodyProperty
	if p.SetEncrypt(); p.Raw != 1024 || p.EncryptType() != msgcomm.RSA {
		t.Fatalf("设置加密错误，应为%d，实际为%d", 1024, p.Raw)
	}
}
