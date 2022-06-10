package jtt808

import (
	"reflect"
	"strings"

	"github.com/mingkid/g-jtt808/binary"
)

type Decoder struct {
	msg any
}

func (d Decoder) Decode(b []byte) (err error) {
	// 去除首位标识位，校验码
	b = binary.EliminateChecksum(binary.EliminateIdentityBit(b))
	er := binary.ErrReader{R: binary.NewReader(b)}

	v := reflect.ValueOf(d.msg).Elem()
	hv := v.FieldByName("Head")
	bv := v.FieldByName("Body")

	d.decode(hv, er)
	if bv.Kind() == reflect.Ptr && bv.IsNil() {
		return er.Err
	}
	d.decode(bv, er)

	return er.Err
}

func (d Decoder) decode(v reflect.Value, r binary.ErrReader) {
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)

		switch fv.Kind() {
		case reflect.Uint8:
			fv.SetUint(uint64(r.ReadByte()))
		case reflect.Uint16:
			fv.SetUint(uint64(r.ReadUint16()))
		case reflect.Uint32:
			fv.SetUint(uint64(r.ReadUint32()))
		case reflect.String:
			d.decodeString(r, fv, v.Type().Field(i), v)
		case reflect.Slice:
			d.decodeSlice(r, fv, v.Type().Field(i), v)
		case reflect.Struct:
			d.decode(fv, r)
		case reflect.Ptr:
			if fv.IsNil() {
				break
			}
			d.decode(fv.Elem(), r)
		}
	}
}

func (d Decoder) decodeString(r binary.ErrReader, v reflect.Value, f reflect.StructField, parent reflect.Value) {
	var (
		t      string
		length int
	)
	t, length, r.Err = Tag(f, parent)

	switch t {
	case "bcd":
		v.Set(reflect.ValueOf(r.ReadBCD(length)))
	case "string":
		v.Set(reflect.ValueOf(strings.Trim(r.ReadString(r.R.Len()), "\u0000")))
	}
}

func (d Decoder) decodeSlice(r binary.ErrReader, v reflect.Value, f reflect.StructField, parent reflect.Value) {
	var length int
	if _, length, r.Err = Tag(f, parent); length == 0 {
		v.Set(reflect.ValueOf(r.ReadBytes(r.R.Len())))
	} else {
		v.Set(reflect.ValueOf(r.ReadBytes(length)))
	}
}

func NewDecoder(msg any) Decoder {
	return Decoder{msg: msg}
}
