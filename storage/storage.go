package storage

import "todo/models"

type StorageI interface {
	User() UserRepoI
	Todo() TodoRepoI
}

type UserRepoI interface {
	Create(entity models.UserCreateModel) (uuid string, err error)
	GetList(query models.Query) (resp []models.UserListItem, err error)
	GetByID(ID string) (resp models.User, err error)
	Update(entity models.UserUpdateModel) (err error)
	Delete(ID string) (effectedRowsNum int, err error)
}

type TodoRepoI interface {
	Create(entity models.TodoCreateModel) (uuid string, err error)
	GetList(query models.Query) (resp []models.TodoListItem, err error)
	GetByID(ID string) (resp models.Todo, err error)
	Update(entity models.TodoUpdateModel) error
	Delete(ID string) (effectedRowsNum int, err error)
}
