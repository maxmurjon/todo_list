package postgres

import (
	"log"
	"todo/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sqlx.DB
	// article storage.ArticleRepoI
	// author  storage.AuthorRepoI
	user storage.UserRepoI
}

func NewPostgres(psqlConnString string) storage.StorageI {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = &userRepo{db: s.db}
	}
	return s.user
}

// func (s *Store) Author() storage.AuthorRepoI {
// 	if s.author == nil {
// 		s.author = &authorRepo{db: s.db}
// 	}
// 	return s.author
// }
