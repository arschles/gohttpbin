package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "flag"
  "fmt"
  "log"
)

var port = flag.Int("port", 8080, "port to bind")

func get(resp http.ResponseWriter, req *http.Request) {
  m := getBase(req)
  output(resp, m)
}

func post(resp http.ResponseWriter, req *http.Request) {
  m := getBase(req)
  m[BodyKey] = string(readBody(req))
  output(resp, m)
}

func put(resp http.ResponseWriter, req *http.Request) {
  m := getBase(req)
  m[BodyKey] = string(readBody(req))
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

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/get", get).Methods("GET")
  router.HandleFunc("/post", post).Methods("POST")
  router.HandleFunc("/put", put).Methods("PUT")
  router.HandleFunc("/delete", delete).Methods("DELETE")
  router.HandleFunc("/head", head).Methods("HEAD")
  router.HandleFunc("/patch", patch).Methods("PATCH")
  flag.Parse()

  log.Printf("starting gohttpbin on port %d", *port)
  http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
