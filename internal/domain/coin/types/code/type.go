package code

import "github.com/pkg/errors"

const MaxLength = 50

var ErrWrongLength = errors.Errorf("code must be less than or equal to %d characters", MaxLength)

type Code string

func New(code string) (*Code, error) {
	if len([]rune(code)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Code(code)
	return &n, nil
}

func (c Code) String() string {
	return string(c)
}
