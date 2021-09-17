
# Task 3
## Problem
1. Request https://malshare.com/daily
2. Use goroutines to speed up performance
3. Save all information in this link to the database

## Getting Started
 1. Install MySQL and MySQL WorkBench
 2. Configure features in file `config/sample_config.json` the following format: 
	 ``` json
	 { 
		 "DB_USERNAME": <YOUR_USERNAME>,
		 "DB_PASSWORD": <YOUR_PASSWORD>,
		 "DB_PORT": "3306",
		 "DB_HOST": "127.0.0.1",
		 "DB_NAME": "malshareDB"
	 }
	 ```
 3. We must install package: `tkanos/gonfig`,  `PuerkitoBio/goquery`, `go-sql-driver/mysql` by the following:
``` console
go get github.com/PuerkitoBio/goquery
go get github.com/go-sql-driver/mysql
go get github.com/tkanos/gonfig
```
 4. Use command:  ``` console 
 		go run main.go 
		```  
		download data to the database

## Steps
1. Request to https://malshare.com/daily.
2. Extract HTML response and get date data by get element table -> tr -> td.
3. Classify data is the md5, sha1,sha256 save to map.
4. Design goroutines to run concurrently
5. Save to database 
