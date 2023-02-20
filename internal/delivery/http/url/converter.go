package url

import (
	"coins/internal/domain/url"
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
