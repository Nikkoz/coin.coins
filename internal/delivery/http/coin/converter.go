package coin

import (
	"coins/internal/domain/coin"
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
