package main

import (
	"database/sql"
	"log"
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"git.ultraware.nl/NiseVoid/qb/driver/pgqb"
	"git.ultraware.nl/NiseVoid/qb/qbdb"

	"github.com/caarlos0/env/v6"

	_ "github.com/lib/pq"
)

var qdb qbdb.DB

func main() {
	err := env.Parse(&settings)
	if err != nil {
		panic(err)
	}

	initDatabase()
	startRouter()
}

func initDatabase() {
	db, err := sql.Open(`postgres`, settings.Database.ConnectionString)
	if err != nil {
		panic(err)
	}
	qdb = pgqb.New(db)
}

func startRouter() {
	r := router.New()

	r.ErrorHandler = func(c *router.Context, i interface{}) {
		log.Println(i)
		_ = c.NoContent(http.StatusInternalServerError)
	}
	r.Use(middlewareCORS(settings.Cors.AllowOrigin))
	r.OPTIONS(`/*path`, func(c *router.Context) error {
		return c.NoContent(http.StatusOK)
	})

	r.GET(`/services`, getServices)
	r.GET(`/services/:id`, getService)

	r.GET(`/subscriptions`, getSubscriptions)
	r.PUT(`/subscriptions`, subscribe)
	r.DELETE(`/subscriptions/:id`, unsubscribe)

	panic(r.Start(`:8080`))
}
