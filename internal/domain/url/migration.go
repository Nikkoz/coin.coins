package url

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Url{})
	if err != nil {
		return err
	}

	//db.Migrator().AlterColumn(&Url{}, "ExternalID")

	return nil
}
