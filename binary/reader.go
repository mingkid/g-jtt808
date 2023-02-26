package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

type Reader struct {
	r *bytes.Reader
}

// ReadBCD 读取BCD编码
func (reader *Reader) ReadBCD(length uint) (string, error) {
	res := make([]byte, length)
	n, err := reader.r.Read(res)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", res[:n]), nil
}

// ReadByte 读取单字节
func (reader *Reader) ReadByte() (res uint8, err error) {
	res, err = reader.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return
}

// ReadUint16 读取双字节
func (reader *Reader) ReadUint16() (uint16, error) {
	tmp := make([]byte, 2)
	_, err := reader.r.Read(tmp)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(tmp), nil
}

// ReadUint32 读取四字节
func (reader *Reader) ReadUint32() (uint32, error) {
	tmp := make([]byte, 4)
	_, err := reader.r.Read(tmp)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(tmp), nil
}

// ReadBytes 读取多个字节
func (reader *Reader) ReadBytes(size int) ([]byte, error) {
	res := make([]byte, size)

	_, err := reader.r.Read(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (reader *Reader) ReadAllBytes() ([]byte, error) {
	return reader.ReadBytes(reader.r.Len())
}

func (reader *Reader) ReadString(size int) (string, error) {
	sourceBuf, err := reader.ReadBytes(size)
	if err != nil {
		return "", err
	}

	targetBuf, err := io.ReadAll(transform.NewReader(bytes.NewReader(sourceBuf), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return "", err
	}

	return strings.Trim(string(targetBuf), "\u0000"), nil
}

func (reader *Reader) ReadAllString() (string, error) {
	targetBuf, err := io.ReadAll(transform.NewReader(reader.r, simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return "", err
	}

	return strings.Trim(string(targetBuf), "\u0000"), nil
}

func NewReader(b []byte) *Reader {
	return &Reader{
		r: bytes.NewReader(b),
	}
}
