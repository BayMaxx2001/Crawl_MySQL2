
# Task 3
## Problem
1. Request https://malshare.com/daily
2. Use goroutines to speed up performance
3. Save all information in this link to the database
4. Write API 
## Getting Started
 1. Install MySQL and optional install MySQL WorkBench to easy use. 
 2. Create file `dev_config.json` in folder `config`  
 3. Configure features in file `config/dev_config.json` the following format: 
	 ``` json
	 { 
		 "DB_USERNAME": <YOUR_USERNAME>,
		 "DB_PASSWORD": <YOUR_PASSWORD>,
		 "DB_PORT": "3306",
		 "DB_HOST": "127.0.0.1",
		 "DB_NAME": "malshareDB"
	 }
	 ```
 4. We must install package: `tkanos/gonfig`,  `PuerkitoBio/goquery`, `go-sql-driver/mysql` by the following:
``` console
go get github.com/PuerkitoBio/goquery
go get github.com/go-sql-driver/mysql
go get github.com/tkanos/gonfig
```
 5. Before running the program, you must set up the database by running the command `go run setupDB`. 
 6. Use command:  `go run crawlData.go` download data to the database
 7. API :
	* Statistics of the number of samples in 1 day: 
	```url
		localhost:8080/get-number-infor-day/?date=<yyyy/mm/dd>
	```
	* Search for 1 hash:
	```url
		localhost:8080/get-date/?hashcode=<hashcode>
	```
	* Hash list in day:
	```url
		localhost:8080/get-list-infor-day/?date=<yyyy/mm/dd>
	```
## Steps
1. Request to https://malshare.com/daily.
2. Extract HTML response and get date data by get element table -> tr -> td.
3. Classify data is the md5, sha1,sha256 save to map.
4. Design goroutines to run concurrently
5. Save to database 
6. Write API