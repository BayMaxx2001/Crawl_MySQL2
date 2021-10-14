package dal

import (
	"crawl_data/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Data Source Name
func DSN(dbName string) string {
	var (
		username = config.GetConfig().DB_USERNAME
		hostname = config.GetConfig().DB_HOST
		password = config.GetConfig().DB_PASSWORD
		port     = config.GetConfig().DB_PORT
	)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbName)
}

func ConnectToDatabase(dbName string) (*sql.DB, error) {

	DSN := DSN(dbName)
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Println("Error at ConnectToDatabase of database/dal/database.go ", err)
		return db, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db, err
}

func CreateDatabase(dbName string, db *sql.DB) error {
	query := "CREATE DATABASE IF NOT EXISTS " + dbName
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error at CreateDatabase of database/dal/database.go ", err)
		return err
	}
	return err
}
func CreateTable(tableName string, db *sql.DB) error {
	queryInforTbl := `
	CREATE TABLE IF NOT EXISTS StorageInfor (
		Date        DATE NOT NULL,
		Type        NVARCHAR(10) NOT NULL,
		LineID    	INT NOT NULL,
		HashCode    TEXT NOT NULL,
		PRIMARY KEY(DATE, Type, LineID),
		INDEX (Date,HashCode(256))
	);
	`
	return ExecQueryDatabase(queryInforTbl, db)
}
