package database

import (
	"coins/internal/repositories/coin/interfaces"
	repository "coins/internal/repositories/url/interfaces"
	"gorm.io/gorm"
	"time"
)

var _ interfaces.Storage = (*Repository)(nil)

type (
	Repository struct {
		db *gorm.DB

		repoUrl repository.Storage

		options Options
	}

	Options struct {
		Timeout      time.Duration
		DefaultLimit uint64
	}
)

func New(db *gorm.DB, s repository.Storage, options Options) *Repository {
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
