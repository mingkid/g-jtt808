package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
	"time"
)

// M9202 数据结构，表示JTT1078协议中的9202数据
type M9202 struct {
	audioVideoChannel uint8  // 音视频通道号
	replayControl     uint8  // 回放控制
	replaySpeed       uint8  // 快进或快退倍数
	dragPlayTime      string // 拖动回放位置，BCD格式，YY-MM-DD-HH-MM-SS
}

// AudioVideoChannel 获取音视频通道号
func (m M9202) AudioVideoChannel() M9202AudioVideoChannel {
	return binary.Uint8ToBytesMap(m.audioVideoChannel)
}

// ReplayControl 获取回放控制类型
func (m M9202) ReplayControl() ReplayControl {
	return ReplayControl(m.replayControl)
}

// ReplaySpeed 获取快进或快退倍数类型
func (m M9202) ReplaySpeed() ReplaySpeed {
	return ReplaySpeed(m.replaySpeed)
}

// DragPlayPosition 获取拖动回放位置
func (m M9202) DragPlayPosition() time.Time {
	layout := "06-01-02-15-04-05" // 对应的时间布局
	dragPlayTime, err := time.Parse(layout, m.dragPlayTime)
	if err != nil {
		return time.Time{}
	}

	return dragPlayTime
}

// SetAudioVideoChannel 设置音视频通道号
func (m *M9202) SetAudioVideoChannel(channel uint8) *M9202 {
	m.audioVideoChannel = channel
	return m
}

// SetReplayControl 设置回放控制类型
func (m *M9202) SetReplayControl(control ReplayControl) *M9202 {
	m.replayControl = uint8(control)
	return m

}

// SetReplaySpeed 设置快进或快退倍数类型
func (m *M9202) SetReplaySpeed(speed ReplaySpeed) *M9202 {
	m.replaySpeed = uint8(speed)
	return m

}

// SetDragPlayTime 设置拖动回放位置
func (m *M9202) SetDragPlayTime(time string) *M9202 {
	m.dragPlayTime = time
	return m

}

// Encode 将 M9202 结构体编码为字节流
func (m M9202) Encode() ([]byte, error) {
	w := binary.NewWriter()

	if err := w.WriteUint8(m.audioVideoChannel); err != nil {
		return nil, fmt.Errorf("编码音视频通道号失败: %v", err)
	}

	if err := w.WriteUint8(m.replayControl); err != nil {
		return nil, fmt.Errorf("编码回放控制失败: %v", err)
	}

	if err := w.WriteUint8(m.replaySpeed); err != nil {
		return nil, fmt.Errorf("编码快进或快退倍数失败: %v", err)
	}

	// 拖动回放时 提供拖动回放时间
	if m.replayControl == uint8(ReplayControlDrag) {
		if err := w.WriteBCD(m.dragPlayTime, 6); err != nil {
			return nil, fmt.Errorf("编码拖动回放位置失败: %v", err)
		}
	}

	return w.Bytes(), nil
}
