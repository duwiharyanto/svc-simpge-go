package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

func InitGorm(db *sql.DB, withLog bool) (*gorm.DB, error) {
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// Logger: newLogger,
	}

	if withLog {
		cfg.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // Disable color
			},
		)
	}

	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: db,
		}),
		cfg,
	)
	if err != nil {
		return nil, fmt.Errorf("error connection gorm, %w", err)
	}
	return gormDB, nil
}
