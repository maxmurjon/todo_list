package postgres

import (
	"database/sql"
	"errors"
	"todo/models"
	"todo/storage"

	"github.com/jmoiron/sqlx"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) storage.AuthorRepoI {
	return authorRepo{
		db: db,
	}
}

func (r authorRepo) Create(entity models.PersonCreateModel) (err error) {
	insertQuery := `INSERT INTO author (
		firstname,
		lastname
	) VALUES (
		$1,
		$2
	)`

	_, err = r.db.Exec(insertQuery,
		entity.Firstname,
		entity.Lastname,
	)

	return err
}

func (r authorRepo) GetList(query models.Query) (resp []models.Person, err error) {

	var rows *sql.Rows
	rows, err = r.db.Query(
		`SELECT
			author.id, author.firstname, author.lastname, author.created_at, author.updated_at 
			FROM author 
			OFFSET $1 LIMIT $2`,
		query.Offset,
		query.Limit,
	)

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.Person
		err = rows.Scan(
			&a.ID, &a.Firstname, &a.Lastname, &a.CreatedAt, &a.UpdatedAt,
		)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r authorRepo) GetByID(ID string) (resp models.Person, err error) {
	var rows *sql.Rows
	rows, err = r.db.Query(
		`SELECT id, firstname, lastname, created_at, updated_at from author where id = $1
		`,
		ID,
	)

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.Person
		err = rows.Scan(
			&a.ID, &a.Firstname, &a.Lastname, &a.CreatedAt, &a.UpdatedAt,
		)
		resp = a
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r authorRepo) Update(entity models.PersonUpdateModel) error {

	query := `UPDATE author
	SET firstname = $2, lastname = $3
	WHERE id = $1
	`

	res, err := r.db.Exec(query,
		entity.ID,
		entity.Firstname,
		entity.Lastname,
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

func (r authorRepo) Delete(ID string) (effectedRowsNum int, err error) {
	insertQuery := `DELETE from author where id = $1`

	_, err = r.db.Exec(insertQuery,
		ID,
	)

	return 0, err

}

func (r authorRepo) GetMostPosted() (resp []models.PersonWhoMostPosted, err error) {

	// author firstname  | author lastname | artcile_name | article_created_at | article_updated_at | count

	return resp, err
}

func (r authorRepo) GetMostFewPosted() (resp []models.Person, err error) {
	return resp, err
}
