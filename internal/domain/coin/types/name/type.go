package name

import "github.com/pkg/errors"

const MaxLength = 100

var ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)

type Name string

func New(name string) (*Name, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Name(name)
	return &n, nil
}

func (n Name) String() string {
	return string(n)
}
