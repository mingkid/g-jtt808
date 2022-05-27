package v2013

import (
	"io"
	"time"

	jt808b "github.com/mingkid/g-jtt808/binary"
	"github.com/mingkid/g-jtt808/message/body/v2013/extra0200"
)

type M0200 struct {
	Warn         M0200Warn
	Status       M0200Status
	RawLatitude  uint32
	RawLongitude uint32
	RawAltitude  uint16
	RawSpeed     uint16
	RawDirection uint16
	RawTime      string `jtt808:"6,bcd"`
	RawExtras    []byte `jtt808:"0,raw"`
}

func (m M0200) Latitude() uint32 {
	return m.RawLatitude
}

func (m M0200) Longitude() uint32 {
	return m.RawLongitude
}

func (m M0200) Altitude() uint16 {
	return m.RawAltitude
}

func (m M0200) Speed() uint16 {
	return m.RawSpeed
}

func (m M0200) Direction() uint16 {
	return m.RawDirection
}

func (m M0200) Time() (t time.Time) {
	t, _ = time.Parse(time.RFC3339, m.RawTime)
	return
}

func (m M0200) Extras() (result extra0200.Extra, err error) {
	result = make(map[byte][]byte, 72)
	if len(m.RawExtras) == 0 {
		return
	}

	er := jt808b.ErrReader{R: jt808b.NewReader(m.RawExtras)}
	for {
		if er.Err != nil && er.Err == io.EOF {
			break
		} else if er.Err != nil {
			return nil, er.Err
		}
		id := er.ReadByte()
		length := er.ReadByte()
		v := er.ReadBytes(int(length))
		result[id] = v
	}
	return
}
