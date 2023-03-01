package v2013

type M0001Result uint8

const (
	M0001_OK          M0001Result = iota // 成功/确认
	M0001_FAIL                           // 失败
	M0001_MSG_ERROR                      // 消息有误
	M0001_UNKNOWN                        // 不支持
	M001_WARNING_DEAL                    // 报警处理确认
)
