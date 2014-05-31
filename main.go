package gohttpbin

import (
	"github.com/gorilla/mux"
	"net/http"
)

func get(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}

func post(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	m[DataKey] = string(readBody(req))
	output(resp, m)
}

func put(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	m[DataKey] = string(readBody(req))
	output(resp, m)
}

func delete(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}

func head(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}

func patch(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/get", get).Methods("GET")
	router.HandleFunc("/post", post).Methods("POST")
	router.HandleFunc("/put", put).Methods("PUT")
	router.HandleFunc("/delete", delete).Methods("DELETE")
	router.HandleFunc("/head", head).Methods("HEAD")
	router.HandleFunc("/patch", patch).Methods("PATCH")
	http.Handle("/", router)
}
