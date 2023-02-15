package icon

import "github.com/pkg/errors"

const MaxLength = 100

var ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)

type Icon string

func New(icon string) (*Icon, error) {
	if len([]rune(icon)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Icon(icon)
	return &n, nil
}

func (i Icon) String() string {
	return string(i)
}
