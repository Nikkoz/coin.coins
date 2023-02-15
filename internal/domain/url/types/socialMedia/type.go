package socialMedia

import (
	"github.com/pkg/errors"
)

const MaxLength = 10

var ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)

type SocialMedia string

func New(sm string) (*SocialMedia, error) {
	if len([]byte(sm)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := SocialMedia(sm)
	return &n, nil
}

func (sm SocialMedia) String() string {
	return string(sm)
}
