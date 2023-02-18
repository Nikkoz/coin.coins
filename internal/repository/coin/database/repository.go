package database

import (
	"coins/internal/useCase/adapters/storage"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB

		repoUrl storage.Url

		options Options
	}

	Options struct{}
)

func New(db *gorm.DB, s storage.Url, options Options) *Repository {
	repo := &Repository{
		db: db,

		repoUrl: s,
	}

	repo.SetOptions(options)

	return repo
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
