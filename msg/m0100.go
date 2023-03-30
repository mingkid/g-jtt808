package msg

import (
	"fmt"
	"github.com/mingkid/g-jtt808/binary"
)

type M0100 struct {
	provincialID      uint16
	cityID            uint16
	manufacturerID    []byte
	termModel         []byte
	termID            []byte
	licensePlateColor byte
	licensePlate      []byte
}

func (m M0100) ProvincialID() uint16 {
	return m.provincialID
}

func (m M0100) CityID() uint16 {
	return m.cityID
}

func (m M0100) ManufacturerID() (string, error) {
	r := binary.NewReader(m.manufacturerID)
	res, err := r.ReadString(5)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m M0100) TermModel() (string, error) {
	r := binary.NewReader(m.termModel)
	res, err := r.ReadString(20)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m M0100) TermID() (string, error) {
	r := binary.NewReader(m.termID)
	res, err := r.ReadString(7)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m M0100) LicensePlateColor() byte {
	return m.licensePlateColor
}

func (m M0100) LicensePlate() (string, error) {
	r := binary.NewReader(m.licensePlate)
	res, err := r.ReadAllString()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m *M0100) Decode(b []byte) (err error) {
	r := binary.NewReader(b)

	if m.provincialID, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("省域ID解码失败")
	}
	if m.cityID, err = r.ReadUint16(); err != nil {
		return fmt.Errorf("市县域ID解码失败")
	}
	if m.manufacturerID, err = r.ReadBytes(5); err != nil {
		return fmt.Errorf("制造商ID解码失败")
	}
	if m.termModel, err = r.ReadBytes(20); err != nil {
		return fmt.Errorf("终端型号解码失败")
	}
	if m.termID, err = r.ReadBytes(7); err != nil {
		return fmt.Errorf("终端ID解码失败")
	}
	if m.licensePlateColor, err = r.ReadByte(); err != nil {
		return fmt.Errorf("车牌颜色解码失败")
	}
	if m.licensePlate, err = r.ReadAllBytes(); err != nil {
		return fmt.Errorf("车牌解码失败")
	}
	return
}
