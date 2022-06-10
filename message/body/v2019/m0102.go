package v2019

import "strings"

type M0102 struct {
	RawTokenLength byte
	RawToken       string `jtt808:"RawTokenLength,string"`
	RawIMEI        []byte `jtt808:"15,slice"`
	RawSoftVer     []byte `jtt808:"20,slice"`
}

func (m M0102) TokenLength() byte {
	return m.RawTokenLength
}

func (m M0102) Token() string {
	return strings.Trim(m.RawToken, "\u0000")
}

func (m *M0102) SetToken(t string) *M0102 {
	m.RawTokenLength = byte(len([]byte(t)))
	m.RawToken = t
	return m
}

func (m M0102) IMEI() string {
	return strings.Trim(string(m.RawIMEI), "\u0000")
}

func (m *M0102) SetIMEI(imei string) *M0102 {
	m.RawIMEI = []byte(imei)
	return m
}

func (m M0102) SoftVer() string {
	return strings.Trim(string(m.RawSoftVer), "\u0000")
}

func (m *M0102) SetSoftVer(v string) *M0102 {
	m.RawSoftVer = []byte(v)
	return m
}
