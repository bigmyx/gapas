package main

import (
	"fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    pass := Pass{Pass: "121212"}
    if err := json.NewEncoder(w).Encode(pass); err != nil {
    	panic(err)
    }
}

func ShowPass(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    passId := vars["passId"]
    fmt.Fprintln(w, "Todo show:", passId)
}