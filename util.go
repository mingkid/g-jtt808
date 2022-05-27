package jtt808

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Tag(f reflect.StructField) (t string, length int, err error) {
	tag, ok := f.Tag.Lookup("jtt808")
	if !ok {
		err = fmt.Errorf("未定义数据长度")
	}

	tagItems := strings.Split(tag, ",")
	t = tagItems[1]
	length, err = strconv.Atoi(tagItems[0])

	if err != nil {
		err = fmt.Errorf("数据长度不正确，请按照：'长度,类型'格式来设置Tag")
	}

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
