package main

import (
	"crawl_data/config"
	"crawl_data/database/dal"
	"log"

	"database/sql"
)

func checkError(msg string, err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Connect to MySQL
func connectMySql() *sql.DB {
	db, err := dal.ConnectToDatabase("")
	checkError("Connect to MySql at connectMySql of setupDB.go", err)
	return db
}

// Create database
func createDB() {
	var (
		dbName = config.GetConfig().DB_NAME
		db     *sql.DB
		err    error
	)
	// Connect to mysql
	db = connectMySql()
	// Create database
	err = dal.CreateDatabase(dbName, db)
	checkError("Create database:"+dbName+" at connectDBAndCreateTBL of crawlData/main.go", err)
}

// Create table in database
func createTable() {
	var (
		dbName = config.GetConfig().DB_NAME
		db     *sql.DB
		err    error
	)
	// Connect to database
	db, err = dal.ConnectToDatabase(dbName)
	checkError("Connect to database:"+dbName+" at connectDBAndCreateTBL of crawlData/main.go", err)
	// Create table in database
	err = dal.CreateTable("INFORMATION", db)
	checkError("Create table INFORMATION at connectDBAndCreateTBL of crawlData/main.go", err)
}

func main() {
	connectMySql()
	createDB()
	createTable()
}
