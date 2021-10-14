package dal

import (
	"crawl_data/database/entities"
	"crawl_data/helpersCrawl/model"
	"crawl_data/helpersCrawl/utils"
	"database/sql"
	"log"
)

func SaveToDatabase(date string, format string, content []model.PageInformation, dbName string, db *sql.DB) {
	year, month, day := utils.GetDateDetail(date)
	date = year + "-" + month + "-" + day
	if format == "MD5" {
		for id, infor := range content {
			if IsExist(date, infor.MD5, db) || infor.MD5 == "" {
				continue
			}
			var newRow = entities.NewInfor(date, "MD5", id+1, infor.MD5)
			if err := InsertStorageInforTbl(newRow, db); err != nil {
				log.Println("Error insert MD5 at SaveToDatabase of database/dal/storeDB.go ", err)
			}
		}
	}
	if format == "SHA1" {
		for id, infor := range content {
			if IsExist(date, infor.SHA1, db) || infor.SHA1 == "" {
				continue
			}
			var newRow = entities.NewInfor(date, "SHA1", id+1, infor.SHA1)
			if err := InsertStorageInforTbl(newRow, db); err != nil {
				log.Println("Error insert SHA1 at SaveToDatabase of database/dal/storeDB.go ", err)
			}
		}
	}
	if format == "SHA256" {
		for id, infor := range content {
			if IsExist(date, infor.SHA256, db) || infor.SHA256 == "" {
				continue
			}
			var newRow = entities.NewInfor(date, "SHA256", id+1, infor.SHA256)
			if err := InsertStorageInforTbl(newRow, db); err != nil {
				log.Println("Error insert SHA256 at SaveToDatabase of database/dal/storeDB.go ", err)
			}
		}
	}
}
