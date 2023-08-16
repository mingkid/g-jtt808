package msg

import (
	"encoding/hex"
	"fmt"
	"github.com/mingkid/g-jtt808/msg/common"
	"testing"
)

func TestM9205_Encode(t *testing.T) {
	var m M9205
	m.SetLogicalChannel(12)
	m.SetStartTime("230817030201")
	m.SetEndTime("230820010203")
	m.SetAlarmFlags(0x0010000000000000)
	m.SetAVResourceType(uint8(common.AVTypeVideoOrAudioVideo))
	m.SetStreamType(uint8(common.StreamTypeMain))
	m.SetStorageType(uint8(common.StorageTypeBackup))
	b, err := m.Encode()
	if err != nil {
		fmt.Println("解码失败:", err)
	}

	if r := hex.EncodeToString(b); r != "0c2308170302012308200102030010000000000000030102" {
		t.Fatalf("消息包组包错误，应为%s，实际为%s", "0c2308170302012308200102030000000000000000030102", r)
	}
}
