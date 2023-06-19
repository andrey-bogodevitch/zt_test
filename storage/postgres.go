package storage

import (
	"database/sql"
	"fmt"

	"zt_test/entity"

	"github.com/redis/go-redis/v9"
)

func PostgresRun(c entity.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PGhost, c.PGport, c.PGuser, c.PGpassword, c.PGname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Repository struct {
	db    *sql.DB
	cache *redis.Client
}

func New(dbpool *sql.DB, redis *redis.Client) *Repository {
	return &Repository{
		db:    dbpool,
		cache: redis,
	}
}

func (r *Repository) AddUser(user entity.User) error {
	query := "INSERT INTO users1 (name, age) values ($1, $2)"
	_, err := r.db.Exec(query, user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetLastUser() (entity.User, error) {
	query := "SELECT * FROM users1 WHERE id = (SELECT MAX(id) FROM users1);"

	var user entity.User

	err := r.db.QueryRow(query).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
	)
	if err != nil {
		return entity.User{}, nil
	}
	return user, nil
}
