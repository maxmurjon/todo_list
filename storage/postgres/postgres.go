package postgres

import (
	"log"
	"todo/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db     *sqlx.DB
	user   storage.UserRepoI
	todo   storage.TodoRepoI
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

func (s *Store) Todo() storage.TodoRepoI {
	if s.todo == nil {
		s.todo = &todoRepo{db: s.db}
	}
	return s.todo
}
