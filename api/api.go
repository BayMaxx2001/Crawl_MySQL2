package api

import (
	"crawl_data/database"
	"crawl_data/database/entities"
	"encoding/json"
	"log"
	"net/http"
)

func getDateInParam(w http.ResponseWriter, r *http.Request) string {
	date, ok := r.URL.Query()["date"]

	if !ok || len(date[0]) < 1 {
		log.Println("Url Param 'date' is missing")
		json.NewEncoder(w).Encode("Url Param 'date' is missing")
		return ""
	}
	log.Println(date[0])
	return date[0]
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
func TestAPI(w http.ResponseWriter, r *http.Request) {
	a := entities.NewInfor("2020", "md5", 1, "123")
	json.NewEncoder(w).Encode(a)
}
