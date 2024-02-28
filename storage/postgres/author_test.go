package postgres

import (
	"database/sql/driver"
	"testing"
	"time"
	"todo/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestAuthorRepo_Create(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Set up the test case data
	testAuthor := models.PersonCreateModel{
		Firstname: fakeData.FirstName(),
		Lastname:  fakeData.LastName(),
	}

	// Set up the expected SQL query and parameters
	expectedQuery := "INSERT INTO author \\(firstname, lastname\\) VALUES \\(\\?, \\?\\)"
	expectedArgs := []driver.Value{testAuthor.Firstname, testAuthor.Lastname}

	// Set up the mock database to expect the SQL query with the specified parameters
	mock.ExpectExec(expectedQuery).WithArgs(expectedArgs...).WillReturnResult(sqlmock.NewResult(0, 1))

	// Create the author repository and inject the mock database
	db := sqlx.NewDb(mockDB, "sqlmock")
	repo := authorRepo{db: db}

	// Call the Create method
	err = repo.Create(testAuthor)
	if err != nil {
		t.Errorf("error calling Create method: %v", err)
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestAuthorRepo_GetList(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Set up the test case data
	testQuery := models.Query{
		Offset: 0,
		Limit:  10,
	}

	// Set up the expected SQL query and parameters
	expectedQuery := "SELECT author.id, author.firstname, author.lastname, author.created_at, author.updated_at FROM author OFFSET \\? LIMIT \\?"
	expectedArgs := []driver.Value{testQuery.Offset, testQuery.Limit}

	// Set up the mock database to expect the SQL query with the specified parameters
	mock.ExpectQuery(expectedQuery).WithArgs(expectedArgs...).WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "created_at", "updated_at"}).
		AddRow(1, fakeData.FirstName, fakeData.LastName, time.Now(), time.Now()).
		AddRow(2, fakeData.FirstName, fakeData.LastName, time.Now(), time.Now()))

	// Create the author repository and inject the mock database
	db := sqlx.NewDb(mockDB, "sqlmock")
	repo := authorRepo{db: db}

	// Call the GetList method
	resp, err := repo.GetList(testQuery)
	if err != nil {
		t.Errorf("error calling GetList method: %v", err)
	}

	// Verify the response
	if len(resp) != 2 {
		t.Errorf("expected 2 rows, got %d", len(resp))
	}
}

func TestAuthorRepo_GetByID(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Set up the test case data
	testID := "1"
	testAuthor := models.Person{
		ID:        "1",
		Firstname: fakeData.FirstName(),
		Lastname:  fakeData.LastName(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set up the expected SQL query and parameters
	expectedQuery := "SELECT id, firstname, lastname, created_at, updated_at from author where id = \\?"
	expectedArgs := []driver.Value{testID}

	// Set up the mock database to expect the SQL query with the specified parameters
	mock.ExpectQuery(expectedQuery).WithArgs(expectedArgs...).WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "created_at", "updated_at"}).
		AddRow(testAuthor.ID, testAuthor.Firstname, testAuthor.Lastname, testAuthor.CreatedAt, testAuthor.UpdatedAt))

	// Create the author repository and inject the mock database
	db := sqlx.NewDb(mockDB, "sqlmock")
	repo := authorRepo{db: db}

	// Call the GetByID method
	resp, err := repo.GetByID(testID)
	if err != nil {
		t.Errorf("error calling GetByID method: %v", err)
	}

	// Verify the response
	if resp != testAuthor {
		t.Errorf("unexpected response. expected: %v, got: %v", testAuthor, resp)
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestAuthorRepo_Delete(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Set up the test case data
	testID := "1"

	// Set up the expected SQL query and parameters
	expectedQuery := "DELETE from author where id = \\?"
	expectedArgs := []driver.Value{testID}

	// Set up the mock database to expect the SQL query with the specified parameters
	mock.ExpectExec(expectedQuery).WithArgs(expectedArgs...).WillReturnResult(sqlmock.NewResult(0, 1))

	// Create the author repository and inject the mock database
	db := sqlx.NewDb(mockDB, "sqlmock")
	repo := authorRepo{db: db}

	// Call the Delete method
	rowsAffected, err := repo.Delete(testID)
	if err != nil {
		t.Errorf("error calling Delete method: %v", err)
	}

	// Verify that the number of affected rows is as expected
	if rowsAffected != 1 {
		t.Errorf("expected 1 row affected, but got %d", rowsAffected)
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func Test_authorRepo_Create(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		entity models.PersonCreateModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := authorRepo{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("authorRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
