package dal

import (
	"crawl_data/config"
	"crawl_data/database/entities"
	"crawl_data/helpersCrawl/utils"
	"database/sql"
	"log"
)

type DateAndType struct {
	Date string
	Type string
}

func GetRowSelect(res *sql.Rows) ([]entities.StorageInforTbl, error) {
	var (
		inforReturn   entities.StorageInforTbl
		lsInforReturn []entities.StorageInforTbl
	)
	for res.Next() {
		var infor entities.StorageInforTbl
		err := res.Scan(&infor.Date, &infor.Type, &infor.LineID, &infor.HashCode)
		if err != nil {
			log.Println("Error at GetRowSelect of database/dal/query.go", err)
			return lsInforReturn, err
		}
		inforReturn = entities.NewInfor(infor.Date, infor.Type, infor.LineID, infor.HashCode)
		lsInforReturn = append(lsInforReturn, inforReturn)
	}
	return lsInforReturn, nil
}

func SelectByDateAndInfor(date string, hashCode string, db *sql.DB) ([]entities.StorageInforTbl, error) {
	var (
		lsInforReturn []entities.StorageInforTbl
	)
	year, month, day := utils.GetDateDetail(date)
	date = year + "-" + month + "-" + day

	query := `
		SELECT *
		FROM StorageInfor st
		WHERE st.Date = ? AND st.HashCode= ?
	`
	//log.Println(query)
	res, err := db.Query(query, date, hashCode)
	if err != nil {
		log.Println("Error at SelectByDateAndInfor of database/dal/query.go when select", err)
		return lsInforReturn, err
	}

	if lsInforReturn, err = GetRowSelect(res); err != nil {
		log.Println("Error at SelectByDateAndInfor of database/dal/query.go when getRowSelect", err)
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
		log.Println("Error ConnectToDatabase at SelectByDate of database/dal/query.go", err)
	}

	query := `
		SELECT *
		FROM StorageInfor st
		WHERE st.Date = ?
	`
	//log.Println(query)
	res, err := db.Query(query, date)
	if err != nil {
		log.Println("Error at SelectByDate of database/dal/query.go when select", err)
		return lsInforReturn, err
	}

	if lsInforReturn, err = GetRowSelect(res); err != nil {
		log.Println("Error at SelectByDate of database/dal/query.go when getRowSelect", err)
		return lsInforReturn, err
	}
	return lsInforReturn, nil
}

func GetStatisticsADayDB(date string) (int, error) {
	var (
		count  int
		dbName = config.GetConfig().DB_NAME
	)
	db, err := ConnectToDatabase(dbName)
	if err != nil {
		log.Println("Error at ConnectToDatabase of GetNumberADay of database/dal/query.go", err)
	}
	query := `
		SELECT count(*) as num
		FROM StorageInfor st
		WHERE st.Date = ?
	`

	err = db.QueryRow(query, date).Scan(&count)
	if err != nil {
		log.Println("Error at GetStatisticsADayDB of database/dal/query.go", err)
		return count, err
	}
	return count, nil
}

func SelectByHashCodeDB(hashCode string) ([]DateAndType, error) {
	var (
		dbName        = config.GetConfig().DB_NAME
		lsDateAndType []DateAndType
	)

	db, err := ConnectToDatabase(dbName)
	if err != nil {
		log.Println("Error ConnectToDatabase at SelectByHashCodeDB of database/dal/query.go", err)
	}
	query := `
		SELECT Date , Type
		FROM StorageInfor st
		WHERE st.HashCode= ?
	`
	res, err := db.Query(query, hashCode)
	if err != nil {
		log.Println("Error at SelectByHashCodeDB of database/dal/query.go when select", err)
		return lsDateAndType, err
	}

	for res.Next() {
		var (
			date       string
			typeOfHash string
		)
		err := res.Scan(&date, &typeOfHash)
		if err != nil {
			log.Println("Error at GetRowSelect of database/dal/query.go", err)
		}
		lsDateAndType = append(lsDateAndType, DateAndType{date, typeOfHash})
	}
	return lsDateAndType, nil
}

func ExecQueryDatabase(query string, db *sql.DB) error {
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error at ExecQueryDatabase of database/dal/database.go ", err)
		return err
	}
	return err
}

func InsertStorageInforTbl(infor entities.StorageInforTbl, db *sql.DB) error {
	// prepare
	query := `
		INSERT INTO StorageInfor(Date, Type, LineID, HashCode) 
		VALUES (?, ?, ?, ?)
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error at InsertStorageInforTbl of database/dal/database.go when prepare statement ", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(infor.Date, infor.Type, infor.LineID, infor.HashCode); err != nil {
		log.Println("Error at InsertStorageInforTbl of database/dal/database.go when execute", err)
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
