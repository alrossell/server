package api 

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "example.com/server/global"

    _ "github.com/lib/pq"
)

func CreateSong(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Creating a new song")

    client := global.GetClient()

    var song global.Song
    _ = json.NewDecoder(request.Body).Decode(&song)

    _, err := client.Exec(
        `INSERT INTO songs (title, artist, album, release_year, genre, duration_seconds) 
         VALUES ($1, $2, $3, $4, $5, $6)`,
        song.Title, song.Artist, song.Album, song.ReleaseYear, song.Genre, song.DurationSeconds)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Song created successfully!")
}

func GetSongs(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Getting all songs")

    client := global.GetClient()

    rows, err := client.Query("SELECT song_id, title FROM songs")
    defer rows.Close()

    if err != nil {
        log.Fatal(err)
    }

    var songs []global.Song

    for rows.Next() {
        var song global.Song 
        err := rows.Scan(&song.Title, &song.Artist, &song.Album, &song.ReleaseYear, &song.Genre, &song.DurationSeconds) 
        if err != nil { 
            log.Fatal(err) 
        }
        songs = append(songs, song)
    }

    json.NewEncoder(response).Encode(songs)

    fmt.Println("Songs retrieved successfully!")
}

func DeleteSong(response http.ResponseWriter, request *http.Request) {
}

func DeleteAllSongs(response http.ResponseWriter, request *http.Request) {

}

