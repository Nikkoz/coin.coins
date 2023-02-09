package models

type CoinUrl struct {
	ID uint `gorm:"primaryKey;AUTO_INCREMENT"`

	ExternalID uint   `gorm:"unique;not null;comment:Внешний идентификатор записи"`
	CoinID     uint   `gorm:"not null;comment:Идентификатор монеты"`
	Link       string `gorm:"not null;comment:Ссылка на источник"`
	Type       string `gorm:"not null;comment:Тип социальной сети" sql:"type:ENUM(twitter, reddit)"`
}
