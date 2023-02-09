package models

import "gorm.io/gorm"

type Coin struct {
	ID uint `gorm:"primaryKey;AUTO_INCREMENT"`

	Name     string     `gorm:"size:100;not null;comment:Название"`
	Code     string     `gorm:"size:50;not null;comment:Код монеты, например BTC"`
	Icon     string     `gorm:"size:100;comment:Ссылка на иконку монеты"`
	CoinUrls []*CoinUrl `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.DeletedAt `gorm:"index"`
}
