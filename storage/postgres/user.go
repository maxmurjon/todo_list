package postgres

import (
	"database/sql"
	"errors"
	"todo/models"
	"todo/storage"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) storage.UserRepoI {
	return userRepo{
		db: db,
	}
}

func (r userRepo) Create(entity models.UserCreateModel) (err error) {
	insertQuery := `INSERT INTO author (
		full_name,
		user_name,
		email,
		user_password
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	_, err = r.db.Exec(insertQuery,
		entity.FullName,
		entity.UserName,
		entity.Email,
		entity.UserPassword,
	)

	return err
}

// func (r userRepo) GetList(query models.Query) (resp []models.Person, err error) {

// 	var rows *sql.Rows
// 	rows, err = r.db.Query(
// 		`SELECT
// 			author.id, author.firstname, author.lastname, author.created_at, author.updated_at 
// 			FROM author 
// 			OFFSET $1 LIMIT $2`,
// 		query.Offset,
// 		query.Limit,
// 	)

// 	if err != nil {
// 		return resp, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		var a models.Person
// 		err = rows.Scan(
// 			&a.ID, &a.Firstname, &a.Lastname, &a.CreatedAt, &a.UpdatedAt,
// 		)
// 		resp = append(resp, a)
// 		if err != nil {
// 			return resp, err
// 		}
// 	}

// 	return resp, err
// }

func (r userRepo) GetByID(ID string) (resp models.User, err error) {
	var rows *sql.Rows
	rows, err = r.db.Query(
		`SELECT id, full_name,
		user_name,
		email,
		user_password from users where id = $1
		`,
		ID,
	)

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.User
		err = rows.Scan(
			&a.Id, &a.FullName, &a.UserName, &a.Email, &a.UserPassword,
		)
		resp = a
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r userRepo) Update(entity models.UserUpdateModel) error {

	query := `UPDATE author
	SET full_name=$2,
	user_name=$3,
	email=$4,
	user_password=$5
	WHERE id = $1
	`

	res, err := r.db.Exec(query,
		entity.Id,
		entity.FullName,
		entity.UserName,
		entity.Email,
		entity.UserPassword,
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

func (r userRepo) Delete(ID string) (effectedRowsNum int, err error) {
	insertQuery := `DELETE from user where id = $1`

	_, err = r.db.Exec(insertQuery,
		ID,
	)

	return 0, err

}