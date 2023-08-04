package scoupes

import "gorm.io/gorm"

func Paginate(limit uint64, page uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * limit

		return db.Offset(int(offset)).Limit(int(limit))
	}
}
