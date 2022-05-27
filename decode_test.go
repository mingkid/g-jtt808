package jtt808

import (
	"testing"
	"time"

	"github.com/mingkid/g-jtt808/message"
	bv2013 "github.com/mingkid/g-jtt808/message/body/v2013"
	hv2013 "github.com/mingkid/g-jtt808/message/head/v2013"
	"github.com/mingkid/g-jtt808/message/msgcomm"
)

func TestDecoder_M0100(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M0100]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128, 0, 44, 1, 44, 55, 48, 49, 49, 49, 66, 83, 74, 45, 65, 54, 66, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 49, 57, 50, 54, 55, 53, 0, 0, 0, 131, 126})
	if msg.Head.ID() != msgcomm.TermRegister {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermRegister, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 128 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 128, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051192675" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051192675", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 39 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}
	if msg.Body.ProvincialID() != 44 {
		t.Fatalf("消息包省域ID解析错误，应为%d，实际为%d", 44, msg.Body.ProvincialID())
	}
	if msg.Body.CityID() != 300 {
		t.Fatalf("消息包市县域ID解析错误，应为%d，实际为%d", 300, msg.Body.CityID())
	}
	if msg.Body.ManufacturerID() != "70111" {
		t.Fatalf("消息包制造商ID解析错误，应为%s，实际为%s", "70111", msg.Body.ManufacturerID())
	}
	if msg.Body.TermModel() != "BSJ-A6BD" {
		t.Fatalf("消息包解析错误，应为%s，实际为%s", "BSJ-A6BD", msg.Body.TermModel())
	}
	if msg.Body.TermID() != "1192675" {
		t.Fatalf("消息包终端ID解析错误，应为%s，实际为%s", "1192675", msg.Body.TermID())
	}
	if msg.Body.LicensePlateColor() != 0 {
		t.Fatalf("消息包车牌颜色解析错误，应为%d，实际为%d", 0, msg.Body.LicensePlateColor())
	}
	if msg.Body.LicensePlate() != "" {
		t.Fatalf("消息包车牌标识解析错误，应为%s，实际为%s", "", msg.Body.LicensePlate())
	}
}

func TestDecoder_M0102(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M0102]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 1, 2, 0, 32, 1, 48, 81, 25, 99, 56, 0, 106, 55, 54, 50, 54, 102, 53, 49, 57, 56, 48, 53, 51, 102, 98, 50, 99, 48, 49, 100, 100, 48, 101, 98, 101, 97, 100, 101, 54, 48, 99, 102, 51, 109, 126})
	if msg.Head.ID() != msgcomm.TermAuth {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermAuth, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 106 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 106, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051196338" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051196338", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 32 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 32, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}
	if msg.Body.Token() != "7626f5198053fb2c01dd0ebeade60cf3" {
		t.Fatalf("消息包鉴权码解析错误，应为%s，实际为%s", "7626f5198053fb2c01dd0ebeade60cf3", msg.Body.Token())
	}
}

func TestDecoder_M0200(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M0200]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 2, 0, 0, 56, 1, 48, 81, 25, 38, 117, 0, 10, 0, 0, 0, 0, 0, 12, 0, 3, 2, 127, 135, 26, 7, 79, 186, 202, 0, 0, 0, 116, 0, 150, 34, 3, 36, 18, 37, 53, 1, 4, 0, 0, 9, 186, 43, 4, 5, 193, 5, 193, 48, 1, 29, 49, 1, 29, 0, 4, 0, 206, 5, 193, 3, 2, 0, 116, 203, 126})
	if msg.Head.ID() != msgcomm.TermLocationRepose {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermLocationRepose, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 10 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 10, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051192675" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051192675", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 56 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 56, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}
	if msg.Body.Latitude() != 41912090 {
		t.Fatalf("消息包纬度解析错误，应为%d，实际为%d", 41912090, msg.Body.Latitude())
	}
	if msg.Body.Longitude() != 122665674 {
		t.Fatalf("消息包经度解析错误，应为%d，实际为%d", 122665674, msg.Body.Longitude())
	}
	if msg.Body.Altitude() != 0 {
		t.Fatalf("消息包高程解析错误，应为%d，实际为%d", 0, msg.Body.Altitude())
	}
	if msg.Body.Direction() != 150 {
		t.Fatalf("消息包方向解析错误，应为%d，实际为%d", 0, msg.Body.Direction())
	}
	if time, _ := time.Parse(time.RFC3339, "2022-03-24T12:25:35"); msg.Body.Time() != time {
		t.Fatalf("消息包时间解析错误，应为%v，实际为%v", time, msg.Body.Time())
	}
	if msg.Body.Status.ACCOn() != true {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", true, msg.Body.Status.ACCOn())
	}
	if msg.Body.Status.PositionOn() != true {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", true, msg.Body.Status.PositionOn())
	}
	if msg.Body.Status.IsSLat() != false {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", false, msg.Body.Status.IsSLat())
	}
	if msg.Body.Status.IsWLong() != false {
		t.Fatalf("消息包状态ACC解析错误，应为%t，实际为%t", false, msg.Body.Status.IsWLong())
	}
	if msg.Body.Status.IsOperation() != true {
		t.Fatalf("消息包状态运营解析错误，应为%t，实际为%t", true, msg.Body.Status.IsOperation())
	}
	if msg.Body.Status.IsEncrypt() != false {
		t.Fatalf("消息包状态经纬度加密解析错误，应为%t，实际为%t", false, msg.Body.Status.IsEncrypt())
	}
	if msg.Body.Status.IsOilChannelNormal() != true {
		t.Fatalf("消息包状态油路解析错误，应为%t，实际为%t", true, msg.Body.Status.IsOilChannelNormal())
	}
	if msg.Body.Status.IsCircuitNormal() != true {
		t.Fatalf("消息包状态电路解析错误，应为%t，实际为%t", true, msg.Body.Status.IsCircuitNormal())
	}
	if msg.Body.Status.IsDoorLock() != false {
		t.Fatalf("消息包状态车门锁解析错误，应为%t，实际为%t", false, msg.Body.Status.IsDoorLock())
	}
	if msg.Body.Status.IsFrontDoorOpen() != false {
		t.Fatalf("消息包状态车前门（门1）锁解析错误，应为%t，实际为%t", false, msg.Body.Status.IsFrontDoorOpen())
	}
	if msg.Body.Status.IsMidDoorOpen() != false {
		t.Fatalf("消息包状态车中门（门2）锁解析错误，应为%t，实际为%t", false, msg.Body.Status.IsMidDoorOpen())
	}
	if msg.Body.Status.IsBackDoorOpen() != false {
		t.Fatalf("消息包状态车后门（门3）锁解析错误，应为%t，实际为%t", false, msg.Body.Status.IsBackDoorOpen())
	}
	if msg.Body.Status.IsDriveRoomDoorOpen() != false {
		t.Fatalf("消息包状态车驾驶席门（门4）锁解析错误，应为%t，实际为%t", false, msg.Body.Status.IsDriveRoomDoorOpen())
	}
	if msg.Body.Status.IsElseRoomDoorOpen() != false {
		t.Fatalf("消息包状态车自定义门（门5）锁解析错误，应为%t，实际为%t", false, msg.Body.Status.IsElseRoomDoorOpen())
	}
	if msg.Body.Status.UsedGPS() != true {
		t.Fatalf("消息包状态使用GPS卫星解析错误，应为%t，实际为%t", true, msg.Body.Status.UsedGPS())
	}
	if msg.Body.Status.UsedBPS() != true {
		t.Fatalf("消息包状态使用北斗卫星解析错误，应为%t，实际为%t", true, msg.Body.Status.UsedBPS())
	}
	if msg.Body.Status.UsedGalileo() != false {
		t.Fatalf("消息包状态使用Galileo卫星解析错误，应为%t，实际为%t", false, msg.Body.Status.UsedGalileo())
	}
	if msg.Body.Status.UsedGLONASS() != false {
		t.Fatalf("消息包状态使用GLONASS卫星解析错误，应为%t，实际为%t", false, msg.Body.Status.UsedGLONASS())
	}
	if r, _ := msg.Body.Extras(); r.Mileage() != 2490 {
		t.Fatalf("消息包附加消息里程解析错误，应为%d，实际为%d", 2490, r.Mileage())
	}
	if r, _ := msg.Body.Extras(); r.Analog() != 96536001 {
		t.Fatalf("消息包附加消息模拟量解析错误，应为%d，实际为%d", 96536001, r.Analog())
	}
	if r, _ := msg.Body.Extras(); r.SignalStrength() != 29 {
		t.Fatalf("消息包附加消息信号强度解析错误，应为%d，实际为%d", 29, r.SignalStrength())
	}
	if r, _ := msg.Body.Extras(); r.GNSSQty() != 29 {
		t.Fatalf("消息包附加消息GNSS卫星数量解析错误，应为%d，实际为%d", 29, r.GNSSQty())
	}
	if r, _ := msg.Body.Extras(); r.Speed() != 116 {
		t.Fatalf("消息包附加消息速度解析错误，应为%d，实际为%d", 116, r.Speed())
	}
}
