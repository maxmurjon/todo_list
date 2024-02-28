package models


type User struct{
	Id string `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	UserName string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserListItem struct {
	ID        string        `json:"id"`
	Todo              // Promoted fields
	Author    Person     `json:"author"` // Nested structs
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserCreateModel struct{
	FullName string `json:"full_name"`
	Email string `json:"email"`
	UserName string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type UserUpdateModel struct{
	Id string `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	UserName string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

