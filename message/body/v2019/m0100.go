package v2019

import "strings"

type M0100 struct {
	RawProvincialID      uint16
	RawCityID            uint16
	RawManufacturerID    []byte `jtt808:"11,slice"`
	RawTermModel         []byte `jtt808:"30,slice"`
	RawTermID            []byte `jtt808:"30,slice"`
	RawLicensePlateColor byte
	RawLicensePlate      string `jtt808:"0,string"`
}

func (m M0100) CityID() uint16 {
	return m.RawCityID
}

func (m *M0100) SetCityID(id uint16) {
	m.RawCityID = id
}

func (m M0100) TermID() string {
	return strings.Trim(string(m.RawTermID), "\u0000")
}

func (m *M0100) SetTermID(id string) *M0100 {
	m.RawTermID = []byte(id)
	return m
}

func (m M0100) TermModel() string {
	return strings.Trim(string(m.RawTermModel), "\u0000")
}

func (m *M0100) SetTermModel(model string) *M0100 {
	m.RawTermModel = []byte(model)
	return m
}

func (m M0100) ManufacturerID() string {
	return strings.Trim(string(m.RawManufacturerID), "\u0000")
}

func (m *M0100) SetManufacturerID(id string) *M0100 {
	m.RawManufacturerID = []byte(id)
	return m
}

func (m M0100) ProvincialID() uint16 {
	return m.RawProvincialID
}

func (m *M0100) SetProvincialID(id uint16) *M0100 {
	m.RawProvincialID = id
	return m
}

func (m M0100) LicensePlateColor() byte {
	return m.RawLicensePlateColor
}

func (m *M0100) SetLicensePlateColor(v byte) *M0100 {
	m.RawLicensePlateColor = v
	return m
}

func (m M0100) LicensePlate() string {
	return m.RawLicensePlate
}

func (m *M0100) SetLicensePlate(v string) *M0100 {
	m.RawLicensePlate = v
	return m
}
