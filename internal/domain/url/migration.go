package url

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	//db.Migrator().AlterColumn(&Url{}, "ExternalID")

	return nil
}
