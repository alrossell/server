package api 

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "example.com/server/global"

    _ "github.com/lib/pq"
)


func GetReviews(response http.ResponseWriter, request *http.Request) {
    fmt.Println("GetReviews called")

    client := global.GetClient()

    rows, err := client.Query("SELECT * FROM reviews")
    defer rows.Close()

    if err != nil {
        log.Fatal(err)
    }

    var reviews []global.Review

    for rows.Next() {
        var review global.Review 
        err := rows.Scan(&review.Id, &review.UserId, &review.SongId, &review.Date, &review.Review) 
        if err != nil { 
            log.Fatal(err) 
        }
        reviews = append(reviews, review)
    }

    json.NewEncoder(response).Encode(reviews)

    fmt.Println("Songs retrieved successfully!")
}

func GetReview(response http.ResponseWriter, request *http.Request) {

}

func PostReview(response http.ResponseWriter, request *http.Request) {
    fmt.Println("PostReview called")

    queryValues := request.URL.Query()
    searchQuery := queryValues.Get("query") // Assuming the query parameter is named 'query'
    _ = searchQuery

    client := global.GetClient()

    var review global.Review
    _ = json.NewDecoder(request.Body).Decode(&review)

    fmt.Println(review)

    _, err := client.Exec(
        `INSERT INTO reviews (user_id, song_id, date, review) 
         VALUES ($1, $2, $3, $4)`,
        review.UserId, review.SongId, review.Date, review.Review)

    if err != nil {
        log.Fatal(err)
    }

    json.NewEncoder(response).Encode(review)
}

