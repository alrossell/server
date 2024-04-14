package api 

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "path"

    "example.com/server/global"

    _ "github.com/lib/pq"
)

func PostUser(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Creating a new user")
    client := global.GetClient()

    var user global.User
    _ = json.NewDecoder(request.Body).Decode(&user)

    _, err := client.Exec(
        `INSERT INTO users (first_name, last_name, email, creation_date, password) 
         VALUES ($1, $2, $3, $4, $5)`,
        user.FirstName, user.LastName, user.Email, user.Date, user.Password)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("User created successfully!")
}

func GetUsers(response http.ResponseWriter, request *http.Request) {
    fmt.Println("GetUsers called")

    client := global.GetClient()

    rows, err := client.Query("SELECT * FROM users")
    defer rows.Close()

    if err != nil {
        log.Fatal(err)
    }

    var users []global.User

    for rows.Next() {
        var user global.User 
        err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date, &user.Password) 
        if err != nil { 
            log.Fatal(err) 
        }
        users = append(users, user)
    }

    json.NewEncoder(response).Encode(users)

    fmt.Println("Users retrieved successfully!")
}

func GetUser(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Get User")

    client := global.GetClient()

    id := path.Base(request.URL.Path)

    rows, err := client.Query("SELECT * FROM users WHERE email = $1", id)
    defer rows.Close()

    if err != nil {
        log.Fatal(err)
    }

    var user global.User

    for rows.Next() {
        err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date, &user.Password) 
        if err != nil { 
            log.Fatal(err) 
        }
    }

    json.NewEncoder(response).Encode(user)

    fmt.Println("User retrieved successfully!")
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
}

func DeleteAllUsers(response http.ResponseWriter, request *http.Request) {

}
