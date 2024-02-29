package postgres

import (
	"database/sql"
	"errors"
	"todo/models"
	"todo/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) storage.TodoRepoI {
	return todoRepo{
		db: db,
	}
}

func (r todoRepo) Create(entity models.TodoCreateModel) (string, error) {
	insertQuery := `INSERT INTO todo (
		id,
		title,
		user_id,
		status
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	uuid := uuid.New()
	_, err := r.db.Exec(insertQuery,
		uuid,
		entity.Title,
		entity.UserId,
		entity.Status,
	)
	return uuid.String(), err
}

func (r todoRepo) GetList(query models.Query) (resp []models.TodoListItem, err error) {

	var rows *sql.Rows
	rows, err = r.db.Query(
		`SELECT id,
		title,
		user_id,
		status from todo
				OFFSET $1 LIMIT $2`,
		query.Offset,
		query.Limit,
	)

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.TodoListItem
		err = rows.Scan(
			&a.ID, &a.Title, &a.Status, &a.UserId,
		)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r todoRepo) GetByID(ID string) (models.Todo, error) {
	var resp models.Todo
	err := r.db.QueryRow(
		`SELECT id,
		title,
		user_id,
		status from todo where id = $1
		`,
		ID,
	).Scan(&resp.ID, &resp.Title, &resp.Status, &resp.UserId)

	if err != nil {
		return resp, err
	}

	return resp, err
}

func (r todoRepo) Update(entity models.TodoUpdateModel) error {

	query := `UPDATE todo
	SET title=$2,
	user_id=$3,
	status=$4
	WHERE id = $1
	`

	res, err := r.db.Exec(query,
		entity.ID,
		entity.Title,
		entity.UserId,
		entity.Status,
	)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()

	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}
	return errors.New("article not updated")
}

func (r todoRepo) Delete(ID string) (effectedRowsNum int, err error) {
	insertQuery := `DELETE from todo where id = $1`

	_, err = r.db.Exec(insertQuery,
		ID,
	)

	return 0, err

}
