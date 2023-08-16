package msg

import (
	"testing"
)

func TestM9101_Decode(t *testing.T) {

    var m M9101
    _ = m.Decode([]byte{13, 49, 57, 50, 46, 49, 54, 56, 46, 48, 46, 49, 50, 51, 0, 123, 1, 200, 3, 0, 0})

    if m.ServerIPAddressLength() != 13 {
        t.Fatalf("期望的服务器IP地址长度为 %d，但得到 %d", 126, m.ServerIPAddressLength())
    }

    if m.ServerIPAddress() != "192.168.0.123" {
        t.Fatalf("期望的服务器IP地址为 %s，但得到 %s", "192.168.0.123", m.ServerIPAddress())
    }

    if m.TCListenPort() != 123 {
        t.Fatalf("期望的TCP监听端口号为 %d，但得到 %d", 123, m.TCListenPort())
    }

    if m.UDListenPort() != 456 {
        t.Fatalf("期望的UDP监听端口号为 %d，但得到 %d", 456, m.UDListenPort())
    }

    if m.LogicalChannel() != 3 {
        t.Fatalf("期望的逻辑通道号为 %d，但得到 %d", 3, m.LogicalChannel())
    }

    if m.DataType() != DataTypeAudioVideo {
        t.Fatalf("期望的数据类型为 %v，但得到 %v", DataTypeAudioVideo, m.DataType())
    }

    if m.StreamType() != StreamTypeMain {
        t.Fatalf("期望的码流类型为 %v，但得到 %v", StreamTypeMain, m.StreamType())
    }
}
