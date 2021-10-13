package utils

import (
	"fmt"
	"log"
)

func GetListDate(url string) ([]string, error) {
	var (
		listDate []string
		err      error
	)
	log.Println("Start getting the list date")

	resp, err := RequestHTTP(url)
	if err != nil {
		log.Println("Error at GetListDate of utils/date.go when get list date ", err)
		return listDate, err
	}
	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Status code error: %d, %s", resp.StatusCode, resp.Status) + "at GetListDate of utils/date.go "
		log.Fatal(msg)
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := LoadHTML(resp)
	if err != nil {
		log.Println("Error at GetListDate of utils/date.go when load HTML document ", err)
		return listDate, err
	}
	// Find the day items
	sel_getTr := doc.Find("table tr")
	for i := range sel_getTr.Nodes {

		query := sel_getTr.Eq(i)
		sel_getTd := query.Find("td")

		for j := range sel_getTd.Nodes {

			day := sel_getTd.Eq(j)

			if j == 1 {
				add := string(day.Text())
				if IsDate(add) {
					listDate = append(listDate, string(add[:len(add)-1]))
				}
				//log.Println(string(day.Text()))
			}
		}
	}
	log.Println("Complete get the list date")
	return listDate, err
}
