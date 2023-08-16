package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

type Writer struct {
	b *bytes.Buffer
}

// WriteByte 写入byte数据
func (w Writer) WriteByte(data byte) error {
	return w.b.WriteByte(data)
}

// WriteUint8 写入8位数据
func (w Writer) WriteUint8(data uint8) error {
	_, err := w.b.Write([]byte{data})
	return err
}

// WriteUint16 写入16位数据
func (w Writer) WriteUint16(data uint16) (err error) {
	var temp = make([]byte, 2)
	binary.BigEndian.PutUint16(temp[:], data)
	_, err = w.b.Write(temp)
	return
}

// WriteUint64 写入64位数据
func (w Writer) WriteUint64(data uint64) (err error) {
	var temp = make([]byte, 8)
	binary.BigEndian.PutUint64(temp[:], data)
	_, err = w.b.Write(temp)
	return
}

// WriteBCD 写入BCD编码
func (w *Writer) WriteBCD(s string, length int) error {
	// 长度不够补位
	if len(s)%4 != 0 {
		for i := 0; i < len(s)%4; i++ {
			s = "0" + s
		}
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	var data []byte
	if length != 0 {
		for i := len(b); i < length; i++ {
			data = append(data, 0x00)
		}
		data = append(data, b...)
	}

	_, err = w.b.Write(data)
	return err
}

func (w Writer) WriteString(s string) error {
	b, err := io.ReadAll(transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder()))
	if err != nil {
		return err
	}

	_, err = w.b.Write(b)
	return err
}

func (w Writer) Write(b []byte) error {
	_, err := w.b.Write(b)
	return err
}

func (w Writer) Bytes() []byte {
	return w.b.Bytes()
}

func NewWriter() *Writer {
	return &Writer{
		b: &bytes.Buffer{},
	}
}

// Uint32ToBytesMap uint32转换为字节映射
func Uint32ToBytesMap(value uint32) map[byte][]byte {
	bytes_ := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes_, value)

	bytesMap := make(map[byte][]byte)
	for i, b := range bytes_ {
		bytesMap[byte(i)] = []byte{b}
	}

	return bytesMap
}

// Uint8ToBytesMap uint8转换为字节映射
func Uint8ToBytesMap(value uint8) map[byte][]byte {
	bytes_ := []byte{value}

	bytesMap := make(map[byte][]byte)
	for i, b := range bytes_ {
		bytesMap[byte(i)] = []byte{b}
	}

	return bytesMap
}
