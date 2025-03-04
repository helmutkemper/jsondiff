package diffServer

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	server    *httptest.Server
	header    map[string]string
	response  string
	url       string
	errorFunc func(error)
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
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New("TestServer().SetResponse().json.Unmarshal().error"), err))
		}
		return
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
		_, err := w.Write([]byte(e.response))
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New("TestServer().Init().w.Write().error"), err))
		}
		return
	}))
}
