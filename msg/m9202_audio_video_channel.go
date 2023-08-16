package msg

import (
	"encoding/binary"
	"github.com/mingkid/g-jtt808/msg/common"
)

// M9202AudioVideoChannel M9202协议中的音视频通道信息
type M9202AudioVideoChannel map[byte][]byte

// RealtimeStreamEncodingMode 获取实时流编码模式
func (w M9202AudioVideoChannel) RealtimeStreamEncodingMode() common.RealtimeStreamEncodingMode {
	val, ok := w[0x00]
	if !ok {
		return 0
	}
	return common.RealtimeStreamEncodingMode(val[0])
}

// RealtimeStreamResolution 获取实时流分辨率
func (w M9202AudioVideoChannel) RealtimeStreamResolution() common.RealtimeStreamResolution {
	val, ok := w[0x01]
	if !ok {
		return 0
	}
	return common.RealtimeStreamResolution(val[0])
}

// RealtimeStreamKeyFrameInterval 获取实时流关键帧间隔
func (w M9202AudioVideoChannel) RealtimeStreamKeyFrameInterval() uint16 {
	val, ok := w[0x02]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// RealtimeStreamTargetFrameRate 获取实时流目标帧率
func (w M9202AudioVideoChannel) RealtimeStreamTargetFrameRate() uint8 {
	val, ok := w[0x04]
	if !ok {
		return 0
	}
	return val[0]
}

// RealtimeStreamTargetBitrate 获取实时流目标码率
func (w M9202AudioVideoChannel) RealtimeStreamTargetBitrate() uint32 {
	val, ok := w[0x05]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint32(val)
}

// StorageStreamEncodingMode 获取存储流编码模式
func (w M9202AudioVideoChannel) StorageStreamEncodingMode() common.StorageStreamEncodingMode {
	val, ok := w[0x09]
	if !ok {
		return 0
	}
	return common.StorageStreamEncodingMode(val[0])
}

// StorageStreamResolution 获取存储流分辨率
func (w M9202AudioVideoChannel) StorageStreamResolution() common.StorageStreamResolution {
	val, ok := w[0x0A]
	if !ok {
		return 0
	}
	return common.StorageStreamResolution(val[0])
}

// StorageStreamKeyFrameInterval 获取存储流关键帧间隔
func (w M9202AudioVideoChannel) StorageStreamKeyFrameInterval() uint16 {
	val, ok := w[0x0B]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint16(val)
}

// StorageStreamTargetFrameRate 获取存储流目标帧率
func (w M9202AudioVideoChannel) StorageStreamTargetFrameRate() uint8 {
	val, ok := w[0x0D]
	if !ok {
		return 0
	}
	return val[0]
}

// StorageStreamTargetBitrate 获取存储流目标码率
func (w M9202AudioVideoChannel) StorageStreamTargetBitrate() uint32 {
	val, ok := w[0x0E]
	if !ok {
		return 0
	}
	return binary.BigEndian.Uint32(val)
}

// OSDOverlaySettings 获取OSD字幕叠加设置
func (w M9202AudioVideoChannel) OSDOverlaySettings() common.OSDOverlaySettingsBitmask {
	val, ok := w[0x12]
	if !ok {
		return 0
	}
	return common.OSDOverlaySettingsBitmask(binary.BigEndian.Uint16(val))
}

// EnableAudioOutput 获取是否启用音频输出
func (w M9202AudioVideoChannel) EnableAudioOutput() uint8 {
	val, ok := w[0x14]
	if !ok {
		return 0
	}
	return val[0]
}
