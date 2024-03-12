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
	task   storage.TaskRepoI
	userRole storage.UserRoleRepoI
	comment storage.CommentRepoI
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

func (s *Store) Task() storage.TaskRepoI {
	if s.task == nil {
		s.task = &todoRepo{db: s.db}
	}
	return s.task
}

func (s *Store) UserRole() storage.UserRoleRepoI {
	if s.userRole == nil {
		s.userRole = &userRoleRepo{db: s.db}
	}
	return s.userRole
}

func (s *Store) Comment() storage.CommentRepoI {
	if s.comment == nil {
		s.comment = &commentRepo{db: s.db}
	}
	return s.comment
}
