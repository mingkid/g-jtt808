package msg

import (
	"testing"
	"time"
)

func TestM0704_Decode(t *testing.T) {
	var m M0704
	_ = m.Decode([]byte{0, 10, 1, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 170, 192, 7, 81, 204, 109, 0, 0, 1, 7, 1, 57, 34, 6, 2, 22, 50, 87, 1, 4, 0, 0, 82, 189, 43, 4, 5, 152, 5, 152, 48, 1, 24, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 7, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 173, 207, 7, 81, 199, 191, 0, 0, 0, 185, 1, 79, 34, 6, 2, 22, 51, 23, 1, 4, 0, 0, 82, 190, 43, 4, 5, 162, 5, 162, 48, 1, 24, 49, 1, 27, 0, 4, 0, 206, 5, 162, 3, 2, 0, 185, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 41, 7, 81, 199, 162, 0, 0, 0, 183, 0, 2, 34, 6, 2, 22, 51, 25, 1, 4, 0, 0, 82, 190, 43, 4, 5, 142, 5, 142, 48, 1, 23, 49, 1, 27, 0, 4, 0, 206, 5, 142, 3, 2, 0, 183, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 88, 7, 81, 199, 177, 0, 0, 0, 203, 0, 23, 34, 6, 2, 22, 51, 32, 1, 4, 0, 0, 82, 190, 43, 4, 5, 142, 5, 142, 48, 1, 23, 49, 1, 27, 0, 4, 0, 206, 5, 142, 3, 2, 0, 203, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 181, 7, 81, 200, 4, 0, 0, 0, 240, 0, 40, 34, 6, 2, 22, 51, 34, 1, 4, 0, 0, 82, 190, 43, 4, 5, 152, 5, 152, 48, 1, 22, 49, 1, 27, 0, 4, 0, 206, 5, 152, 3, 2, 0, 240, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 175, 76, 7, 81, 200, 182, 0, 0, 1, 36, 0, 41, 34, 6, 2, 22, 51, 37, 1, 4, 0, 0, 82, 190, 43, 4, 5, 131, 5, 131, 48, 1, 22, 49, 1, 28, 0, 4, 0, 206, 5, 131, 3, 2, 1, 36, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 182, 173, 7, 81, 208, 224, 0, 0, 0, 0, 0, 41, 34, 6, 2, 22, 51, 84, 1, 4, 0, 0, 82, 192, 43, 4, 5, 131, 5, 131, 48, 1, 28, 49, 1, 28, 0, 4, 0, 206, 5, 131, 3, 2, 0, 0, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 190, 79, 7, 81, 217, 142, 0, 0, 1, 52, 0, 39, 34, 6, 2, 22, 52, 36, 1, 4, 0, 0, 82, 194, 43, 4, 5, 152, 5, 152, 48, 1, 24, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 52, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 196, 55, 7, 81, 224, 4, 0, 0, 1, 9, 0, 38, 34, 6, 2, 22, 52, 84, 1, 4, 0, 0, 82, 196, 43, 4, 5, 152, 5, 152, 48, 1, 28, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 9, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 202, 73, 7, 81, 231, 50, 0, 0, 1, 6, 0, 41, 34, 6, 2, 22, 53, 36, 1, 4, 0, 0, 82, 197, 43, 4, 5, 142, 5, 142, 48, 1, 29, 49, 1, 29, 0, 4, 0, 206, 5, 142, 3, 2, 1, 6, 239})

	// 消息体
	if res := m.Count(); res != 10 {
		t.Fatalf("消息包数据项个数解析错误，应为%d，实际为%d", 10, res)
	}
	if res := m.Type(); res != 1 {
		t.Fatalf("消息包位置数据类型解析错误，应为%d，实际为%d", 1, res)
	}

	// 位置汇报数据项
	items, _ := m.Items()
	if length := len(items); length != 10 {
		t.Fatalf("消息包汇报数据行解析错误，应为%d，实际为%d", 10, length)
	}
	locationA := items[0]
	if res := locationA.Warn(); res != 0 {
		t.Fatalf("消息包报警标志解析错误，应为%d，实际为%d", 0, res)
	}
	if res := locationA.Status(); res != 786435 {
		t.Fatalf("消息包状态解析错误，应为%d，实际为%d", 786435, res)
	}
	if res := locationA.Latitude(); res != 41.986752 {
		t.Fatalf("消息包纬度解析错误，应为%f，实际为%f", 41.986752, res)
	}
	if res := locationA.Longitude(); res != 122.801261 {
		t.Fatalf("消息包经度解析错误，应为%f，实际为%f", 122.801261, res)
	}
	if res := locationA.Altitude(); res != 0 {
		t.Fatalf("消息包高程解析错误，应为%d，实际为%d", 0, res)
	}
	if res := locationA.Speed(); res != 26.3 {
		t.Fatalf("消息包速度解析错误，应为%f，实际为%f", 26.3, res)
	}
	if res := locationA.Direction(); res != 313 {
		t.Fatalf("消息包方向解析错误，应为%d，实际为%d", 313, res)
	}
	a, _ := time.ParseInLocation("20060102150405", "20220602163257", time.Local)
	b, _ := locationA.Time()
	if a != b {
		t.Fatalf("消息包时间解析错误，应为%s，实际为%s", a, b)
	}
	if r, _ := locationA.Extras(); r.Mileage() != 2118.1 {
		t.Fatalf("消息包附加消息里程解析错误，应为%f，实际为%f", 2118.1, r.Mileage())
	}
	if r, _ := locationA.Extras(); r.Analog() != 93848984 {
		t.Fatalf("消息包附加消息模拟量解析错误，应为%d，实际为%d", 93848984, r.Analog())
	}
	if r, _ := locationA.Extras(); r.SignalStrength() != 24 {
		t.Fatalf("消息包附加消息信号强度解析错误，应为%d，实际为%d", 24, r.SignalStrength())
	}
	if r, _ := locationA.Extras(); r.GNSSQty() != 29 {
		t.Fatalf("消息包附加消息GNSS卫星数量解析错误，应为%d，实际为%d", 29, r.GNSSQty())
	}
	if r, _ := locationA.Extras(); r.Speed() != 26.3 {
		t.Fatalf("消息包附加消息速度解析错误，应为%f，实际为%f", 26.3, r.Speed())
	}
}
