package externalId

import "strconv"

type ExternalID uint

func New(id uint) (*ExternalID, error) {
	n := ExternalID(id)
	return &n, nil
}

func (e ExternalID) String() string {
	return strconv.FormatUint(uint64(e), 10)
}
