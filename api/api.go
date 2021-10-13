package api

import (
	"crawl_data/database"
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

func GetNumberInforADayAPI(w http.ResponseWriter, r *http.Request) {
	date, err := getDateInParam(w, r)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	number, err := database.GetNumberADayDB(date)
	if err != nil {
		log.Println("Error at GetNumberInforADayAPI of api/api.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	json.NewEncoder(w).Encode(number)
}

func SelectByDateAPI(w http.ResponseWriter, r *http.Request) {
	date, err := getDateInParam(w, r)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	listHashCode, err := database.SelectByDateDB(date)
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
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	lsDate, err := database.SelectByHashCodeDB(hashCode)
	if err != nil {
		log.Println("Error at SelectByHashCodeAPI of api/api.go", err)
		json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}
	json.NewEncoder(w).Encode(lsDate)
}
