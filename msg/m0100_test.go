package msg

import "testing"

func TestM0100_Decode(t *testing.T) {
	m := M0100{}
	_ = m.Decode([]byte{0, 44, 1, 44, 55, 48, 49, 49, 49, 66, 83, 74, 45, 65, 54, 66, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 49, 57, 50, 54, 55, 53, 0, 0, 0})
	if m.ProvincialID() != 44 {
		t.Fatalf("消息包省域ID解析错误，应为%d，实际为%d", 44, m.ProvincialID())
	}
	if m.CityID() != 300 {
		t.Fatalf("消息包市县域ID解析错误，应为%d，实际为%d", 300, m.CityID())
	}
	if res, _ := m.ManufacturerID(); res != "70111" {
		t.Fatalf("消息包制造商ID解析错误，应为%s，实际为%s", "70111", res)
	}
	if res, _ := m.TermModel(); res != "BSJ-A6BD" {
		t.Fatalf("消息包解析错误，应为%s，实际为%s", "BSJ-A6BD", res)
	}
	if res, _ := m.TermID(); res != "1192675" {
		t.Fatalf("消息包终端ID解析错误，应为%s，实际为%s", "1192675", res)
	}
	if m.LicensePlateColor() != 0 {
		t.Fatalf("消息包车牌颜色解析错误，应为%d，实际为%d", 0, m.LicensePlateColor())
	}
	if res, _ := m.LicensePlate(); res != "" {
		t.Fatalf("消息包车牌标识解析错误，应为%s，实际为%s", "", res)
	}
}
