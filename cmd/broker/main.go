package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	"monitor/internal/migration"
	"monitor/service"

	"git.ultraware.nl/NiseVoid/qb/driver/pgqb"
	"git.ultraware.nl/NiseVoid/qb/qbdb"

	"github.com/caarlos0/env/v6"

	_ "github.com/lib/pq"
)

var (
	qdb qbdb.DB
)

func main() {
	err := env.Parse(&settings)
	if err != nil {
		panic(err)
	}

	initDB()

	go func() {
		for {
			func() {
				defer func() {
					v := recover()
					if v != nil {
						log.Println(v)
					}
				}()

				services := service.GetAll(qdb)
				services = updateServices(services)

				markUnavailableAsDead(services)
			}()

			time.Sleep(time.Second * 5)
		}
	}()

	// TODO: Listen for subscription in the database
	// listenForSubscriptions()

	// TODO: update subscriptions based on the database
	// updateSubsciptions()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}

func initDB() {
	db, err := sql.Open(`postgres`, settings.Database.ConnectionString)
	if err != nil {
		panic(err)
	}

	migration.Migrate(db)
	qdb = pgqb.New(db)
}
