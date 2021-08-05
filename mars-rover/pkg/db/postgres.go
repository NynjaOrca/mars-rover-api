package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"mars-rover-api/mars-rover/pkg/config"
)

type Postgres struct {
	Config config.Postgres
	DB     *sql.DB
}

type DAO interface {
	InsertRoverImages(images []string, earthDate, camera string) error
	FetchRoverImages(earthDate, camera string) ([]string, error)
	EnsureTables() error
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

func (p *Postgres) FetchRoverImages(earthDate, camera string) ([]string, error) {
	q := `SELECT * FROM daily_images WHERE earth_date = $1 and camera = $2`
	rows, err := p.DB.Query(q, earthDate, camera)
	if err != nil {
		return nil, err
	}

	var ed string
	var cm string
	var img images
	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&ed, &cm, &img)
			if err != nil {
				return nil, err
			}
		}
	}
	return img, nil
}

func (p *Postgres) InsertRoverImages(images []string, earthDate, camera string) error {
	q := `INSERT INTO daily_images(earth_date, camera, images) 
			VALUES($1, $2, $3) ON CONFLICT (earth_date)
			DO NOTHING`

	b, err := json.Marshal(images)
	result, err := p.DB.Exec(q, earthDate, camera, b)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	} else {
		if n != 0 {
			log.Println("number of rows inserted into mongo:", n)
		}
	}
	return nil
}
