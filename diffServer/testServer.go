package diffServer

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	server   *httptest.Server
	header   map[string]string
	response string
	url      string
}

func (e *TestServer) GetUrl() (url string) {
	if e.server != nil {
		e.Init()
	}

	return e.server.URL
}

func (e *TestServer) SetResponse(response any) {
	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err)
	}
	e.response = string(data)
}

func (e *TestServer) AddHeader(key, value string) {
	if e.header == nil {
		e.header = make(map[string]string)
	}

	e.header[key] = value
}

func (e *TestServer) Init() {
	e.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(5 * time.Second) // todo: remover
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		for k, v := range e.header {
			w.Header().Set(k, v)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(e.response))
	}))
}
