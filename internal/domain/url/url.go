package url

import (
	"coins/internal/domain/url/types/externalId"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
)

type Url struct {
	ID uint `gorm:"primaryKey;AUTO_INCREMENT"`

	ExternalID  externalId.ExternalID   `gorm:"unique;not null;comment:Внешний идентификатор записи"`
	CoinID      uint                    `gorm:"not null;comment:Идентификатор монеты"`
	Link        link.Link               `gorm:"not null;comment:Ссылка на источник"`
	SocialMedia socialMedia.SocialMedia `gorm:"not null;comment:Тип социальной сети" sql:"type:ENUM(twitter, reddit)"`
}

func (Url) Table() string {
	return "coin_urls"
}
