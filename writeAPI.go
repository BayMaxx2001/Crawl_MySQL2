package main

import (
	"crawl_data/api"
	"log"
	"net/http"
)

func writeAPI() {
	http.HandleFunc("/stats-day/", api.GetStatisticsADayAPI)
	http.HandleFunc("/get-date/", api.SelectByHashCodeAPI)
	http.HandleFunc("/get-list-infor-day/", api.SelectByDateAPI)
}

func main() {
	writeAPI()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
