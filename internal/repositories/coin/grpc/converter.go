package grpc

import "github.com/Nikkoz/coin.sync/pkg/protobuf/coins"

func ToCoinRequest(ids []uint64) *coins.GetCoinsRequest {
	request := make(map[uint32]uint64)

	for k, id := range ids {
		request[uint32(k)] = id
	}

	return &coins.GetCoinsRequest{
		Ids: request,
	}
}
