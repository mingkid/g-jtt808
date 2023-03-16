package msg

import (
	"github.com/mingkid/g-jtt808/binary"
)

type M0102 struct {
	token []byte
}

func (m M0102) Token() (string, error) {
	r := binary.NewReader(m.token)
	res, err := r.ReadAllString()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m *M0102) Decode(b []byte) (err error) {
	r := binary.NewReader(b)

	if m.token, err = r.ReadAllBytes(); err != nil {
		return err
	}
	return nil
}
