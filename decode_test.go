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
	if msg.Body.Latitude() != 41.912090 {
		t.Fatalf("消息包纬度解析错误，应为%f，实际为%f", 41.912090, msg.Body.Latitude())
	}
	if msg.Body.Longitude() != 122.665674 {
		t.Fatalf("消息包经度解析错误，应为%f，实际为%f", 122.665674, msg.Body.Longitude())
	}
	if msg.Body.Altitude() != 0 {
		t.Fatalf("消息包高程解析错误，应为%d，实际为%d", 0, msg.Body.Altitude())
	}
	if msg.Body.Direction() != 150 {
		t.Fatalf("消息包方向解析错误，应为%d，实际为%d", 0, msg.Body.Direction())
	}
	if time, _ := time.ParseInLocation("20060102150405", "20220324122535", time.Local); msg.Body.Time() != time {
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
	if r, _ := msg.Body.Extras(); r.Mileage() != 249.0 {
		t.Fatalf("消息包附加消息里程解析错误，应为%f，实际为%f", 249.0, r.Mileage())
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
	if r, _ := msg.Body.Extras(); r.Speed() != 11.6 {
		t.Fatalf("消息包附加消息速度解析错误，应为%f，实际为%f", 11.6, r.Speed())
	}
}

func TestDecoder_M0002(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, *interface{}]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 0, 2, 0, 0, 1, 48, 81, 25, 38, 117, 1, 56, 17, 126})
	if msg.Head.ID() != msgcomm.TermHeartbeat {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermHeartbeat, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 312 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 312, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051192675" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051192675", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 0 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 0, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}
	if msg.Body != nil {
		t.Fatalf("消息包消息体解析错误，应为%v，实际为%v", nil, msg.Body)
	}
}

func TestDecoder_V2013_M0704(t *testing.T) {
	var msg message.Message[hv2013.MsgHead, bv2013.M0704]
	d := NewDecoder(&msg)
	_ = d.Decode([]byte{126, 7, 4, 2, 71, 1, 48, 81, 33, 131, 99, 0, 5, 0, 10, 1, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 170, 192, 7, 81, 204, 109, 0, 0, 1, 7, 1, 57, 34, 6, 2, 22, 50, 87, 1, 4, 0, 0, 82, 189, 43, 4, 5, 152, 5, 152, 48, 1, 24, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 7, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 173, 207, 7, 81, 199, 191, 0, 0, 0, 185, 1, 79, 34, 6, 2, 22, 51, 23, 1, 4, 0, 0, 82, 190, 43, 4, 5, 162, 5, 162, 48, 1, 24, 49, 1, 27, 0, 4, 0, 206, 5, 162, 3, 2, 0, 185, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 41, 7, 81, 199, 162, 0, 0, 0, 183, 0, 2, 34, 6, 2, 22, 51, 25, 1, 4, 0, 0, 82, 190, 43, 4, 5, 142, 5, 142, 48, 1, 23, 49, 1, 27, 0, 4, 0, 206, 5, 142, 3, 2, 0, 183, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 88, 7, 81, 199, 177, 0, 0, 0, 203, 0, 23, 34, 6, 2, 22, 51, 32, 1, 4, 0, 0, 82, 190, 43, 4, 5, 142, 5, 142, 48, 1, 23, 49, 1, 27, 0, 4, 0, 206, 5, 142, 3, 2, 0, 203, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 181, 7, 81, 200, 4, 0, 0, 0, 240, 0, 40, 34, 6, 2, 22, 51, 34, 1, 4, 0, 0, 82, 190, 43, 4, 5, 152, 5, 152, 48, 1, 22, 49, 1, 27, 0, 4, 0, 206, 5, 152, 3, 2, 0, 240, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 175, 76, 7, 81, 200, 182, 0, 0, 1, 36, 0, 41, 34, 6, 2, 22, 51, 37, 1, 4, 0, 0, 82, 190, 43, 4, 5, 131, 5, 131, 48, 1, 22, 49, 1, 28, 0, 4, 0, 206, 5, 131, 3, 2, 1, 36, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 182, 173, 7, 81, 208, 224, 0, 0, 0, 0, 0, 41, 34, 6, 2, 22, 51, 84, 1, 4, 0, 0, 82, 192, 43, 4, 5, 131, 5, 131, 48, 1, 28, 49, 1, 28, 0, 4, 0, 206, 5, 131, 3, 2, 0, 0, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 190, 79, 7, 81, 217, 142, 0, 0, 1, 52, 0, 39, 34, 6, 2, 22, 52, 36, 1, 4, 0, 0, 82, 194, 43, 4, 5, 152, 5, 152, 48, 1, 24, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 52, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 196, 55, 7, 81, 224, 4, 0, 0, 1, 9, 0, 38, 34, 6, 2, 22, 52, 84, 1, 4, 0, 0, 82, 196, 43, 4, 5, 152, 5, 152, 48, 1, 28, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 9, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 202, 73, 7, 81, 231, 50, 0, 0, 1, 6, 0, 41, 34, 6, 2, 22, 53, 36, 1, 4, 0, 0, 82, 197, 43, 4, 5, 142, 5, 142, 48, 1, 29, 49, 1, 29, 0, 4, 0, 206, 5, 142, 3, 2, 1, 6, 239, 126})

	// 消息头
	if msg.Head.ID() != msgcomm.TermLocationBatch {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msgcomm.TermLocationBatch, msg.Head.ID())
	}
	if msg.Head.SerialNum() != 5 {
		t.Fatalf("消息包流水号解析错误，应为%d，实际为%d", 5, msg.Head.SerialNum())
	}
	if msg.Head.Phone() != "013051218363" {
		t.Fatalf("消息包手机号解析错误，应为%s，实际为%s", "013051218363", msg.Head.Phone())
	}
	if msg.Head.BodyProperty.IsSub() != false {
		t.Fatalf("消息包分包解析错误，应为%t，实际为%t", false, msg.Head.BodyProperty.IsSub())
	}
	if msg.Head.BodyProperty.BodyLength() != 583 {
		t.Fatalf("消息包消息长度解析错误，应为%d，实际为%d", 583, msg.Head.BodyProperty.BodyLength())
	}
	if msg.Head.BodyProperty.EncryptType() != msgcomm.None {
		t.Fatalf("消息包加密解析错误，应为%d，实际为%d", 39, msg.Head.BodyProperty.EncryptType())
	}

	// 消息体
	if res := msg.Body.Count(); res != 10 {
		t.Fatalf("消息包数据项个数解析错误，应为%d，实际为%d", 10, res)
	}
	if res := msg.Body.Type(); res != 1 {
		t.Fatalf("消息包位置数据类型解析错误，应为%d，实际为%d", 1, res)
	}

	// 位置汇报数据项
	items, _ := msg.Body.Items()
	if length := len(items); length != 10 {
		t.Fatalf("消息包汇报数据行解析错误，应为%d，实际为%d", 10, length)
	}
	locationA := items[0]
	if locationA.Warn != 0 {
		t.Fatalf("消息包报警标志解析错误，应为%d，实际为%d", 0, locationA.Warn)
	}
	if locationA.Status != 786435 {
		t.Fatalf("消息包状态解析错误，应为%d，实际为%d", 786435, locationA.Status)
	}
	if locationA.RawLatitude != 41986752 {
		t.Fatalf("消息包纬度解析错误，应为%d，实际为%d", 41986752, locationA.RawLatitude)
	}
	if locationA.RawLongitude != 122801261 {
		t.Fatalf("消息包经度解析错误，应为%d，实际为%d", 122801261, locationA.RawLongitude)
	}
	if locationA.RawAltitude != 0 {
		t.Fatalf("消息包高程解析错误，应为%d，实际为%d", 0, locationA.RawAltitude)
	}
	if locationA.RawSpeed != 263 {
		t.Fatalf("消息包速度解析错误，应为%d，实际为%d", 263, locationA.RawSpeed)
	}
	if locationA.RawDirection != 313 {
		t.Fatalf("消息包方向解析错误，应为%d，实际为%d", 313, locationA.RawDirection)
	}
	if locationA.RawTime != "220602163257" {
		t.Fatalf("消息包时间解析错误，应为%s，实际为%s", "220602163257", locationA.RawTime)
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

func TestName(t *testing.T) {
	//ConvHEX([]byte{126, 0, 3, 0, 0, 1, 48, 81, 25, 34, 135, 0, 124, 163, 126})                                                                                                                                                       // 0003
	//ConvHEX([]byte{126, 129, 0, 0, 35, 1, 48, 80, 137, 129, 115, 0, 0, 0, 32, 0, 56, 102, 48, 51, 98, 56, 48, 51, 97, 100, 97, 97, 56, 101, 50, 102, 101, 53, 101, 101, 51, 50, 98, 48, 49, 102, 51, 100, 55, 57, 97, 52, 200, 126}) // 8100
	//ConvHEX([]byte{126, 128, 1, 0, 5, 1, 48, 80, 137, 129, 115, 0, 0, 1, 88, 1, 2, 1, 197, 126})                                                                                                                                     // 8001
	ConvHEX([]byte{126, 7, 4, 2, 71, 1, 48, 81, 33, 131, 99, 0, 5, 0, 10, 1, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 170, 192, 7, 81, 204, 109, 0, 0, 1, 7, 1, 57, 34, 6, 2, 22, 50, 87, 1, 4, 0, 0, 82, 189, 43, 4, 5, 152, 5, 152, 48, 1, 24, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 7, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 173, 207, 7, 81, 199, 191, 0, 0, 0, 185, 1, 79, 34, 6, 2, 22, 51, 23, 1, 4, 0, 0, 82, 190, 43, 4, 5, 162, 5, 162, 48, 1, 24, 49, 1, 27, 0, 4, 0, 206, 5, 162, 3, 2, 0, 185, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 41, 7, 81, 199, 162, 0, 0, 0, 183, 0, 2, 34, 6, 2, 22, 51, 25, 1, 4, 0, 0, 82, 190, 43, 4, 5, 142, 5, 142, 48, 1, 23, 49, 1, 27, 0, 4, 0, 206, 5, 142, 3, 2, 0, 183, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 88, 7, 81, 199, 177, 0, 0, 0, 203, 0, 23, 34, 6, 2, 22, 51, 32, 1, 4, 0, 0, 82, 190, 43, 4, 5, 142, 5, 142, 48, 1, 23, 49, 1, 27, 0, 4, 0, 206, 5, 142, 3, 2, 0, 203, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 174, 181, 7, 81, 200, 4, 0, 0, 0, 240, 0, 40, 34, 6, 2, 22, 51, 34, 1, 4, 0, 0, 82, 190, 43, 4, 5, 152, 5, 152, 48, 1, 22, 49, 1, 27, 0, 4, 0, 206, 5, 152, 3, 2, 0, 240, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 175, 76, 7, 81, 200, 182, 0, 0, 1, 36, 0, 41, 34, 6, 2, 22, 51, 37, 1, 4, 0, 0, 82, 190, 43, 4, 5, 131, 5, 131, 48, 1, 22, 49, 1, 28, 0, 4, 0, 206, 5, 131, 3, 2, 1, 36, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 182, 173, 7, 81, 208, 224, 0, 0, 0, 0, 0, 41, 34, 6, 2, 22, 51, 84, 1, 4, 0, 0, 82, 192, 43, 4, 5, 131, 5, 131, 48, 1, 28, 49, 1, 28, 0, 4, 0, 206, 5, 131, 3, 2, 0, 0, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 190, 79, 7, 81, 217, 142, 0, 0, 1, 52, 0, 39, 34, 6, 2, 22, 52, 36, 1, 4, 0, 0, 82, 194, 43, 4, 5, 152, 5, 152, 48, 1, 24, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 52, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 196, 55, 7, 81, 224, 4, 0, 0, 1, 9, 0, 38, 34, 6, 2, 22, 52, 84, 1, 4, 0, 0, 82, 196, 43, 4, 5, 152, 5, 152, 48, 1, 28, 49, 1, 29, 0, 4, 0, 206, 5, 152, 3, 2, 1, 9, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 202, 73, 7, 81, 231, 50, 0, 0, 1, 6, 0, 41, 34, 6, 2, 22, 53, 36, 1, 4, 0, 0, 82, 197, 43, 4, 5, 142, 5, 142, 48, 1, 29, 49, 1, 29, 0, 4, 0, 206, 5, 142, 3, 2, 1, 6, 239, 126})
	//ConvHEX([]byte{126, 2, 0, 0, 56, 1, 48, 81, 33, 131, 99, 0, 71, 0, 0, 0, 0, 0, 12, 0, 3, 2, 128, 171, 240, 7, 82, 56, 111, 0, 0, 1, 45, 1, 9, 34, 6, 2, 34, 86, 85, 1, 4, 0, 0, 84, 10, 43, 4, 5, 172, 5, 172, 48, 1, 30, 49, 1, 25, 0, 4, 0, 206, 5, 172, 3, 2, 1, 45, 23, 126}) // 0200
	//ConvHEX([]byte{126, 2, 0, 0, 86, 1, 48, 81, 35, 32, 0, 0, 2, 0, 0, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 4, 0, 0, 163, 132, 43, 4, 5, 19, 5, 19, 48, 1, 31, 49, 1, 0, 0, 4, 0, 206, 5, 19, 3, 2, 0, 0, 0, 12, 0, 178, 137, 134, 4, 162, 25, 33, 194, 86, 50, 148, 235, 14, 0, 12, 0, 178, 137, 134, 4, 162, 25, 33, 194, 86, 50, 148, 18, 126}) // 异常的0200
}
