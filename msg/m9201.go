package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
	"github.com/mingkid/g-jtt808/msg/common"
	"time"
)

// M9201 数据结构，表示JTT1078协议中的9201数据
type M9201 struct {
	serverIPAddressLength uint8  // 服务器IP 地址长度
	serverIPAddress       string // 服务器IP 地址
	tcpListenPort         uint16 // 服务器视频通道监听端口号(TCP)
	udpListenPort         uint16 // 服务器视频通道监听端口号(UDP)
	logicalChannel        uint8  // 逻辑通道号
	avType                uint8  // 音视频类型
	streamType            uint8  // 码流类型
	storageType           uint8  // 存储器类型
	replayMode            uint8  // 回放方式
	replaySpeed           uint8  // 回放倍速
	startTime             string // 开始时间，BCD格式，YY-MM-DD-HH-MM-SS
	endTime               string // 结束时间，BCD格式，YY-MM-DD-HH-MM-SS
}

// ServerIPAddressLength 获取服务器IP地址长度
func (m M9201) ServerIPAddressLength() byte {
	return m.serverIPAddressLength
}

// ServerIPAddress 获取服务器IP地址
func (m M9201) ServerIPAddress() string {
	return m.serverIPAddress
}

// TCListenPort 获取服务器视频通道监听端口号(TCP)
func (m M9201) TCListenPort() uint16 {
	return m.tcpListenPort
}

// UDListenPort 获取服务器视频通道监听端口号(UDP)
func (m M9201) UDListenPort() uint16 {
	return m.udpListenPort
}

// LogicalChannel 获取逻辑通道号
func (m M9201) LogicalChannel() byte {
	return m.logicalChannel
}

// AVType 获取音视频类型
func (m M9201) AVType() common.AVType {
	return common.AVType(m.avType)
}

// StreamType 获取码流类型
func (m M9201) StreamType() common.StreamType {
	return common.StreamType(m.streamType)
}

// StorageType 获取存储器类型
func (m M9201) StorageType() common.StorageType {
	return common.StorageType(m.storageType)
}

// ReplayMode 获取回放方式
func (m M9201) ReplayMode() M9201ReplayMode {
	return M9201ReplayMode(m.replayMode)
}

// ReplaySpeed 获取回放速度
func (m M9201) ReplaySpeed() M9201ReplaySpeed {
	return M9201ReplaySpeed(m.replaySpeed)
}

// StartTime 获取开始时间
func (m M9201) StartTime() time.Time {
	layout := "06-01-02-15-04-05" // 对应的时间布局
	endTime, err := time.Parse(layout, m.endTime)
	if err != nil {
		return time.Time{}
	}
	return endTime
}

// EndTime 获取结束时间
func (m M9201) EndTime() time.Time {
	layout := "06-01-02-15-04-05" // 对应的时间布局
	endTime, err := time.Parse(layout, m.endTime)
	if err != nil {
		return time.Time{}
	}
	return endTime
}

// SetServerIPAddressLength 设置服务器IP地址长度
func (m *M9201) SetServerIPAddressLength(length uint8) *M9201 {
	m.serverIPAddressLength = length
	return m
}

// SetServerIPAddress 设置服务器IP地址
func (m *M9201) SetServerIPAddress(ipAddress string) *M9201 {
	m.serverIPAddress = ipAddress
	return m
}

// SetTCListenPort 设置服务器视频通道监听端口号(TCP)
func (m *M9201) SetTCListenPort(port uint16) *M9201 {
	m.tcpListenPort = port
	return m
}

// SetUDListenPort 设置服务器视频通道监听端口号(UDP)
func (m *M9201) SetUDListenPort(port uint16) *M9201 {
	m.udpListenPort = port
	return m
}

// SetLogicalChannel 设置逻辑通道号
func (m *M9201) SetLogicalChannel(channel uint8) *M9201 {
	m.logicalChannel = channel
	return m
}

// SetAVType 设置音视频类型
func (m *M9201) SetAVType(avType uint8) *M9201 {
	m.avType = avType
	return m
}

// SetStreamType 设置码流类型
func (m *M9201) SetStreamType(streamType uint8) *M9201 {
	m.streamType = streamType
	return m
}

// SetStorageType 设置存储器类型
func (m *M9201) SetStorageType(storageType uint8) *M9201 {
	m.storageType = storageType
	return m
}

// SetReplayMode 设置回放方式
func (m *M9201) SetReplayMode(mode uint8) *M9201 {
	m.replayMode = mode
	return m
}

// SetReplaySpeed 设置回放速度
func (m *M9201) SetReplaySpeed(speed uint8) *M9201 {
	m.replaySpeed = speed
	return m
}

// SetStartTime 设置开始时间
func (m *M9201) SetStartTime(startTime time.Time) *M9201 {
	m.startTime = startTime.Format("06-01-02-15-04-05")
	return m
}

// SetEndTime 设置结束时间
func (m *M9201) SetEndTime(endTime time.Time) *M9201 {
	m.endTime = endTime.Format("06-01-02-15-04-05")
	return m
}

// Encode 将 M9201 结构体编码为字节流
func (m M9201) Encode() ([]byte, error) {
	w := binary.NewWriter()

	if err := w.WriteUint8(m.serverIPAddressLength); err != nil {
		return nil, fmt.Errorf("编码服务器IP地址长度失败: %v", err)
	}

	if err := w.WriteString(m.serverIPAddress); err != nil {
		return nil, fmt.Errorf("编码服务器IP地址失败: %v", err)
	}

	if err := w.WriteUint16(m.tcpListenPort); err != nil {
		return nil, fmt.Errorf("编码TCP监听端口号失败: %v", err)
	}

	if err := w.WriteUint16(m.udpListenPort); err != nil {
		return nil, fmt.Errorf("编码UDP监听端口号失败: %v", err)
	}

	if err := w.WriteUint8(m.logicalChannel); err != nil {
		return nil, fmt.Errorf("编码逻辑通道号失败: %v", err)
	}

	if err := w.WriteUint8(m.avType); err != nil {
		return nil, fmt.Errorf("编码音视频类型失败: %v", err)
	}

	if err := w.WriteUint8(m.streamType); err != nil {
		return nil, fmt.Errorf("编码码流类型失败: %v", err)
	}

	if err := w.WriteUint8(m.storageType); err != nil {
		return nil, fmt.Errorf("编码存储器类型失败: %v", err)
	}

	if err := w.WriteUint8(m.replayMode); err != nil {
		return nil, fmt.Errorf("编码回放方式失败: %v", err)
	}

	if err := w.WriteUint8(m.replaySpeed); err != nil {
		return nil, fmt.Errorf("编码回放速度失败: %v", err)
	}

	if err := w.WriteBCD(m.startTime, 6); err != nil {
		return nil, fmt.Errorf("编码开始时间失败: %v", err)
	}

	if err := w.WriteBCD(m.endTime, 6); err != nil {
		return nil, fmt.Errorf("编码结束时间失败: %v", err)
	}

	return w.Bytes(), nil
}
