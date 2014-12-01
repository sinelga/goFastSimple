package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"strconv"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var hitsFlag = flag.Int("hits", 0, "if not will be 0")

var keywordsarr []domains.Keyword

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	flag.Parse() // Scan the arguments list

	locale := *localeFlag
	themes := *themesFlag
	hitsint := *hitsFlag

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgiLight/goFastCgiLight/singo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hits := strconv.Itoa(hitsint)
	sqlstr := "select keyword,hits from keywords where locale='" + locale + "' and themes='" + themes + "' and hits>=" + hits + " and Cl=0"

	golog.Info(sqlstr)

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var keyword string
		var hits int
		rows.Scan(&keyword, &hits)
		keywordObj := domains.Keyword{Keyword: keyword, Hits: hits}
		keywordsarr = append(keywordsarr, keywordObj)

	}
	rows.Close()
	db.Close()

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	queuename := locale + ":" + themes + ":keywords"

	for _, keywordObj := range keywordsarr {

		if r, err := c.Do("ZADD", queuename, strconv.Itoa(keywordObj.Hits), keywordObj.Keyword); err != nil {
			golog.Err("syncpushdomains: " + err.Error())

		} else {

			//				golog.Info(r)
			fmt.Println("r ", r)

		}

	}

}
