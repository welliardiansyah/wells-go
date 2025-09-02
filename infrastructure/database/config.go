package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var pool *sql.DB

func ConnectRawSQL(user, password, host, dbname string, port int) error {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	pool, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error creating DB pool: %v", err)
	}

	pool.SetMaxOpenConns(25)
	pool.SetMaxIdleConns(5)

	if err := pool.Ping(); err != nil {
		return fmt.Errorf("error pinging PostgreSQL DB: %v", err)
	}

	fmt.Println("PostgreSQL (Raw SQL) connection established")
	return nil
}

func GetPool() *sql.DB {
	return pool
}

func Close() error {
	if pool != nil {
		return pool.Close()
	}
	return nil
}
