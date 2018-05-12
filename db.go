package main

import (
	"database/sql"
	"strconv"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/olling/slog"
)

type Recipe struct {
	ID int
	Title string
	Added string
	Blog string
	Instructions string
}

func InitializeDBMigration() {
	db, err := sql.Open("postgres", "postgres://"+Conf.DatabaseConf.User+":"+Conf.DatabaseConf.Password+"@"+Conf.DatabaseConf.Host+":"+strconv.Itoa(Conf.DatabaseConf.Port)+"/"+Conf.DatabaseConf.DatabaseName+"?sslmode=disable")
	if err != nil {
		slog.PrintError("could not connect to database: ", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.PrintError("Driver error: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migrations",
		"postgres", driver)
	if err != nil {
		slog.PrintError("Migration error: ", err)
	}
	m.Up()
}
