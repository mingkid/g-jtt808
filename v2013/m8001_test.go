package v2013

import (
	"encoding/hex"
	"jtt808/msg"
	"testing"
)

func TestM8001_Encode(t *testing.T) {
	var m M8001
	m.SetSerialNumber(321)
	m.SetMsgID(msg.TermLocationRepose)
	m.SetResult(M8001Fail)
	b, _ := m.Encode()

	if r := hex.EncodeToString(b); r != "0141020001" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "0141020001", r)
	}
}
