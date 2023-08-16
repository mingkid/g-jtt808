package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
	"github.com/mingkid/g-jtt808/msg/common"
	"time"
)

// M9205 数据结构，表示JTT1078协议中的9205数据
type M9205 struct {
	logicalChannel uint8  // 逻辑通道号
	startTime      string // 开始时间，BCD格式，YY-MM-DD-HH-MM-SS
	endTime        string // 结束时间，BCD格式，YY-MM-DD-HH-MM-SS
	alarmFlags     uint64 // 报警标志
	avResourceType uint8  // 音视频资源类型
	streamType     uint8  // 码流类型
	storageType    uint8  // 存储器类型
}

// LogicalChannel 获取逻辑通道号
func (m M9205) LogicalChannel() byte {
	return m.logicalChannel
}

// StartTime 获取开始时间
func (m M9205) StartTime() time.Time {
	layout := "06-01-02-15-04-05" // 对应的时间布局
	startTime, err := time.Parse(layout, m.startTime)
	if err != nil {
		return time.Time{}
	}
	return startTime
}

// EndTime 获取结束时间
func (m M9205) EndTime() time.Time {
	layout := "06-01-02-15-04-05" // 对应的时间布局
	endTime, err := time.Parse(layout, m.endTime)
	if err != nil {
		return time.Time{}
	}
	return endTime
}

// AlarmFlags 获取报警标志
func (m M9205) AlarmFlags() M9205AlarmFlags {
	return M9205AlarmFlags{
		M0200Warn: M0200Warn(uint32(m.alarmFlags)),
		M9205Warn: binary.Uint32ToBytesMap(uint32(m.alarmFlags >> 32)),
	}
}

// AVResourceType 获取音视频资源类型
func (m M9205) AVResourceType() common.AVType {
	return common.AVType(m.avResourceType)
}

// StreamType 获取码流类型
func (m M9205) StreamType() common.StreamType {
	return common.StreamType(m.streamType)
}

// StorageType 获取存储器类型
func (m M9205) StorageType() common.StorageType {
	return common.StorageType(m.storageType)
}

// SetLogicalChannel 设置逻辑通道号
func (m *M9205) SetLogicalChannel(channel uint8) *M9205 {
	m.logicalChannel = channel
	return m
}

// SetStartTime 设置开始时间
func (m *M9205) SetStartTime(startTime string) *M9205 {
	m.startTime = startTime
	return m
}

// SetEndTime 设置结束时间
func (m *M9205) SetEndTime(endTime string) *M9205 {
	m.endTime = endTime
	return m
}

// SetAlarmFlags 设置报警标志
func (m *M9205) SetAlarmFlags(alarmFlags uint64) *M9205 {
	m.alarmFlags = alarmFlags
	return m
}

// SetAVResourceType 设置音视频资源类型
func (m *M9205) SetAVResourceType(resourceType uint8) *M9205 {
	m.avResourceType = resourceType
	return m
}

// SetStreamType 设置码流类型
func (m *M9205) SetStreamType(streamType uint8) *M9205 {
	m.streamType = streamType
	return m
}

// SetStorageType 设置存储器类型
func (m *M9205) SetStorageType(storageType uint8) *M9205 {
	m.storageType = storageType
	return m
}

// Encode 将 M9205 结构体编码为字节流
func (m M9205) Encode() ([]byte, error) {
	w := binary.NewWriter()

	if err := w.WriteUint8(m.logicalChannel); err != nil {
		return nil, fmt.Errorf("编码逻辑通道号失败: %v", err)
	}

	if err := w.WriteBCD(m.startTime, 6); err != nil {
		return nil, fmt.Errorf("编码开始时间失败: %v", err)
	}

	if err := w.WriteBCD(m.endTime, 6); err != nil {
		return nil, fmt.Errorf("编码结束时间失败: %v", err)
	}

	if err := w.WriteUint64(m.alarmFlags); err != nil {
		return nil, fmt.Errorf("编码报警标志失败: %v", err)
	}

	if err := w.WriteUint8(m.avResourceType); err != nil {
		return nil, fmt.Errorf("编码音视频资源类型失败: %v", err)
	}

	if err := w.WriteUint8(m.streamType); err != nil {
		return nil, fmt.Errorf("编码码流类型失败: %v", err)
	}

	if err := w.WriteUint8(m.storageType); err != nil {
		return nil, fmt.Errorf("编码存储器类型失败: %v", err)
	}

	return w.Bytes(), nil
}
