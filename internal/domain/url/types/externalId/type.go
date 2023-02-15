package externalId

import (
	"github.com/pkg/errors"
	"strconv"
)

const MaxLength uint = 10

var ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)

type ExternalID uint

func New(id uint) (*ExternalID, error) {
	if id > MaxLength {
		return nil, ErrWrongLength
	}

	n := ExternalID(id)
	return &n, nil
}

func (e ExternalID) String() string {
	return strconv.FormatUint(uint64(e), 10)
}
