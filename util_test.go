package gohttpbin

import (
	"encoding/json"
	"fmt"
	. "github.com/franela/goblin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOutput(t *testing.T) {
	g := Goblin(t)
	g.Describe("output", func() {
		g.It("should marshal the entire map", func() {
			respWriter := httptest.NewRecorder()
			m := map[string]interface{}{"a": "b"}
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
			req, err := http.NewRequest("GET", "http://test.com/abc", nil)
			g.Assert(err).Equal(nil)

			headers := http.Header{
				"a": []string{"b", "c"},
				"b": []string{"c"},
			}
			req.Header = headers
			resultMap := getBase(req)
			g.Assert(resultMap[HeadersKey]).Equal(parseHeaders(req))
			g.Assert(resultMap[ArgsKey]).Equal(parseArgs(req))
			g.Assert(resultMap[UrlKey]).Equal(fmt.Sprintf("%s%v", req.Host, req.RequestURI))
			g.Assert(resultMap[OriginKey]).Equal(req.Host)
		})
	})
}

func TestParseHeaders(t *testing.T) {
	g := Goblin(t)
	g.Describe("parseHeaders", func() {
		g.It("should parse the headers and output a Headers type", func() {
			inputHeaders := http.Header{
				"a": []string{"b", "c"},
				"b": []string{"c"},
			}
			req := http.Request{Header: inputHeaders}
			resultHeaders := parseHeaders(&req)
			g.Assert(resultHeaders["a"]).Equal("b")
			g.Assert(resultHeaders["b"]).Equal("c")
		})
	})
}

func TestParseArgs(t *testing.T) {
	g := Goblin(t)
	g.Describe("parseArgs", func() {
		g.It("should parse query string arguments and choose the first value if there are duplicates", func() {
			req, err := http.NewRequest("GET", "http://test.com?a=b&a=c", nil)
			g.Assert(err).Equal(nil)
			args := parseArgs(req)
			g.Assert(args["a"]).Equal("b")
		})
	})
}
