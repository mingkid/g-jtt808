package common

type EncryptType uint8

const (
	None EncryptType = iota
	// RSA 加密
	RSA
)
