package postgres

import (
	"todo/models"
	"todo/storage"

	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) storage.CommentRepoI {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) CreateComment(entity models.CommentCreateModel) (string, error) {
	insertQuery := `INSERT INTO comments (
		task_id,
		user_id,
		text
	) VALUES (
		$1,
		$2,
		$3
	) RETURNING comment_id`

	var commentID string
	err := r.db.QueryRow(insertQuery,
		entity.TaskID,
		entity.UserID,
		entity.Text,
	).Scan(&commentID)

	return commentID, err
}

func (r *commentRepo) GetByCommentID(taskID string) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := r.db.Query(
		`SELECT comment_id, task_id, user_id, text, created_at
		 FROM comments
		 WHERE task_id = $1`,
		taskID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.TaskID, &comment.UserID, &comment.Text, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *commentRepo) DeleteComment(commentID string) error {
	_, err := r.db.Exec(`DELETE FROM comments WHERE comment_id = $1`, commentID)
	return err
}
