package link

import (
	"github.com/pkg/errors"
)

const MaxLength = 10

var ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)

type Link string

func New(link string) (*Link, error) {
	if len([]byte(link)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Link(link)
	return &n, nil
}

func (l Link) String() string {
	return string(l)
}
