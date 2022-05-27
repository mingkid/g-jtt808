package jtt808

import (
	"encoding/hex"
	"testing"

	"github.com/mingkid/g-jtt808/message"
	bv2013 "github.com/mingkid/g-jtt808/message/body/v2013"
	hv2013 "github.com/mingkid/g-jtt808/message/head/v2013"
	"github.com/mingkid/g-jtt808/message/msgcomm"
)

func TestEncoder_M8001(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M8001]

	// 不分包
	msg.Head.SetID(msgcomm.PlatformCommResp)
	msg.Head.SetPhone("13680179679")
	msg.Head.SetSerialNum(123)
	msg.Body.SetMsgID(msgcomm.TermLocationRepose).SetSerialNumber(321).SetResult(bv2013.M8001Fail)

	e := NewEncoder(&msg)
	b, _ := e.Encode(message.CalcChecksum)
	if r := hex.EncodeToString(b); r != "7e80010005013680179679007b0141020001f37e" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "7e80010005013680179679007b0141020001f37e", r)
	}

	//	分包
	msg.Head.BodyProperty = 0
	msg.Head.BodyProperty = msg.Head.BodyProperty.SetIsSub().(hv2013.MsgBodyProperty)
	msg.Head.PackagePacking = &hv2013.MsgPackagePacking{}
	msg.Head.PackagePacking.SetTotal(1).SetSerialNum(1)

	e = NewEncoder(&msg)
	b, _ = e.Encode(message.CalcChecksum)
	if r := hex.EncodeToString(b); r != "7e80012005013680179679007b000100010141020001d37e" {
		t.Fatalf("消息包分包组包错误，应为%s，实际为%s", "7e80012005013680179679007b000100010141020001d37e", r)
	}
}

func TestEncoder_M8100(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M8100]

	msg.Head.SetID(msgcomm.TermRegResp)
	msg.Head.SetPhone("13680179679")
	msg.Head.SetSerialNum(123)
	msg.Body.SetToken("1234567890")
	msg.Body.SetSerialNumber(321)
	msg.Body.SetResult(bv2013.Success)

	e := NewEncoder(&msg)
	b, _ := e.Encode(message.CalcChecksum)
	if r := hex.EncodeToString(b); r != "7e8100000d013680179679007b01410031323334353637383930f97e" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "7e8100000d013680179679007b01410031323334353637383930f97e", r)
	}
}
