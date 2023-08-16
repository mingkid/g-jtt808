package msg

import "encoding/binary"

// MemoryFailureAlarm 表示存储器故障报警信息
type MemoryFailureAlarm []byte

// MainMemoryStatus 获取主存储器状态
func (w MemoryFailureAlarm) MainMemoryStatus() uint16 {
	return binary.BigEndian.Uint16(w[0x00:0x11])
}

// BackupMemoryStatus 获取备份存储器状态
func (w MemoryFailureAlarm) BackupMemoryStatus() uint16 {
	return binary.BigEndian.Uint16(w[0x12:0x15])
}
