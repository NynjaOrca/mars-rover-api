package api

import (
	"fmt"
	"log"
	"mars-rover-api/mars-rover/pkg/config"
	"mars-rover-api/mars-rover/pkg/db"
	"mars-rover-api/mars-rover/pkg/nasaclient"
	"time"
)

const (
	earthDateFormat = "2006-01-02"
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

func buildMap() map[string][]string {
	t := time.Now()
	m := make(map[string][]string, 0)
	m[t.Format(earthDateFormat)] = []string{}
	fmt.Println("i:", 0, "t:", t)

	for i := 1; i < 10; i++ {
		t = t.Add(time.Hour * -24)
		fmt.Println("i:", i, "t:", t)
		m[t.Format(earthDateFormat)] = []string{}
	}
	fmt.Println(m)
	return m
}

func (a *API) LastTenDays(camera string) {
	m := buildMap()
	for k, v := range buildMap() {
		fmt.Println("k", k, "m", v)
		images, err := a.Postgres.FetchRoverImages(k, camera)
		if err != nil {
			log.Println("api - error fetching images:", err)
		}
		if len(images) == 0 {
			images = a.NasaClient.GetRoverImages(k, camera)
			if err != nil {
				log.Println("unable to get rover images:", err)
			}
			err = a.Postgres.InsertRoverImages(images, k, camera)
			if err != nil {
				log.Println("unable to insert rover images:", err)
			}
		}

		m[k] = images
	}
	fmt.Println("final map:", m)
}
