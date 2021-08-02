package api

import (
	"fmt"
	"log"
	"mars-rover-api/mars-rover/pkg/config"
	"mars-rover-api/mars-rover/pkg/db"
	"mars-rover-api/mars-rover/pkg/nasaclient"
)

type API struct {
	Config     config.Config
	NasaClient nasaclient.Client
	Postgres   db.DAO
}

func New(c config.Config, dao db.DAO) (*API, error) {
	cli := nasaclient.New(c.NasaClient)

	return &API{
		Config:     c,
		NasaClient: cli,
		Postgres:   dao,
	}, nil
}

func (a *API) LastTenDays(earthDate, camera string) {
	err := a.NasaClient.GetRoverImages(earthDate, camera)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to get rover images: %v", err))
	}

	err = a.Postgres.FetchRoverImages(earthDate, camera)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to get rover images from postgres: %v", err))
	}
}
