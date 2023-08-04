package url

import (
	"coins/internal/domain/url"
	"coins/pkg/types/query"
	"strconv"
)

func ToResponse(url *url.Url) *Response {
	return &Response{
		ID: strconv.Itoa(int(url.ID)),
		Short: Short{
			Link: url.Link.String(),
			Type: url.SocialMedia.String(),
		},
	}
}

func ToListResponse(count uint64, params query.Query, urls []*url.Url) *List {
	list := &List{
		Total: count,
		Limit: params.Limit,
		Page:  params.Page,
		Data:  []*Response{},
	}

	for _, value := range urls {
		list.Data = append(list.Data, ToResponse(value))
	}

	return list
}
