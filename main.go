package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "flag"
  "fmt"
  "log"
  "encoding/json"
)

var port = flag.Int("port", 8080, "port to bind")

const (
  ArgsKey = "args"
  HeadersKey = "headers"
  UrlKey = "url"
  OriginKey = "origin"
)

func get(resp http.ResponseWriter, req *http.Request) {
  m := map[string]interface{}{}
  m[HeadersKey] = parseHeaders(req)
  m[ArgsKey] = parseArgs(req)
  m[UrlKey] = fmt.Sprintf("%s%v", req.Host, req.RequestURI)
  m[OriginKey] = req.Host

  marshalled, err := json.MarshalIndent(m, "", "  ")
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
    return
  }
  resp.Write(marshalled)
}

func post(resp http.ResponseWriter, req *http.Request) {

}

func put(resp http.ResponseWriter, req *http.Request) {
}

func delete(resp http.ResponseWriter, req *http.Request) {

}

func head(resp http.ResponseWriter, req *http.Request) {

}

func patch(resp http.ResponseWriter, req *http.Request) {

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
