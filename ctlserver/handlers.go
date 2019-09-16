package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func GetCurrent(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(RepoGetCurrentState()); err != nil {
		panic(err)
	}
}

func SetCurrent(w http.ResponseWriter, r *http.Request) {
	var state State
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &state); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	s := RepoCreateState(state)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}
