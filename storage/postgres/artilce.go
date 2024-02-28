package postgres

import (
	"bootcamp/article/models"
	"bootcamp/article/storage"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(db *sqlx.DB) storage.ArticleRepoI {
	return articleRepo{
		db: db,
	}
}

func (r articleRepo) Create(entity models.ArticleCreateModel) (err error) {
	insertQuery := `INSERT INTO article (
		title,
		body,
		author_id
	) VALUES (
		$1,
		$2,
		$3
	)`

	_, err = r.db.Exec(insertQuery,
		entity.Title,
		entity.Body,
		entity.AuthorID,
	)

	return err
}

func (r articleRepo) GetList(query models.Query) (resp []models.ArticleListItem, err error) {
	var rows *sql.Rows
		rows, err = r.db.Query(
			`SELECT
			ar.id, ar.title, ar.body, ar.created_at, ar.updated_at,
			au.id, au.firstname, au.lastname, au.created_at, au.updated_at
			FROM article AS ar JOIN author AS au ON ar.author_id = au.id
			OFFSET $1 LIMIT $2`,
			query.Offset,
			query.Limit,
		)
	

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.ArticleListItem
		err = rows.Scan(
			&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.UpdatedAt,
			&a.Author.ID, &a.Author.Firstname, &a.Author.Lastname, &a.Author.CreatedAt, &a.Author.UpdatedAt,
		)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r articleRepo) GetByID(ID string) (resp models.Article, err error) {
	
	var rows *sql.Rows
	rows, err = r.db.Query(
		`SELECT id, title, body, created_at, updated_at from article where id = $1
		`,
		ID,
	)

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.Article
		err = rows.Scan(
			&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.UpdatedAt,
		)
		resp = a
		if err != nil{
			return resp, err
		}
	}

	return resp, err
}

func (r articleRepo) Update(entity models.ArticleUpdateModel) ( int64, error) {
		query := `UPDATE article
		SET title = $2, body = $3
		WHERE id = $1
		`
	
		 res, err := r.db.Exec(query,
			entity.ID,
			entity.Title,
			entity.Body,
	)
	if err != nil {
		return 0,err
	}
	n, err := res.RowsAffected()
	
	if err!= nil {
		return 0,err
	}
	if n>0{
		return 0,nil
	}
	return n,errors.New("article not updated")

}


func (r articleRepo) Delete(ID string) (effectedRowsNum int, err error) {

	insertQuery := `DELETE from article where id = $1`

	_, err = r.db.Exec(insertQuery,
		ID,
	)

	return 0,err

}
