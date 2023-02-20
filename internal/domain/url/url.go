package url

import (
	"coins/internal/domain/url/types/externalId"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
)

type Url struct {
	ID uint `gorm:"primaryKey;AUTO_INCREMENT"`

	ExternalID  *externalId.ExternalID  `gorm:"unique;comment:Внешний идентификатор записи"`
	CoinID      uint                    `gorm:"not null;comment:Идентификатор монеты"`
	Link        link.Link               `gorm:"not null;comment:Ссылка на источник"`
	SocialMedia socialMedia.SocialMedia `gorm:"not null;comment:Тип социальной сети" sql:"type:ENUM(twitter, reddit)"`
}

func (Url) Table() string {
	return "coin_urls"
}

func New(externalId *externalId.ExternalID, link link.Link, social socialMedia.SocialMedia) *Url {
	url := &Url{
		Link:        link,
		SocialMedia: social,
	}

	if externalId != nil {
		url.ExternalID = externalId
	}

	return url
}
