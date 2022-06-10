package jtt808

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Tag 获取 message.Message 中的指定字段 reflect.StructField 对应的JTT808类型和长度
func Tag(f reflect.StructField, parent reflect.Value) (t string, length int, err error) {
	tag, ok := f.Tag.Lookup("jtt808")
	if !ok {
		err = fmt.Errorf("未定义数据长度")
	}

	tagItems := strings.Split(tag, ",")
	t = tagItems[1]
	if length, err = strconv.Atoi(tagItems[0]); err != nil {
		// 变量长度获取
		lenV := parent.FieldByName(tagItems[0])
		if lenV.IsValid() {
			if length, err = strconv.Atoi(lenV.String()); err != nil {
				err = fmt.Errorf("数据长度不正确，请按照：'长度,类型'格式来设置Tag")
			}
		}
	}

	// 常量长度获取
	return
}

func ConvHEX(data []byte) (s string) {
	s = hex.EncodeToString(data)
	for i := 2; i <= len(data)*2; i += 2 {
		fmt.Printf("%s ", s[i-2:i])
	}
	fmt.Println("")
	return
}
