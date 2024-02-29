package models

type Todo struct {
	ID           string `json:"id"`
	Title     string `json:"title"`
	Status        string `json:"status"`
	UserId     string `json:"user_id"`
}

type TodoListItem struct {
	ID           string `json:"id"`
	Title     string `json:"title"`
	Status        string `json:"status"`
	UserId     string `json:"user_id"`
}

type TodoCreateModel struct {
	Title     string `json:"title"`
	Status        string `json:"status"`
	UserId     string `json:"user_id"`
}

type TodoUpdateModel struct {
	ID           string `json:"id"`
	Title     string `json:"title"`
	Status        string `json:"status"`
	UserId     string `json:"user_id"`
}
