package api 

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "example.com/server/global"

    _ "github.com/lib/pq"
)

func GetSearchResults(response http.ResponseWriter, request *http.Request) {
    fmt.Println("GetSearchResults called")

    queryValues := request.URL.Query()
    searchQuery := queryValues.Get("query") // Assuming the query parameter is named 'query'
    _ = searchQuery

    client := global.GetClient()

    rows, err := client.Query("SELECT * FROM songs")
    defer rows.Close()

    if err != nil {
        log.Fatal(err)
    }

    var songs []global.Song

    for rows.Next() {
        var song global.Song 
        err := rows.Scan(&song.Id, &song.Title, &song.Artist, &song.Album, &song.ReleaseYear, &song.Genre, &song.DurationSeconds) 
        
        if err != nil { 
            log.Fatal(err) 
        }

        songs = append(songs, song)
        // if song.Title == searchQuery || song.Artist == searchQuery || song.Album == searchQuery || song.Genre == searchQuery {
        //     songs = append(songs, song)
        // }
    }

    json.NewEncoder(response).Encode(songs)
}

