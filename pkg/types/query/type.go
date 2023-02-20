package query

import (
	"coins/pkg/types/sort"
	"github.com/gin-gonic/gin"
)

type (
	Query struct {
		Sorts  sort.Sorts
		Page   uint64
		Limit  uint64
		Offset uint64
	}

	SortOptions struct {
	}

	Options struct {
		// Тут можно добавить фильтр
		Sorts SortsOptions
	}

	SortsOptions map[string]SortOptions // map[front_key]FilterOptions
)

var (
	sortKey  = "sort"
	limitKey = "limit"
	pageKey  = "page"
)

func Parse(c *gin.Context, options Options) (*Query, error) {
	data, ok := c.GetQueryMap(sortKey)
	if !ok {
		data = map[string]string{}
	}

	sorts, err := parseSorts(data, options.Sorts)
	if err != nil {
		return nil, err
	}

	page, limit, offset := parsePagination(c.Query(limitKey), c.Query(pageKey))

	return &Query{
		Sorts:  sorts,
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}, nil
}
