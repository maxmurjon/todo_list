package models

type User struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserListItem struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserCreateModel struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserUpdateModel struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}
