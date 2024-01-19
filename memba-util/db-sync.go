package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gookit/config/v2"
  "github.com/gookit/config/v2/yaml"
  "database/sql"
  _ "github.com/lib/pq"
  "net/url"
  "log"
  "time"
)

var db *sql.DB

type StreamItem struct {
  Id      int     `json:"id"`
  Title   string  `json:"title"`
  Comment string  `json:"comment"`
  Url     string  `json:"url"`
}

func main() {
  // configuration/setup
  config.AddDriver(yaml.Driver)
  err := config.LoadFiles("config.yml")
  if err != nil {
    log.Fatal(err)
  }
  // validate db connection or bail
  connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
    config.String("database.user"),
    url.QueryEscape(config.String("database.pass")),
    config.String("database.host"),
    config.String("database.port"),
    config.String("database.db_name"),
  )
  db, err = sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  // get external information or bail
  client := &http.Client{Timeout: 10 * time.Second}
  res, err := client.Get(config.String("stream.url"))
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  err = json.NewDecoder(r.Body).Decode(&items)
  if err != nil {
    log.Fatal(err)
  }

  // debug only
  fmt.Println(items[0].Title)
  // upsert existing "stream" data cache
  // i don't think this will be the final way to do this - i think ultimately we'll upsert into the mixed memba_media table or something as long as memba_media.stream_item = true or something like that
  

}
