package msg

import (
	"encoding/hex"
	"testing"
)

func TestM8100_Encode(t *testing.T) {
	var m M8100
	m.SetToken("1234567890")
	m.SetSerialNumber(321)
	m.SetResult(M8100Success)

	b, _ := m.Encode()
	if r := hex.EncodeToString(b); r != "01410031323334353637383930" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "01410031323334353637383930", r)
	}
}
