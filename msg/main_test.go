package msg

import (
	"encoding/hex"
	"github.com/mingkid/g-jtt808/msg/common"
	"testing"
)

func TestPlantFormMsg_Encode(t *testing.T) {
	var head Head
	head.SetID(common.TermRegResp)
	head.SetPhone("13680179679")
	head.SetSerialNum(123)

	var body M8100
	body.SetToken("1234567890")
	body.SetSerialNumber(321)
	body.SetResult(M8100Success)

	m := PlantFormMsg{
		head: &head,
		body: body,
	}
	b, _ := m.Encode()

	if r := hex.EncodeToString(b); r != "7e8100000d013680179679007b01410031323334353637383930f97e" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "7e8100000d013680179679007b01410031323334353637383930f97e", r)
	}
}
