package diffServer

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HttpRequest Faz o request http dos dados
type HttpRequest struct {
	header  []keyValue
	param   []keyValue
	url     string
	method  string
	content string

	event func([]string)
}

// SetEvent Ponteiro para função externa, chamada quando um evento acontece
func (e *HttpRequest) SetEvent(event func([]string)) {
	e.event = event
}

// AddHeader Adiciona um header a chamada http
func (e *HttpRequest) AddHeader(key, value string) {
	e.header = append(e.header, keyValue{key, value})
}

// AddParam Adiciona um parâmetro a chamada http
func (e *HttpRequest) AddParam(key, value string) {
	e.param = append(e.param, keyValue{key, value})
}

// SetMethod Define o método da chamada http, GET, POST, etc.
func (e *HttpRequest) SetMethod(method string) {
	e.method = method
}

// SetContent Define o conteúdo da chamada http
func (e *HttpRequest) SetContent(content string) {
	e.content = content
}

// SetUrl Define a URL da chamada
func (e *HttpRequest) SetUrl(url string) {
	e.url = url
}

// BasicAuth Define a autenticação básica
func (e *HttpRequest) BasicAuth(username, password string) string {
	// todo: fazer
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Request Executa a requisição da chamada http
func (e *HttpRequest) Request() (bodyResponse string) {
	var err error
	var body []byte
	var request *http.Request
	var response *http.Response

	urlStr := e.url
	for key, param := range e.param {
		if key == 0 {
			urlStr += "?"
		} else {
			urlStr += "&"
		}

		urlStr += url.QueryEscape(param.Key + "=" + param.Value)
	}

	request, err = http.NewRequest(e.method, urlStr, strings.NewReader(e.content))
	if err != nil {
		panic(err)
	}

	for _, header := range e.header {
		request.Header.Add(header.Key, header.Value)
	}

	response, err = http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if e.event != nil {
		authorization := response.Header.Get("Authorization")
		e.event([]string{authorization})
	}

	return string(body)
}
