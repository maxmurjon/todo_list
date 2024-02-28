package postgres

import (
	"database/sql"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"
)

var (
	db       *sql.DB
	fakeData *faker.Faker
)

func createRandomId(t *testing.T) string {
	id, err := uuid.NewRandom()
	assert.NoError(t, err)
	return id.String()
}

func TestMain(m *testing.M) {
	// cfg := config.Load()
	// db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresDatabase,
	// ))
	// if err != nil {
	// 	panic(err)
	// }

	db, err := sql.Open("postgres",
		"postgres://postgres:postgres@localhost:5432/article?sslmode=disable")
	if err != nil {
		panic(err)
	}

	fakeData, _ = faker.New("en")

	if err := db.Ping(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
