package repository

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/GangOfThrees/Obi-wan/kenobi/internal/constants"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/repository/db_queries"
	_ "github.com/lib/pq"
)

var (
	obiWanDb *sql.DB
	queries  *db_queries.Queries
)

func SetupDatabase() error {
	db, err := sql.Open(os.Getenv(constants.DB_FLAVOUR), os.Getenv(constants.DB_CONN_STRING))
	if err != nil {
		return err
	}
	obiWanDb = db

	obiWanDb.SetConnMaxLifetime(30 * time.Minute)

	maxOpenConns := 10
	if n, err := strconv.Atoi(os.Getenv(constants.DB_MAX_CONNS)); err == nil {
		maxOpenConns = n
	}
	obiWanDb.SetMaxOpenConns(maxOpenConns)

	maxIdleConns := 10
	if n, err := strconv.Atoi(os.Getenv(constants.DB_MAX_CONNS)); err == nil {
		maxIdleConns = n
	}
	obiWanDb.SetMaxIdleConns(maxIdleConns)

	queries = db_queries.New(obiWanDb)

	return nil
}

func GetDbQueries() (*db_queries.Queries, *sql.DB) {
	return queries, obiWanDb
}
