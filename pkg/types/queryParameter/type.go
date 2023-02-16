package queryParameter

import (
	"coins/pkg/types/pagination"
	"coins/pkg/types/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
