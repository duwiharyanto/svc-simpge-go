package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbms         = os.Getenv("DBMS")
	host         = os.Getenv("DB_HOST")
	port         = os.Getenv("DB_PORT")
	user         = os.Getenv("DB_USER")
	password     = os.Getenv("DB_PASSWORD")
	databaseName = os.Getenv("DB_NAME")
)

func Connect() (*sql.DB, error) {
	var dbSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=3s", user, password, host, port, databaseName)

	db, err := sql.Open(dbms, dbSource)
	if err != nil {
		return nil, err
	}

	dbMaxOpenConnection, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION"))
	if err != nil {
		// default
		dbMaxOpenConnection = 10
	}
	dbMaxIdleConnection, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION"))
	if err != nil {
		// default
		dbMaxIdleConnection = 20
	}
	dbMaxConnectionLifetime, err := time.ParseDuration(os.Getenv("DB_MAX_CONNECTION_LIFETIME"))
	if err != nil {
		// default
		dbMaxConnectionLifetime = 5 * time.Minute
	}

	// credit: Alex Edwards. https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(dbMaxOpenConnection)
	db.SetMaxIdleConns(dbMaxIdleConnection)
	db.SetConnMaxLifetime(dbMaxConnectionLifetime)

	return db, nil
}

func Healthz(ctx context.Context, db *sql.DB) error {
	var res string
	err := db.QueryRowContext(ctx, `SELECT COUNT(*) FROM personal_data_pribadi`).Scan(&res)
	if err != nil {
		return fmt.Errorf("error querying healthz, %w", err)
	}
	return nil
}
