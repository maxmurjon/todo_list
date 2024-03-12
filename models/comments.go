package models



type Comment struct {
	ID string `json:"comment_id"`
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
	Text string `json:"text"`
	CreatedAt string `json:"created_at"`
}


type CommentListItem struct {
	
	ID string `json:"comment_id"`
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
	Text string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type CommentCreateModel struct {
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
	Text string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type CommentUpdateModel struct {
	ID string `json:"comment_id"`
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
	Text string `json:"text"`
	CreatedAt string `json:"created_at"`
}
