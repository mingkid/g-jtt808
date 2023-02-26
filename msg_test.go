package jtt808

import (
	"encoding/hex"
	"jtt808/msg"
	v2013 "jtt808/v2013"
	"testing"
)

func TestPlantFormMsg_Encode(t *testing.T) {
	var head v2013.Head
	head.SetID(msg.TermRegResp)
	head.SetPhone("13680179679")
	head.SetSerialNum(123)

	var body v2013.M8100
	body.SetToken("1234567890")
	body.SetSerialNumber(321)
	body.SetResult(v2013.M8100Success)

	m := PlantFormMsg{
		head: &head,
		body: body,
	}
	b, _ := m.Encode()
	if r := hex.EncodeToString(b); r != "7e8100000d01368023179679007b01410031323334353637383930f97e" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "7e8100000d013680179679007b01410031323334353637383930f97e", r)
	}
}
