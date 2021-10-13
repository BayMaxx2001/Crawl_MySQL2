package main

import (
	"crawl_data/config"
	"crawl_data/database"
	"crawl_data/helpers/model"
	"crawl_data/helpers/page"
	"crawl_data/helpers/utils"
	"database/sql"
	"log"
	"sync"
)

var url string = "https://malshare.com/daily/"

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
	if err != nil {
		log.Panic("Get list date of crawlData/main.go ", err)
	}
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

func main() {
	var dbName = config.GetConfig().DB_NAME
	//connect database
	db, err := database.ConnectToDatabase(dbName)
	if err != nil {
		log.Panic("Error at main of crawlData.go when connect to database")
	}
	crawlData(dbName, db)
}
