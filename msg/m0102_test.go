package msg

import "testing"

func TestM0102_Decode(t *testing.T) {
	var m M0102
	_ = m.Decode([]byte{55, 54, 50, 54, 102, 53, 49, 57, 56, 48, 53, 51, 102, 98, 50, 99, 48, 49, 100, 100, 48, 101, 98, 101, 97, 100, 101, 54, 48, 99, 102, 51})

	if res, _ := m.Token(); res != "7626f5198053fb2c01dd0ebeade60cf3" {
		t.Fatalf("消息包鉴权码解析错误，应为%s，实际为%s", "7626f5198053fb2c01dd0ebeade60cf3", res)
	}
}
