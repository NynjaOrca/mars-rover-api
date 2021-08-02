package api

import (
	"mars-rover-api/mars-rover/pkg/config"
	"mars-rover-api/mars-rover/pkg/db"
	"mars-rover-api/mars-rover/pkg/nasaclient"
)

type API struct {
	Config     config.Config
	NasaClient nasaclient.Client
	Postgres   *db.DAO
}

func New(c config.Config, dao db.DAO) (*API, error) {
	cli := nasaclient.New(c.NasaClient)

	return &API{
		Config:     c,
		NasaClient: cli,
		Postgres:   &dao,
	}, nil
}
