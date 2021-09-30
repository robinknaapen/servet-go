package migration

import (
	"database/sql"
	"embed"

	"git.fuyu.moe/Fuyu/migrate/v2"
)

//go:embed *.sql
var migrations embed.FS

// Migrate migrates the database
func Migrate(db *sql.DB) {
	err := migrate.Migrate(db, migrate.Options{}, migrations)
	if err != nil {
		panic(err)
	}
}
