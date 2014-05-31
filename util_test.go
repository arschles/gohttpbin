package main

import (
  "testing"
  . "github.com/franela/goblin"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "net/url"
  "fmt"
)

func TestOutput(t *testing.T) {
  g := Goblin(t)
  g.Describe("output", func() {
    g.It("should marshal the entire map", func() {
      respWriter := httptest.NewRecorder()
      m := map[string]interface{} {"a": "b"}
      output(respWriter, m)
      expectedEncoded, err := json.MarshalIndent(m, "", "  ")
      g.Assert(err).Equal(nil)
      g.Assert(expectedEncoded).Equal(respWriter.Body.Bytes())
    })
  })
}

func TestGetBase(t *testing.T) {
  g := Goblin(t)
  g.Describe("getBase", func() {
    g.It("should correctly assemble all basic output into a map", func() {
      headers := http.Header{
        "a":[]string{"b", "c"},
        "b":[]string{"c"},
      }
      rawQuery := "a=b&c=d"
      url := url.URL{RawQuery:rawQuery}
      request := http.Request {
        Header:headers,
        URL:&url,
        Host:"testHost",
        RequestURI:"/abc",
      }
      resultMap := getBase(&request)
      g.Assert(resultMap[HeadersKey]).Equal(parseHeaders(&request))
      g.Assert(resultMap[ArgsKey]).Equal(parseArgs(&request))
      g.Assert(resultMap[UrlKey]).Equal(fmt.Sprintf("%s%v", request.Host, request.RequestURI))
      g.Assert(resultMap[OriginKey]).Equal(request.Host)
    })
  })
}

func TestParseHeaders(t *testing.T) {
  g := Goblin(t)
  g.Describe("parseHeaders", func() {
      g.It("should parse the headers and output a Headers type", func() {
        inputHeaders := http.Header{
          "a": []string{"b", "c"},
          "b":[]string{"c"},
        }
        req := http.Request{Header:inputHeaders}
        resultHeaders := parseHeaders(&req)
        g.Assert(resultHeaders["a"]).Equal("b")
        g.Assert(resultHeaders["b"]).Equal("c")
      })
  })
}
