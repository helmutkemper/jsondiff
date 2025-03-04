package diffServer

import (
	"encoding/json"
	"errors"
	"github.com/rivo/tview"
	"strconv"
)

type ConnectServerToken struct {
	header      []keyValue
	param       []keyValue
	url         string
	method      string
	content     string
	fieldWidth  int
	form        *tview.Form
	formConnect FormConnect
	formToken   formToken
	tokenServer testServerToken
	token       Token

	event     func(Token)
	errorFunc func(error)
}

// SetEvent Recebe o ponteiro de função para quando o token é recebido
//
//	Esta é a única maneira de passar um evento de dentro para fora de um struct
func (e *ConnectServerToken) SetEvent(event func(Token)) {
	e.event = event
}

func (e *ConnectServerToken) SetUrl(url string) {
	e.url = url
}

func (e *ConnectServerToken) SetMethod(method string) {
	e.method = method
}

func (e *ConnectServerToken) SetContent(content string) {
	e.content = content
}

func (e *ConnectServerToken) SetHeader(key, value string) {
	for k := range e.header {
		if e.header[k].Key == key {
			e.header[k].Value = value
			return
		}
	}
}

func (e *ConnectServerToken) SetParam(key, value string) {
	for k := range e.param {
		if e.param[k].Key == key {
			e.param[k].Value = value
			return
		}
	}
}

func (e *ConnectServerToken) GetHeader() []keyValue {
	return e.header
}

func (e *ConnectServerToken) GetParam() []keyValue {
	return e.param
}

func (e *ConnectServerToken) GetUrl() string {
	return e.url
}

func (e *ConnectServerToken) GetMethod() string {
	return e.method
}

func (e *ConnectServerToken) GetContent() string {
	return e.content
}

func (e *ConnectServerToken) GetFieldWidth() int {
	return e.fieldWidth
}

func (e *ConnectServerToken) GetFormToken() *tview.Form {
	return e.form
}

func (e *ConnectServerToken) GetTokenServer() testServerToken {
	return e.tokenServer
}

func (e *ConnectServerToken) Init(fieldWidth int) {
	e.fieldWidth = fieldWidth
	e.tokenServer.init()

	e.header = make([]keyValue, 0)
	//e.header = append(e.header, keyValue{"token", ""})
	//e.header = append(e.header, keyValue{"type", ""})

	e.param = make([]keyValue, 0)
	//e.param = append(e.param, keyValue{"user", ""})
	//e.param = append(e.param, keyValue{"password", ""})

	e.mountFormToken()
}

func (e *ConnectServerToken) GetToken() {

	// Monta o token a ser enviado pelo servidor de teste
	tk := new(Token)
	tk.Init()

	// Define a resposta do servidor de teste
	e.tokenServer.init()
	e.tokenServer.TestServer.Init()
	e.tokenServer.TestServer.SetResponse(&tk)

	// Prepara a requisição
	req := new(HttpRequest)
	req.SetMethod(e.method)

	url := e.formConnect.url.GetText()
	if url == "" {
		if e.errorFunc != nil {
			e.errorFunc(errors.New("URL em branco"))
		}
		return
	}
	req.SetUrl(url)

	for k := range e.header {
		req.AddHeader(e.header[k].Key, e.header[k].Value)
	}
	for k := range e.param {
		req.AddHeader(e.param[k].Key, e.param[k].Value)
	}
	response := req.Request()

	// Monta a resposta do servidor de teste
	err := json.Unmarshal([]byte(response), &e.tokenServer.Token)
	if err != nil {
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New("ConnectServerToken().getToken().json.Unmarshal().error"), err))
		}
		return
	}

	// Escreve o token recebido na UI
	e.formToken.token.SetText(e.tokenServer.Token.AccessToken)
	e.formToken.tokenType.SetText(e.tokenServer.Token.TokenType)
	e.formToken.status.SetText(e.tokenServer.Token.Status)
	e.formToken.issuedAt.SetText(strconv.FormatInt(e.tokenServer.Token.IssuedAt, 10))
	e.formToken.expiresIn.SetText(strconv.FormatInt(e.tokenServer.Token.ExpiresIn, 10))

	if e.event != nil {
		e.event(e.tokenServer.Token)
	}
}

func (e *ConnectServerToken) SetConfigTokenUrl() {
	e.formConnect.url.SetText(e.tokenServer.TestServer.GetUrl())
}
func (e *ConnectServerToken) mountFormToken() {

	e.form = tview.NewForm()

	e.form.AddTextView("Envio:", "", e.fieldWidth, 1, false, false)

	e.formConnect.Init(e)
	e.formConnect.Mount(e.form)

	e.form.AddButton("Set teste url", e.SetConfigTokenUrl)
	e.form.AddButton("Get token", e.GetToken)

	e.form.AddTextView("Resposta:", "", e.fieldWidth, 1, false, false)

	e.formToken.Init(&e.token, e.fieldWidth)
	e.formToken.Mount(e.form)

}
