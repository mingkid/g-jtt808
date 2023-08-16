package msg

// M9101StreamType 码流类型
type M9101StreamType byte

const (
	// StreamTypeMain 主码流
	StreamTypeMain M9101StreamType = iota

	// StreamTypeSub 子码流
	StreamTypeSub
)
