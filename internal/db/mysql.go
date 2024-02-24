package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	address  = "localhost:3306"
	user     = "root"
	password = ""
	dbname   = "test"
)

func InitMySQLConn() *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, address, dbname)
	dbConn, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to connect mysql, error: %v", err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("failed to connect mysql, error: %v", err)
	}

	return dbConn
}
