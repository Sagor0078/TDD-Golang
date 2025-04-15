package Accteptance_Test

import (
    "encoding/json"
    "net/http"
    "sync"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var (
    users = []User{}
    mu    sync.Mutex
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    mu.Lock()
    defer mu.Unlock()
    user.ID = len(users) + 1
    users = append(users, user)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    json.NewEncoder(w).Encode(users)
}

func setupRoutes() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            createUserHandler(w, r)
        } else if r.Method == http.MethodGet {
            getUsersHandler(w, r)
        }
    })
    return mux
}
