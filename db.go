package main

import (
	"database/sql"
	"strconv"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/olling/slog"
)

type Recipe struct {
	ID           int
	Title        string
	Added        string
	Blog         string
	Instructions string
}

func InitializeDBMigration() {
	db, err := getDbConnection()
	if err != nil {
		slog.PrintError("could not connect to database: ", err)
	}
	defer db.Close()

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

func bulkInsertRecipes(unsavedRows []Recipe) error {
	db := getDbConnection()
	defer db.Close()

	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*4)
	for _, post := range unsavedRows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*4+1, i*4+2, i*4+3))
		valueArgs = append(valueArgs, post.Title)
		valueArgs = append(valueArgs, post.Added)
		valueArgs = append(valueArgs, post.Blog)
		valueArgs = append(valueArgs, post.Instructions)
	}
	stmt := fmt.Sprintf("INSERT INTO recipes (title, added, blog_id, instrutions_id) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func getDbConnection() (*sqlx.DB, error) {
	//Connect to database
	db, err := sqlx.Connect("postgres", "postgres://"+Conf.DatabaseConf.User+":"+Conf.DatabaseConf.Password+"@"+Conf.DatabaseConf.Host+":"+strconv.Itoa(Conf.DatabaseConf.Port)+"/"+Conf.DatabaseConf.DatabaseName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
