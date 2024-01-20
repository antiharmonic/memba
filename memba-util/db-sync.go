package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	_ "github.com/lib/pq"

	// "math/rand"
	"os"

	"github.com/antiharmonic/memba/memba-server/memba"
)

var db *sql.DB

type StreamItem struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Url     string `json:"url"`
}

func main() {
	// configuration/setup
	log.SetOutput(os.Stdout)
	var config memba.Config
	err := memba.LoadConfig(&config)
	if err != nil {
		log.Fatalln(err)
	}
	// validate db connection or bail
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Database.User,
		url.QueryEscape(config.Database.Password),
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	var items []StreamItem
	// get external information or bail
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(config.Web.URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		log.Fatalln(err)
	}

	// debug only
	// fmt.Println(items[rand.Intn(len(items))].Title)
	// upsert existing "stream" data cache
	// i don't think this will be the final way to do this - i think ultimately we'll upsert into the mixed memba_media table or something as long as memba_media.stream_item = true or something like that
	sql := `
    insert into stream_cache (id, title, comment, url) values ($1, $2, $3, $4)
    on conflict (id) do update set title = $2, comment = $3, url = $4
  `
	sth, err := db.Prepare(sql)
	if err != nil {
		log.Fatalln(err)
	}
	defer sth.Close()

	for _, item := range items {
		_, err = sth.Exec(item.Id, item.Title, item.Comment, item.Url)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("Stream data cached")
}
