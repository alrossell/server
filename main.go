package main

import (
    "fmt"
    "net/http"

    "example.com/server/global"
    "example.com/server/api"

    "github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // Allow all origins
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Allowed methods
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization") // Allowed headers

         w.Header().Set("Access-Control-Allow-Credentials", "true")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
    fmt.Println("Starting Server") 

    global.CreateClient()
    defer global.CloseClient()
    fmt.Println("Database connection created")

    router := mux.NewRouter()

    apiRouter := router.PathPrefix("/").Subrouter()
    apiRouter.Use(corsMiddleware)

    apiRouter.HandleFunc("/songs", api.CreateSong).Methods("POST", "OPTIONS") // Include OPTIONS to handle preflight requests
    apiRouter.HandleFunc("/songs", api.GetSongs).Methods("GET", "OPTIONS")    // Include OPTIONS to handle preflight requests
    apiRouter.HandleFunc("/songs/{id}", api.DeleteSong).Methods("DELETE", "OPTIONS") // Include OPTIONS to handle preflight requests
    apiRouter.HandleFunc("/songs", api.DeleteAllSongs).Methods("DELETE", "OPTIONS") // Include OPTIONS to handle preflight requests
    apiRouter.HandleFunc("/songs", api.DeleteAllSongs).Methods("DELETE", "OPTIONS") // Include OPTIONS to handle preflight requests
   
    apiRouter.HandleFunc("/songs", api.GetSearchResults).Methods("GET", "OPTIONS")    // Include OPTIONS to handle preflight requests

    http.ListenAndServe(":5000", router)
}
