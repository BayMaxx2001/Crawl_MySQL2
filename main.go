package main

import (
	"crawl_data/config"
	"crawl_data/database"
	"crawl_data/model"
	"crawl_data/page"
	"crawl_data/utils"
	"fmt"
	"time"

	"database/sql"
	"log"
	"sync"
)

var url string = "https://malshare.com/daily/"

func checkError(msg string, err error) {
	if err != nil {
		log.Panic(err)
	}
}

func sendDateToChanDate(chanDate chan<- string, listDate []string) {
	for i := range listDate {
		chanDate <- listDate[i]
	}
	defer close(chanDate)
}

func processInformationOfDate(chanDate chan string, chanPageInformation chan []model.PageInformation, wg *sync.WaitGroup) {
	for date := range chanDate {
		info, err := page.GetInformationOfDate(date)
		if err != nil {
			log.Println("Error at processInformationOfDate of crawlData/main.go ", err)
			continue
		}
		chanPageInformation <- info
	}
	defer wg.Done()
}

func saveOutput(chanPageInformation chan []model.PageInformation, wg *sync.WaitGroup, dbName string, db *sql.DB) {
	for info := range chanPageInformation {
		day := info[0].DAY
		month := info[0].MONTH
		year := info[0].YEAR

		date := utils.GetDateToString(day, month, year)
		go database.SaveToDatabase(date, "MD5", info, dbName, db)
		go database.SaveToDatabase(date, "SHA1", info, dbName, db)
		go database.SaveToDatabase(date, "SHA256", info, dbName, db)
	}
	defer wg.Done()
}

func crawlData(dbName string, db *sql.DB) {
	var (
		chanPageInformation        = make(chan []model.PageInformation, 500)
		chanDate                   = make(chan string, 500)
		listDate                   []string
		numOfWorkerDate            = 100
		numOfWorkerPageInformation = 100
		err                        error
	)
	listDate, err = utils.GetListDate(url)
	checkError("Get list date of crawlData/main.go", err)
	go sendDateToChanDate(chanDate, listDate)
	var wg1, wg2 sync.WaitGroup
	for i := 0; i < numOfWorkerPageInformation; i++ {
		wg1.Add(1)
		go saveOutput(chanPageInformation, &wg1, dbName, db)
	}
	for i := 0; i < numOfWorkerDate; i++ {
		wg2.Add(1)
		go processInformationOfDate(chanDate, chanPageInformation, &wg2)
	}
	wg2.Wait()
	close(chanPageInformation)
	wg1.Wait()
}

func connectDBAndCreateTBL() *sql.DB {
	var dbName = config.GetConfig().DB_NAME
	// Connect to MySQL
	db, err := database.ConnectToDatabase("")
	checkError("Connect to MySql at connectDBAndCreateTBL of crawlData/main.go", err)

	// Create database
	db, err = database.CreateDatabase(dbName, db)
	checkError("Create database:"+dbName+" at connectDBAndCreateTBL of crawlData/main.go", err)
	// Connect to database
	db, err = database.ConnectToDatabase(dbName)
	checkError("Connect to database:"+dbName+" at connectDBAndCreateTBL of crawlData/main.go", err)

	// Create table in database
	db, err = database.CreateTable("INFORMATION", db)
	checkError("Create table INFORMATION at connectDBAndCreateTBL of crawlData/main.go", err)

	return db
}

func main() {
	var dbName = config.GetConfig().DB_NAME
	start := time.Now()
	// Crawl data into the database
	db := connectDBAndCreateTBL()
	crawlData(dbName, db)
	fmt.Println("Time to run program: ", time.Since(start))
}
