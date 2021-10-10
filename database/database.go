package database

import (
	"crawl_data/config"
	"crawl_data/database/entities"
	"crawl_data/model"
	"crawl_data/utils"
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
		log.Println("Error at ConnectToDatabase of database/database.go ", err)
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
		log.Println("Error at CreateDatabase of database/database.go ", err)
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

func ExecQueryDatabase(query string, db *sql.DB) error {
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error at ExecQueryDatabase of database/database.go ", err)
		return err
	}
	return err
}

func InsertStorageInforTbl(dbName string, infor entities.StorageInforTbl, db *sql.DB) error {
	// prepare
	query := `
		INSERT INTO ` + dbName + `.StorageInfor(Date, Type, LineID, HashCode) 
		VALUES (?, ?, ?, ?)
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error at InsertStorageInforTbl of database/database.go when prepare statement ", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(infor.Date, infor.Type, infor.LineID, infor.HashCode); err != nil {
		log.Println("Error at InsertStorageInforTbl of database/database.go when execute", err)
		return err
	}

	log.Println("Insert " + infor.Date + " into InsertStorageInforTbl successfully.")
	return err
}

func IsExist(date string, hashCode string, db *sql.DB) bool {
	info, err := SelectByDateAndInfor(date, hashCode, db)
	if err != nil {
		log.Println("Error IsExist of database/database.go", err)
	}
	if len(info) != 0 {
		return (info[0].Date != "" && info[0].HashCode != "")
	}
	return false
}

func SaveToDatabase(date string, format string, content []model.PageInformation, dbName string, db *sql.DB) {
	year, month, day := utils.GetDateDetail(date)
	date = year + "-" + month + "-" + day
	if format == "MD5" {
		for id, infor := range content {
			if IsExist(date, infor.MD5, db) || infor.MD5 == "" {
				continue
			}
			var newRow = entities.NewInfor(date, "MD5", id+1, infor.MD5)
			if err := InsertStorageInforTbl(dbName, newRow, db); err != nil {
				log.Println("Error insert MD5 at SaveToDatabase of database/database.go ", err)
			}
		}
	}
	if format == "SHA1" {
		for id, infor := range content {
			if IsExist(date, infor.SHA1, db) || infor.SHA1 == "" {
				continue
			}
			var newRow = entities.NewInfor(date, "SHA1", id+1, infor.SHA1)
			if err := InsertStorageInforTbl(dbName, newRow, db); err != nil {
				log.Println("Error insert SHA1 at SaveToDatabase of database/database.go ", err)
			}
		}
	}
	if format == "SHA256" {
		for id, infor := range content {
			if IsExist(date, infor.SHA256, db) || infor.SHA256 == "" {
				continue
			}
			var newRow = entities.NewInfor(date, "SHA256", id+1, infor.SHA256)
			if err := InsertStorageInforTbl(dbName, newRow, db); err != nil {
				log.Println("Error insert SHA256 at SaveToDatabase of database/database.go ", err)
			}
		}
	}
}
