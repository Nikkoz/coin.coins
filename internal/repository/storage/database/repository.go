package database

import "gorm.io/gorm"

type (
	Repository struct {
		db      *gorm.DB
		options Options
	}

	Options struct{}
)

func New(db *gorm.DB, options Options) *Repository {
	repo := &Repository{
		db: db,
	}

	repo.SetOptions(options)

	return repo
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
