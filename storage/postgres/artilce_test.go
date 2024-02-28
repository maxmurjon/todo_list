package postgres

import (
	"bootcamp/article/models"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func Test_articleRepo_Create(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		entity models.ArticleCreateModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TO DO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := articleRepo{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("articleRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_articleRepo_GetList(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		query models.Query
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp []models.ArticleListItem
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := articleRepo{
				db: tt.fields.db,
			}
			gotResp, err := r.GetList(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleRepo.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("articleRepo.GetList() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
