package database

import (
	"crawl_data/config"
	"crawl_data/database/entities"
	"crawl_data/utils"
	"database/sql"
	"log"
)

func GetRowSelect(res *sql.Rows) (entities.StorageInforTbl, error) {
	var inforReturn entities.StorageInforTbl
	for res.Next() {
		var infor entities.StorageInforTbl
		err := res.Scan(&infor.Date, &infor.Type, &infor.LineID, &infor.HashCode)
		if err != nil {
			log.Println("Error at GetRowSelect of database/query.go", err)
			return inforReturn, err
		}
		inforReturn = entities.NewInfor(infor.Date, infor.Type, infor.LineID, infor.HashCode)
	}
	return inforReturn, nil
}

func SelectByDateAndInfor(date string, hashCode string, db *sql.DB) (entities.StorageInforTbl, error) {
	var (
		dbName      = config.GetConfig().DB_NAME
		inforReturn entities.StorageInforTbl
	)
	year, month, day := utils.GetDateDetail(date)
	date = year + "-" + month + "-" + day

	query := `
		SELECT *
		FROM ` + dbName + `.StorageInfor st
		WHERE st.Date = '` + date + `' AND st.HashCode= '` + hashCode + `'
	`
	//log.Println(query)
	res, err := db.Query(query)
	if err != nil {
		log.Println("Error at SelectByDateAndInfor of database/query.go when select", err)
		return inforReturn, err
	}

	if inforReturn, err = GetRowSelect(res); err != nil {
		log.Println("Error at SelectByDateAndInfor of database/query.go when getRowSelect", err)
		return inforReturn, err
	}
	return inforReturn, nil

}
