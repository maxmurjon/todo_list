package postgres

import (
	"todo/models"
	"todo/storage"

	"github.com/jmoiron/sqlx"
)

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) storage.TaskRepoI {
	return &todoRepo{
		db: db,
	}
}

// CRUD funksiyalari quyidagicha bo'lishi kerak

func (r *todoRepo) Create(entity models.TaskCreateModel) (string, error) {
	insertQuery := `INSERT INTO tasks (
		title,
		description,
		due_date,
		status,
		user_id
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING task_id`

	var id string
	err := r.db.QueryRow(insertQuery,
		entity.Title,
		entity.Description,
		entity.DueDate,
		entity.Status,
		entity.UserID,
	).Scan(&id)

	return id, err
}

func (r *todoRepo) GetList(query models.Query) ([]models.TaskListItem, error) {
	var tasks []models.TaskListItem

	rows, err := r.db.Query(
		`SELECT task_id, title, description, due_date, status, user_id
		 FROM tasks
		 OFFSET $1 LIMIT $2`,
		query.Offset,
		query.Limit,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.TaskListItem
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status, &task.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *todoRepo) GetByID(ID string) (models.Task, error) {
	var task models.Task
	err := r.db.QueryRow(
		`SELECT task_id, title, description, due_date, status, user_id
		 FROM tasks
		 WHERE task_id = $1`,
		ID,
	).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status, &task.UserID)

	return task, err
}

func (r *todoRepo) Update(entity models.TaskUpdateModel) error {
	query := `UPDATE tasks
	SET title = $2, description = $3, due_date = $4, status = $5, user_id = $6
	WHERE task_id = $1`

	_, err := r.db.Exec(query,
		entity.ID,
		entity.Title,
		entity.Description,
		entity.DueDate,
		entity.Status,
		entity.UserID,
	)

	return err
}

func (r *todoRepo) Delete(ID string) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE task_id = $1`, ID)
	return err
}
