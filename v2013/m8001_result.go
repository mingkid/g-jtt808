package v2013

type M8001Result uint8

const (
	// M8001Success 成功
	M8001Success M8001Result = iota

	// M8001Fail 失败
	M8001Fail

	// M8001MsgErr 消息有误
	M8001MsgErr

	// M8001Unsupported 不支持
	M8001Unsupported

	// M8001WarnConfirm 报警处理确认
	M8001WarnConfirm
)
