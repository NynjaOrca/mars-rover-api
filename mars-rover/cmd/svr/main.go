package main

import (
	"fmt"
	"log"
	"mars-rover-api/mars-rover/pkg/api"
	"mars-rover-api/mars-rover/pkg/config"
	"mars-rover-api/mars-rover/pkg/db"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(fmt.Errorf("unable to build config: %v", err))
	}

	dao, err := db.New(conf.Postgres)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to build dao: %v", err))
	}

	app, err := api.New(conf, dao)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to build api: %v", err))
	}

	d := app.LastTenDays("NAVCAM")
	_ = dao.DB.Close()
	log.Println(d)
}
