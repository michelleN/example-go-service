package main

import (
	"encoding/json"
	"net/http"
)

type Greeting struct {
	Message string
	Color   string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	greeting := Greeting{"Hello", "ffbf00"}

	js, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
