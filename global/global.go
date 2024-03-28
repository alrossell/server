package global

import (
    "log"
    "database/sql"

    _ "github.com/lib/pq"
)

type Review struct {
    Id int `json:"id"`
    UserId int `json:"user_id"` 
    SongId int `json:"song_id"`
    Date string `json:"date"`
    Review string `json:"review"`
}

type Song struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Album string `json:"album"`
    ReleaseYear int `json:"release_year"`
    Genre string `json:"genre"`
    DurationSeconds int `json:"duration_seconds"`
}

type Test struct {
    id, releaseYear, durationSeconds int
    title, artist, album, genre string
}

var client *sql.DB

func CreateClient() {
    connStr := "user=postgres password=mypasswpord dbname=new_database sslmode=disable"
    
    var err error
    client, err = sql.Open("postgres", connStr)

    if err != nil {
		log.Fatal(err)
	}
}

func GetClient() *sql.DB {
    if client == nil {
        CreateClient()
    }

    return client
}

func CloseClient() {
    if client != nil {
        client.Close()
    } else {
        log.Fatal("Client is nil")
    }
}
