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
	FetchRoverImages(earthDate, camera string) error
}

func New(c config.Postgres) (*Postgres, error) {
	psqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.SSLMode)
	db, err := sql.Open("postgres", psqlConnection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p := &Postgres{
		Config: c,
		DB:     db,
	}

	err = p.EnsureTables()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Postgres) EnsureTables() error {
	q := `
	CREATE TABLE IF NOT EXISTS daily_images (
  		earth_date varchar(45) NOT NULL,
  		camera varchar(45) NOT NULL,
  		images jsonb NOT NULL,
  		PRIMARY KEY (earth_date)
	)`
	_, err := p.DB.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) FetchRoverImages(earthDate, camera string) error {
	q := `SELECT * FROM daily_images WHERE earth_date = $1 and camera = $2`
	rows, err := p.DB.Query(q, earthDate, camera)
	if err != nil {
		return err
	}

	if rows != nil {
		for rows.Next() {
			// TODO: scan each row into a model
		}
	}
	return nil
}
