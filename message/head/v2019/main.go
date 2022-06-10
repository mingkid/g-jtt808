package v2019

type MsgHead struct {
	RawID          uint16 `jtt808:""`
	BodyProperty   MsgBodyProperty
	RawPhone       string `jtt808:"6,bcd"`
	RawSerialNum   uint16 `jtt808:""`
	PackagePacking *MsgPackagePacking
}
