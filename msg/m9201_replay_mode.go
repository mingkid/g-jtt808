package msg

// M9201ReplayMode 回放模式枚举
type M9201ReplayMode uint8

const (
	// M9201ReplayModeNormal 正常回放
	M9201ReplayModeNormal M9201ReplayMode = iota

	// M9201ReplayModeFastForward 快进回放
	M9201ReplayModeFastForward

	// M9201ReplayModeKeyFrameFastRewind 关键帧快退回放
	M9201ReplayModeKeyFrameFastRewind

	// M9201ReplayModeKeyFramePlay 关键帧播放
	M9201ReplayModeKeyFramePlay

	// M9201ReplayModeSingleFrameUpload 单帧上传
	M9201ReplayModeSingleFrameUpload
)
