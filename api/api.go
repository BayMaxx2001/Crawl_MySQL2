package api

import (
	"crawl_data/database/dal"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func getDateInParam(w http.ResponseWriter, r *http.Request) (string, error) {
	date, check := r.URL.Query()["date"]

	if !check || len(date[0]) < 1 {
		log.Println("Url Param 'date' is missing")
		return "", errors.New("url param 'date' is missing")
	}
	log.Println(date[0])
	return date[0], nil
}
func getHashCodeInParam(w http.ResponseWriter, r *http.Request) (string, error) {
	hashCode, check := r.URL.Query()["hashcode"]

	if !check || len(hashCode[0]) < 1 {
		log.Println("Url Param 'hashCode' is missing")
		return "", errors.New("url param 'hashCode' is missing")
	}
	log.Println(hashCode[0])
	return hashCode[0], nil
}

func GetStatisticsADayAPI(w http.ResponseWriter, r *http.Request) {
	date, err := getDateInParam(w, r)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	number, err := dal.GetStatisticsADayDB(date)
	if err != nil {
		log.Println("Error at GetStatisticsADayAPI of api/api.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	json.NewEncoder(w).Encode(number)
}

func SelectByDateAPI(w http.ResponseWriter, r *http.Request) {
	date, err := getDateInParam(w, r)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	listHashCode, err := dal.SelectByDateDB(date)
	if err != nil {
		log.Println("Error at SelectByDateAPI of api/api.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	json.NewEncoder(w).Encode(listHashCode)
}
func SelectByHashCodeAPI(w http.ResponseWriter, r *http.Request) {
	hashCode, err := getHashCodeInParam(w, r)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	lsDate, err := dal.SelectByHashCodeDB(hashCode)
	if err != nil {
		log.Println("Error at SelectByHashCodeAPI of api/api.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	json.NewEncoder(w).Encode(lsDate)
}
