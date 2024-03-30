package coin

import (
	response "coins/internal/delivery/http/responses/coin"
	"coins/internal/domain/coin"
	"coins/pkg/types/query"
	"strconv"
)

func toResponse(coin *coin.Coin) *response.Response {
	return &response.Response{
		ID: strconv.Itoa(int(coin.ID)),
		Short: response.Short{
			Name: coin.Name.String(),
			Code: coin.Code.String(),
		},
	}
}

func toListResponse(count uint64, params query.Query, coins []*coin.Coin) *response.List {
	list := &response.List{
		Total: count,
		Limit: params.Limit,
		Page:  params.Page,
		Data:  []*response.Response{},
	}

	for _, value := range coins {
		list.Data = append(list.Data, toResponse(value))
	}

	return list
}
