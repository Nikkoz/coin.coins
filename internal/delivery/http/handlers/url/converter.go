package url

import (
	response "coins/internal/delivery/http/responses/url"
	"coins/internal/domain/url"
	"coins/pkg/types/query"
	"strconv"
)

func ToResponse(url *url.Url) *response.Response {
	return &response.Response{
		ID: strconv.Itoa(int(url.ID)),
		Short: response.Short{
			Link: url.Link.String(),
			Type: url.SocialMedia.String(),
		},
	}
}

func ToListResponse(count uint64, params query.Query, urls []*url.Url) *response.List {
	list := &response.List{
		Total: count,
		Limit: params.Limit,
		Page:  params.Page,
		Data:  []*response.Response{},
	}

	for _, value := range urls {
		list.Data = append(list.Data, ToResponse(value))
	}

	return list
}
