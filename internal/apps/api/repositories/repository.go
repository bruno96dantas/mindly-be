package repositories

import "database/sql"

type RepoFields interface {
	GetDB() *sql.DB
}

type Repository struct{}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
