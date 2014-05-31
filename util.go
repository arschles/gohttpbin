package gohttpbin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Headers map[string]string
type Args map[string][]string

const (
	ArgsKey    = "args"
	HeadersKey = "headers"
	UrlKey     = "url"
	OriginKey  = "origin"
	DataKey    = "data"
)

func output(resp http.ResponseWriter, m map[string]interface{}) {
	marshalled, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Write(marshalled)
}

func getBase(req *http.Request) map[string]interface{} {
	m := map[string]interface{}{}
	m[HeadersKey] = parseHeaders(req)
	m[ArgsKey] = parseArgs(req)
	m[UrlKey] = fmt.Sprintf("%s%v", req.Host, req.RequestURI)
	m[OriginKey] = req.Host
	return m
}

func parseHeaders(req *http.Request) Headers {
	hdrs := req.Header
	ret := Headers{}
	for name, vals := range hdrs {
		if len(vals) > 0 {
			ret[name] = vals[0]
		}
	}
	return ret
}

func parseArgs(req *http.Request) Args {
	return Args(req.URL.Query())
}

func readBody(req *http.Request) []byte {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return []byte{}
	}
	return body
}
