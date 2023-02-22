package database

import (
	"coins/internal/useCase/adapters/storage"
	"gorm.io/gorm"
	"time"
)

type (
	Repository struct {
		db *gorm.DB

		repoUrl storage.Url

		options Options
	}

	Options struct {
		Timeout      time.Duration
		DefaultLimit uint64
	}
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
	if options.Timeout == 0 {
		options.Timeout = time.Second * 30
	}

	if options.DefaultLimit == 0 {
		options.DefaultLimit = 15
	}

	if r.options != options {
		r.options = options
	}
}
