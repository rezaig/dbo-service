package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func InitMySQLConn() *sql.DB {
	mysqlConf := viper.Sub("mysql")
	var (
		host     = mysqlConf.GetString("host")
		username = mysqlConf.GetString("username")
		password = mysqlConf.GetString("password")
		database = mysqlConf.GetString("database")
	)

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, database)
	dbConn, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to connect mysql, error: %v", err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetConnMaxIdleTime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("failed to connect mysql, error: %v", err)
	}

	return dbConn
}
