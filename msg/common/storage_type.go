package common

// StorageType 存储位置
type StorageType byte

const (
	// StorageTypeMainOrBackup 存储位置：主存储器或灾备存储器
	StorageTypeMainOrBackup StorageType = iota

	// StorageTypeMain 存储位置：主存储器
	StorageTypeMain

	// StorageTypeBackup 存储位置：灾备存储器
	StorageTypeBackup
)
