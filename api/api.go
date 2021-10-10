package api

import (
	"crawl_data/database"
	"encoding/json"
	"log"
	"net/http"
)

func getDateInParam(w http.ResponseWriter, r *http.Request) string {
	date, check := r.URL.Query()["date"]

	if !check || len(date[0]) < 1 {
		log.Println("Url Param 'date' is missing")
		json.NewEncoder(w).Encode("Url Param 'date' is missing")
		return ""
	}
	log.Println(date[0])
	return date[0]
}
func getHashCodeInParam(w http.ResponseWriter, r *http.Request) string {
	hashCode, check := r.URL.Query()["hashcode"]

	if !check || len(hashCode[0]) < 1 {
		log.Println("Url Param 'hashCode' is missing")
		json.NewEncoder(w).Encode("Url Param 'hashCode' is missing")
		return ""
	}
	log.Println(hashCode[0])
	return hashCode[0]
}

func GetNumberInforADayAPI(w http.ResponseWriter, r *http.Request) {
	date := getDateInParam(w, r)
	number, err := database.GetNumberADayDB(date)
	if err != nil {
		log.Println("Error at GetNumberInforADayAPI of api/api.go", err)
		json.NewEncoder(w).Encode(nil)
	}
	json.NewEncoder(w).Encode(number)
}

func SelectByDateAPI(w http.ResponseWriter, r *http.Request) {
	date := getDateInParam(w, r)
	listHashCode, err := database.SelectByDateDB(date)
	if err != nil {
		log.Println("Error at SelectByDateAPI of api/api.go", err)
		json.NewEncoder(w).Encode(nil)
	}
	json.NewEncoder(w).Encode(listHashCode)
}
func SelectByHashCodeAPI(w http.ResponseWriter, r *http.Request) {
	hashCode := getHashCodeInParam(w, r)
	lsDate, err := database.SelectByHashCodeDB(hashCode)
	if err != nil {
		log.Println("Error at SelectByHashCodeAPI of api/api.go", err)
		json.NewEncoder(w).Encode(nil)
	}
	json.NewEncoder(w).Encode(lsDate)
}
