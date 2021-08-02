package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"mars-rover-api/mars-rover/pkg/config"
)

type Postgres struct {
	Config config.Postgres
	DB     *sql.DB
}

type DAO interface {
}

func New(c config.Postgres) (*Postgres, error) {
	psqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	db, err := sql.Open("postgres", psqlConnection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Config: c,
		DB:     db,
	}, nil
}
