package msg

// ReplayControl 回放控制类型枚举
type ReplayControl uint8

const (
	// ReplayControlStart 表示开始回放
	ReplayControlStart ReplayControl = iota

	// ReplayControlPause 表示暂停回放
	ReplayControlPause

	// ReplayControlStop 表示停止回放
	ReplayControlStop

	// ReplayControlFastForward 表示快进回放
	ReplayControlFastForward

	// ReplayControlKeyFrameFastRewind 表示关键帧快退回放
	ReplayControlKeyFrameFastRewind

	// ReplayControlDrag 表示拖动回放
	ReplayControlDrag

	// ReplayControlKeyFramePlay 表示关键帧播放
	ReplayControlKeyFramePlay
)
