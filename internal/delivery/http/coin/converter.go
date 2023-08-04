package coin

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/query"
	"strconv"
)

func ToResponse(coin *coin.Coin) *Response {
	return &Response{
		ID: strconv.Itoa(int(coin.ID)),
		Short: Short{
			Name: coin.Name.String(),
			Code: coin.Code.String(),
		},
	}
}

func ToListResponse(count uint64, params query.Query, coins []*coin.Coin) *List {
	list := &List{
		Total: count,
		Limit: params.Limit,
		Page:  params.Page,
		Data:  []*Response{},
	}

	for _, value := range coins {
		list.Data = append(list.Data, ToResponse(value))
	}

	return list
}
