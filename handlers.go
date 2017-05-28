package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	hello := "hello"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hello); err != nil {
		panic(err)
	}
}

func ShowPass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passId := vars["passId"]
	fmt.Fprintln(w, "Todo show:", passId)
}

func SetPass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	var pass Pass
	if err := decoder.Decode(&pass); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer r.Body.Close()
	pass.Id = uuid.New()
	SetKey(pass)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pass.Id)
}
