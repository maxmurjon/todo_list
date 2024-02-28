package models

import "time"

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Article struct {
	ID        string        `json:"id"`
	Content              // Promoted fields
	AuthorID  int        `json:"author_id"` // Nested structs
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ArticleListItem struct {
	ID        string        `json:"id"`
	Content              // Promoted fields
	Author    Person     `json:"author"` // Nested structs
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ArticleCreateModel struct {
	Content      // Promoted fields
	AuthorID string `json:"author_id"`
}

type ArticleUpdateModel struct {
	ID       string `json:"id"`
	Content      // Promoted fields
	AuthorID int `json:"author_id"`
}
