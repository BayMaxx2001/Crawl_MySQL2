package database

import (
	"crawl_data/config"
	"crawl_data/database/entities"
	"crawl_data/utils"
	"database/sql"
	"log"
)

func GetRowSelect(res *sql.Rows) ([]entities.StorageInforTbl, error) {
	var (
		inforReturn   entities.StorageInforTbl
		lsInforReturn []entities.StorageInforTbl
	)
	for res.Next() {
		var infor entities.StorageInforTbl
		err := res.Scan(&infor.Date, &infor.Type, &infor.LineID, &infor.HashCode)
		if err != nil {
			log.Println("Error at GetRowSelect of database/query.go", err)
			return lsInforReturn, err
		}
		inforReturn = entities.NewInfor(infor.Date, infor.Type, infor.LineID, infor.HashCode)
		lsInforReturn = append(lsInforReturn, inforReturn)
	}
	return lsInforReturn, nil
}

func SelectByDateAndInfor(date string, hashCode string, db *sql.DB) ([]entities.StorageInforTbl, error) {
	var (
		dbName        = config.GetConfig().DB_NAME
		lsInforReturn []entities.StorageInforTbl
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
		return lsInforReturn, err
	}

	if lsInforReturn, err = GetRowSelect(res); err != nil {
		log.Println("Error at SelectByDateAndInfor of database/query.go when getRowSelect", err)
		return lsInforReturn, err
	}
	return lsInforReturn, nil
}

func SelectByDateDB(date string) ([]entities.StorageInforTbl, error) {
	var (
		dbName        = config.GetConfig().DB_NAME
		lsInforReturn []entities.StorageInforTbl
	)

	db, err := ConnectToDatabase(dbName)
	if err != nil {
		log.Println("Error ConnectToDatabase at SelectByDate of api/api.go", err)
	}

	query := `
		SELECT *
		FROM ` + dbName + `.StorageInfor st
		WHERE st.Date = '` + date + `'
	`
	//log.Println(query)
	res, err := db.Query(query)
	if err != nil {
		log.Println("Error at SelectByDate of database/query.go when select", err)
		return lsInforReturn, err
	}

	if lsInforReturn, err = GetRowSelect(res); err != nil {
		log.Println("Error at SelectByDate of database/query.go when getRowSelect", err)
		return lsInforReturn, err
	}
	return lsInforReturn, nil
}

func GetNumberADayDB(date string) (int, error) {
	var (
		count  int
		dbName = config.GetConfig().DB_NAME
	)
	db, err := ConnectToDatabase(dbName)
	if err != nil {
		log.Println("Error at ConnectToDatabase of GetNumberADay of api/api.go", err)
	}
	query := `
		SELECT count(*) as num
		FROM ` + dbName + `.StorageInfor st
		WHERE st.Date = '` + date + `'
	`

	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Error at GetNumberADayDB of database/query.go", err)
		return count, err
	}
	return count, nil
}

func SelectByHashCodeDB(hashCode string) ([]string, error) {
	var (
		dbName       = config.GetConfig().DB_NAME
		lsDateReturn []string
	)
	db, err := ConnectToDatabase(dbName)
	if err != nil {
		log.Println("Error ConnectToDatabase at SelectByHashCodeDB of api/api.go", err)
	}
	query := `
		SELECT Date
		FROM ` + dbName + `.StorageInfor st
		WHERE st.HashCode= '` + hashCode + `'
	`
	res, err := db.Query(query)
	if err != nil {
		log.Println("Error at SelectByHashCodeDB of database/query.go when select", err)
		return lsDateReturn, err
	}

	for res.Next() {
		var date string
		err := res.Scan(&date)
		if err != nil {
			log.Println("Error at GetRowSelect of database/query.go", err)
		}
		lsDateReturn = append(lsDateReturn, date)
	}
	return lsDateReturn, nil
}
