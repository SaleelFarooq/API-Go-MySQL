package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "mux-master"
)
import "time"


type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}


type Todos []Todo

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    json.NewEncoder(w).Encode(todos)
}
func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}



func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
