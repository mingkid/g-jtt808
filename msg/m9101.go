package msg

import (
    "fmt"
    "github.com/mingkid/g-jtt808/binary"
)

// M9101 JT/T1078 9101数据
type M9101 struct {
    serverIPAddressLength uint8  // 服务器IP 地址长度
    serverIPAddress       string // 服务器IP 地址
    tcpListenPort         uint16 // 服务器视频通道监听端口号(TCP)
    udpListenPort         uint16 // 服务器视频通道监听端口号(UDP)
    logicalChannel        uint8  // 逻辑通道号
    dataType              uint8  // 数据类型
    streamType            uint8  // 码流类型
}

// ServerIPAddressLength 获取服务器IP地址长度
func (m M9101) ServerIPAddressLength() uint8 {
    return m.serverIPAddressLength
}

// ServerIPAddress 获取服务器IP地址
func (m M9101) ServerIPAddress() string {
    return m.serverIPAddress
}

// TCListenPort 获取服务器视频通道监听端口号(TCP)
func (m M9101) TCListenPort() uint16 {
    return m.tcpListenPort
}

// UDListenPort 获取服务器视频通道监听端口号(UDP)
func (m M9101) UDListenPort() uint16 {
    return m.udpListenPort
}

// LogicalChannel 获取逻辑通道号
func (m M9101) LogicalChannel() uint8 {
    return m.logicalChannel
}

// DataType 获取数据类型
func (m M9101) DataType() M9101DataType {
    return M9101DataType(m.dataType)
}

// StreamType 获取码流类型
func (m M9101) StreamType() M9101StreamType {
    return M9101StreamType(m.streamType)
}

// Decode 解码M9101数据包
func (m *M9101) Decode(b []byte) (err error) {
    r := binary.NewReader(b)

    if m.serverIPAddressLength, err = r.ReadByte(); err != nil {
        return fmt.Errorf("服务器IP地址长度解码失败")
    }

    if m.serverIPAddress, err = r.ReadString(int(m.serverIPAddressLength)); err != nil {
        return fmt.Errorf("服务器IP地址解码失败")
    }
    if m.tcpListenPort, err = r.ReadUint16(); err != nil {
        return fmt.Errorf("TCP监听端口号解码失败")
    }
    if m.udpListenPort, err = r.ReadUint16(); err != nil {
        return fmt.Errorf("UDP监听端口号解码失败")
    }
    if m.logicalChannel, err = r.ReadByte(); err != nil {
        return fmt.Errorf("逻辑通道号解码失败")
    }
    if m.dataType, err = r.ReadByte(); err != nil {
        return fmt.Errorf("数据类型解码失败")
    }
    if m.streamType, err = r.ReadByte(); err != nil {
        return fmt.Errorf("码流类型解码失败")
    }

    return nil
}
