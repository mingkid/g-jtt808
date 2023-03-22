package msg

import (
	"testing"
	"time"
)

func TestM0200_Decode(t *testing.T) {
	var m M0200
	_ = m.Decode([]byte{0, 0, 0, 0, 0, 12, 0, 3, 2, 127, 135, 26, 7, 79, 186, 202, 0, 0, 0, 116, 0, 150, 34, 3, 36, 18, 37, 53, 1, 4, 0, 0, 9, 186, 43, 4, 5, 193, 5, 193, 48, 1, 29, 49, 1, 29, 0, 4, 0, 206, 5, 193, 3, 2, 0, 116})
	if m.Latitude() != 41.912090 {
		t.Fatalf("消息包纬度解析错误，应为%f，实际为%f", 41.912090, m.Latitude())
	}
	if m.Longitude() != 122.665674 {
		t.Fatalf("消息包经度解析错误，应为%f，实际为%f", 122.665674, m.Longitude())
	}
	if m.Altitude() != 0 {
		t.Fatalf("消息包高程解析错误，应为%d，实际为%d", 0, m.Altitude())
	}
	if m.Direction() != 150 {
		t.Fatalf("消息包方向解析错误，应为%d，实际为%d", 0, m.Direction())
	}
	a, _ := time.ParseInLocation("20060102150405", "20220324122535", time.Local)
	b, err := m.Time()
	if err != nil {
		t.Error(err)
	}
	if a != b {
		t.Fatalf("消息包时间解析错误，应为%v，实际为%v", a, b)
	}
	if m.Status().ACCOn() != true {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", true, m.Status().ACCOn())
	}
	if m.Status().PositionOn() != true {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", true, m.Status().PositionOn())
	}
	if m.Status().IsSLat() != false {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", false, m.Status().IsSLat())
	}
	if m.Status().IsWLong() != false {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", false, m.Status().IsWLong())
	}
	if m.Status().IsOperation() != true {
		t.Fatalf("消息包状态运营解析错误，应为%t，实际为%t", true, m.Status().IsOperation())
	}
	if m.Status().IsEncrypt() != false {
		t.Fatalf("消息包状态经纬度加密解析错误，应为%t，实际为%t", false, m.Status().IsEncrypt())
	}
	if m.Status().IsOilChannelNormal() != true {
		t.Fatalf("消息包状态油路解析错误，应为%t，实际为%t", true, m.Status().IsOilChannelNormal())
	}
	if m.Status().IsCircuitNormal() != true {
		t.Fatalf("消息包状态电路解析错误，应为%t，实际为%t", true, m.Status().IsCircuitNormal())
	}
	if m.Status().IsDoorLock() != false {
		t.Fatalf("消息包状态车门锁解析错误，应为%t，实际为%t", false, m.Status().IsDoorLock())
	}
	if m.Status().IsFrontDoorOpen() != false {
		t.Fatalf("消息包状态车前门（门1）锁解析错误，应为%t，实际为%t", false, m.Status().IsFrontDoorOpen())
	}
	if m.Status().IsMidDoorOpen() != false {
		t.Fatalf("消息包状态车中门（门2）锁解析错误，应为%t，实际为%t", false, m.Status().IsMidDoorOpen())
	}
	if m.Status().IsBackDoorOpen() != false {
		t.Fatalf("消息包状态车后门（门3）锁解析错误，应为%t，实际为%t", false, m.Status().IsBackDoorOpen())
	}
	if m.Status().IsDriveRoomDoorOpen() != false {
		t.Fatalf("消息包状态车驾驶席门（门4）锁解析错误，应为%t，实际为%t", false, m.Status().IsDriveRoomDoorOpen())
	}
	if m.Status().IsElseRoomDoorOpen() != false {
		t.Fatalf("消息包状态车自定义门（门5）锁解析错误，应为%t，实际为%t", false, m.Status().IsElseRoomDoorOpen())
	}
	if m.Status().UsedGPS() != true {
		t.Fatalf("消息包状态使用GPS卫星解析错误，应为%t，实际为%t", true, m.Status().UsedGPS())
	}
	if m.Status().UsedBPS() != true {
		t.Fatalf("消息包状态使用北斗卫星解析错误，应为%t，实际为%t", true, m.Status().UsedBPS())
	}
	if m.Status().UsedGalileo() != false {
		t.Fatalf("消息包状态使用Galileo卫星解析错误，应为%t，实际为%t", false, m.Status().UsedGalileo())
	}
	if m.Status().UsedGLONASS() != false {
		t.Fatalf("消息包状态使用GLONASS卫星解析错误，应为%t，实际为%t", false, m.Status().UsedGLONASS())
	}
	if r, _ := m.Extras(); r.Mileage() != 249.0 {
		t.Fatalf("消息包附加消息里程解析错误，应为%f，实际为%f", 249.0, r.Mileage())
	}
	if r, _ := m.Extras(); r.Analog() != 96536001 {
		t.Fatalf("消息包附加消息模拟量解析错误，应为%d，实际为%d", 96536001, r.Analog())
	}
	if r, _ := m.Extras(); r.SignalStrength() != 29 {
		t.Fatalf("消息包附加消息信号强度解析错误，应为%d，实际为%d", 29, r.SignalStrength())
	}
	if r, _ := m.Extras(); r.GNSSQty() != 29 {
		t.Fatalf("消息包附加消息GNSS卫星数量解析错误，应为%d，实际为%d", 29, r.GNSSQty())
	}
	if r, _ := m.Extras(); r.Speed() != 11.6 {
		t.Fatalf("消息包附加消息速度解析错误，应为%f，实际为%f", 11.6, r.Speed())
	}
}
