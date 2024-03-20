package global

import (
    "log"
    "database/sql"

    _ "github.com/lib/pq"
)


type Song struct {
    Id int
    Title string 
    Artist string
    Album string
    ReleaseYear int
    Genre string
    DurationSeconds int
}

const DataBase = "myDatabase"
const Collection = "myCollection"

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
