package common

type MsgID uint16

const (
	// TermCommResp 终端通用应答
	TermCommResp MsgID = 0x0001

	// TermRegister 终端注册
	TermRegister = 0x0100

	// TermRegResp 终端注册应答
	TermRegResp = 0x8100

	// TermLocationRepose 终端位置汇报
	TermLocationRepose = 0x0200

	// TermLocationBatch 终端位置批量上传
	TermLocationBatch = 0x0704

	// TermAuth 终端鉴权
	TermAuth = 0x0102

	// PlatformCommResp 平台通用应答
	PlatformCommResp MsgID = 0x8001

	// TermHeartbeat 终端心跳
	TermHeartbeat MsgID = 0x0002
)
