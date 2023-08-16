package msg

// M9101DataType 数据类型枚举
type M9101DataType uint8

const (
	// DataTypeAudioVideo 音视频
	DataTypeAudioVideo M9101DataType = iota

	// DataTypeVideo 视频
	DataTypeVideo

	// DataTypeTwoWayIntercom 双向对讲
	DataTypeTwoWayIntercom

	// DataTypeMonitor 监听
	DataTypeMonitor

	// DataTypeCenterBroadcast 中心广播
	DataTypeCenterBroadcast

	// DataTypePassThrough 透传
	DataTypePassThrough
)
