package storage

import "todo/models"

type StorageI interface {
	User() UserRepoI
	Task() TaskRepoI
	UserRole() UserRoleRepoI
	Comment() CommentRepoI
}

type UserRepoI interface {
	Create(entity models.UserCreateModel) (uuid string, err error)
	GetList(query models.Query) (resp []models.UserListItem, err error)
	GetByID(ID string) (resp models.User, err error)
	Update(entity models.UserUpdateModel) (err error)
	Delete(ID string) (effectedRowsNum int, err error)
}

type TaskRepoI interface {
	Create(entity models.TaskCreateModel) (uuid string, err error)
	GetList(query models.Query) (resp []models.TaskListItem, err error)
	GetByID(ID string) (resp models.Task, err error)
	Update(entity models.TaskUpdateModel) error
	Delete(ID string) error
}

type UserRoleRepoI interface {
	GetUserRoles(userID string) ([]models.UserRole, error)
	CreateUserRole(userID string, roleID int) error
	DeleteUserRole(userID string, roleID int) error
}

type UserRoleRoleRepoI interface {
	GetUserRoles(userID string) ([]models.UserRole, error)
	CreateUserRole(userID string, roleID int) error
	DeleteUserRole(userID string, roleID int) error
}


type CommentRepoI interface {
	CreateComment(entity models.CommentCreateModel) (string, error)
	GetByCommentID(taskID string) ([]models.Comment, error)
	DeleteComment(commentID string) error
}
