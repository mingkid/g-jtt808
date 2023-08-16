package common

// StreamType 码流类型
type StreamType byte

const (
	// StreamTypeMainOrSub 码流类型：主码流或子码流
	StreamTypeMainOrSub StreamType = iota

	// StreamTypeMain 码流类型：主码流
	StreamTypeMain

	// StreamTypeSub 码流类型：子码流
	StreamTypeSub
)
