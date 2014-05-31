package main

import "net/http"

type Headers map[string]string
type Args map[string][]string

func parseHeaders(req *http.Request) Headers {
  hdrs := req.Header
  ret := Headers{}
  for name, vals := range(hdrs) {
    if len(vals) > 0 {
      ret[name] = vals[0]
    }
  }
  return ret
}

func parseArgs(req *http.Request) Args {
  return Args(req.URL.Query())
}
