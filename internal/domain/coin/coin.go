package coin

import (
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/icon"
	"coins/internal/domain/coin/types/name"
	"coins/internal/domain/url"
	"gorm.io/gorm"
)

type Coin struct {
	ID uint `gorm:"primaryKey;AUTO_INCREMENT"`

	Name     name.Name  `gorm:"size:100;not null;comment:Название"`
	Code     code.Code  `gorm:"size:50;not null;comment:Код монеты, например BTC"`
	Icon     icon.Icon  `gorm:"size:100;comment:Ссылка на иконку монеты"`
	CoinUrls []*url.Url `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func New(name name.Name, code code.Code, icon *icon.Icon) *Coin {
	coin := &Coin{
		Name: name,
		Code: code,
	}

	if icon != nil {
		coin.Icon = *icon
	}

	return coin
}
