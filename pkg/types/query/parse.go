package query

import (
	"coins/pkg/types/columnCode"
	"coins/pkg/types/logger"
	"coins/pkg/types/sort"
	"strconv"
)

const (
	DefaultValueForLimit uint64 = 10
	MaxValueForLimit     uint64 = 100
)

func parseSorts(sorts map[string]string, options SortsOptions) (sort.Sorts, error) {
	var result = make(sort.Sorts, 0)
	if len(sorts) == 0 {
		return result, nil
	}

	for field, d := range sorts {
		key, err := columnCode.New(field)
		if err != nil {
			logger.Error(err)

			return nil, err
		}

		if _, ok := options[key.String()]; !ok {
			continue
		}

		direction, ok := sort.DirectionFromQuery(d)
		if !ok {
			continue
		}

		result = append(result, &sort.Sort{
			Key:       key,
			Direction: direction,
		})
	}

	return result, nil
}

func parsePagination(strLimit, strPage string) (uint64, uint64) {
	page, err := strconv.ParseUint(strPage, 10, 64)
	if err != nil {
		page = 1
	}

	limit := parseLimit(strLimit)

	return page, limit
}

func parseLimit(value string) uint64 {
	limit, err := strconv.ParseUint(value, 10, 64)
	if err != nil || limit == 0 {
		return DefaultValueForLimit
	}

	if limit > MaxValueForLimit {
		return MaxValueForLimit
	}

	return limit
}
