package common

// AVType 音视频类型
type AVType byte

const (
	// AVTypeAudioVideo 音视频类型：音视频
	AVTypeAudioVideo AVType = iota

	// AVTypeAudio 音视频类型：音频
	AVTypeAudio

	// AVTypeVideo 音视频类型：视频
	AVTypeVideo

	// AVTypeVideoOrAudioVideo 音视频类型：视频或音视频
	AVTypeVideoOrAudioVideo
)
