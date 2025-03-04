package diffServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rivo/tview"
	"strconv"
)

type ConnectServerData struct {
	ConnectServerToken ConnectServerToken

	header              []keyValue
	param               []keyValue
	url                 string
	method              string
	content             string
	fieldWidth          int
	form                *tview.Form
	formConnect         FormConnect
	dataServer          TestServer
	dataSendPointer     *Data
	dataReceiverPointer *Data
	//eventOnButtonGetTestToken func()
	//eventOnButtonGetRealToken func()
	//eventOnButtonGetTestData  func()
	errorFunc func(error)
}

func (e *ConnectServerData) SetErrorFunc(f func(error)) {
	e.errorFunc = f
	e.ConnectServerToken.errorFunc = f
}

func (e *ConnectServerData) SetDataSend(data *Data) {
	e.dataSendPointer = data
}

func (e *ConnectServerData) SetDataReceiver(data *Data) {
	e.dataReceiverPointer = data
}

// SetStatus define o estado mostrado ao usuário na tela
func (e *ConnectServerData) SetStatus(status string) {
	e.formConnect.SetStatus(status)
}

func (e *ConnectServerData) SetStatusNeutro() {
	e.SetStatus(KStatusWaitUserAction)
}

func (e *ConnectServerData) SetUrl(url string) {
	e.url = url
}

func (e *ConnectServerData) SetMethod(method string) {
	e.method = method
}

func (e *ConnectServerData) SetContent(content string) {
	e.content = content
}

func (e *ConnectServerData) SetHeader(key, value string) {
	for k := range e.header {
		if e.header[k].Key == key {
			e.header[k].Value = value
			return
		}
	}
}

func (e *ConnectServerData) SetParam(key, value string) {
	for k := range e.param {
		if e.param[k].Key == key {
			e.param[k].Value = value
			return
		}
	}
}

func (e *ConnectServerData) GetHeader() []keyValue {
	return e.header
}

func (e *ConnectServerData) GetParam() []keyValue {
	return e.param
}

func (e *ConnectServerData) GetUrl() string {
	return e.url
}

func (e *ConnectServerData) GetMethod() string {
	return e.method
}

func (e *ConnectServerData) GetContent() string {
	return e.content
}

func (e *ConnectServerData) GetFieldWidth() int {
	return e.fieldWidth
}

func (e *ConnectServerData) GetFormToken() *tview.Form {
	return e.form
}

func (e *ConnectServerData) GetTokenServer() TestServer {
	return e.dataServer
}

func (e *ConnectServerData) Init(fieldWidth int) {
	e.ConnectServerToken.Init(fieldWidth)

	e.fieldWidth = fieldWidth
	e.dataServer.Init()

	e.header = make([]keyValue, 0)

	// Não apagar ou modificar este header, ele afeta a função onToken(tk token)
	e.header = append(e.header, keyValue{"Authorization", ""})

	e.param = make([]keyValue, 0)
	//e.param = append(e.param, keyValue{"user", ""})
	//e.param = append(e.param, keyValue{"password", ""})

	// Define um ponteiro de função para quando o token é recebido
	e.ConnectServerToken.SetEvent(e.onToken)

	e.mountFormData()
	e.SetStatusNeutro()
}

// onToken função chamada quando o token é recebido
func (e *ConnectServerData) onToken(tk Token) {
	for k := range e.header {
		if e.header[k].Key == "Authorization" {
			e.header[k].Value = tk.TokenType + " " + tk.AccessToken

			for kh := range e.formConnect.header {
				if e.formConnect.header[kh].GetLabel() == "Authorization" {
					e.formConnect.header[kh].SetText(e.header[k].Value)
				}
			}
			break
		}
	}
}

func (e *ConnectServerData) mountFormData() {

	e.form = tview.NewForm()

	e.form.AddTextView("Envio:", "", e.fieldWidth, 1, false, false)

	e.formConnect.Init(e)
	e.formConnect.Mount(e.form)

	//e.form.AddButton("Get token", e.buttonGetTestToken)
	//e.form.AddButton("real token", e.buttonGetRealToken)
	e.form.AddButton("test data", e.buttonGetTestData)

	//e.form.AddTextView("Resposta:", "", e.fieldWidth, 1, false, false)

	//e.formToken.Init(&e.token, e.fieldWidth)
	//e.formToken.Mount(e.form)

}

//func (e *ConnectServerData) SetEventOnButtonGetTestToken(f func()) {
//	e.eventOnButtonGetTestToken = f
//}
//
//func (e *ConnectServerData) SetEventOnButtonGetRealToken(f func()) {
//	e.eventOnButtonGetRealToken = f
//}
//
//func (e *ConnectServerData) SetEventOnButtonGetTestData(f func()) {
//	e.eventOnButtonGetTestData = f
//}

func (e *ConnectServerData) buttonGetTestData() {
	//for kh := range e.formConnect.header {
	//	if e.formConnect.header[kh].GetLabel() == "Authorization" {
	//		if e.formConnect.header[kh].GetText() == "" {
	//			e.buttonGetTestToken()
	//			break
	//		}
	//	}
	//}
	e.SetStatus("Aguardando dado de teste")
	//go func() {
	//defer e.SetStatusNeutro()
	e.GetTestData()

	l := len(e.dataReceiverPointer.Data)
	i := "chaves"
	if l == 1 {
		i = "chave"
	}
	e.SetStatus(fmt.Sprintf("Recebido: %v %v", strconv.FormatInt(int64(l), 10), i))

	//if e.eventOnButtonGetTestData != nil {
	//	e.eventOnButtonGetTestData()
	//}
	//}()
}

//func (e *ConnectServerData) buttonGetRealToken() {
//	e.SetStatus("Aguardando token real")
//	//go func() {
//	//	defer e.SetStatusNeutro()
//	e.ConnectServerToken.GetRealToken()
//
//	//if e.eventOnButtonGetRealToken != nil {
//	//	e.eventOnButtonGetRealToken()
//	//}
//	//}()
//	e.SetStatusNeutro()
//}

//func (e *ConnectServerData) buttonGetTestToken() {
//	e.SetStatus("Aguardando token de teste")
//	//go func() {
//	//	defer e.SetStatusNeutro()
//	e.ConnectServerToken.GetToken()
//
//	//if e.eventOnButtonGetTestToken != nil {
//	//	e.eventOnButtonGetTestToken()
//	//}
//	//}()
//	e.SetStatusNeutro()
//}

func (e *ConnectServerData) GetTestData() {
	e.getData(true)
}

func (e *ConnectServerData) getData(test bool) {

	// Define a resposta do servidor de teste
	e.dataServer.Init()
	e.dataServer.SetResponse(e.dataSendPointer)

	// Informa a URL do servidor de teste na UI
	if test && e.formConnect.url.GetText() == "" {
		e.formConnect.url.SetText(e.dataServer.GetUrl())
	}

	// Prepara a requisição
	req := new(HttpRequest)
	req.SetMethod(e.method)

	if test {
		req.SetUrl(e.dataServer.GetUrl())
	} else {
		req.SetUrl(e.formConnect.url.GetText())
	}

	for k := range e.header {
		req.AddHeader(e.header[k].Key, e.header[k].Value)
	}
	for k := range e.param {
		req.AddHeader(e.param[k].Key, e.param[k].Value)
	}
	response := req.Request()

	// Monta a resposta do servidor de teste
	err := json.Unmarshal([]byte(response), e.dataReceiverPointer)
	if err != nil {
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New(fmt.Sprintf("ConnectServerData().getData(%v).json.Unmarshal().error", test)), err))
		}
		return
	}

	//// Escreve o token recebido na UI
	//e.formToken.token.SetText(e.tokenServer.token.AccessToken)
	//e.formToken.tokenType.SetText(e.tokenServer.token.TokenType)
	//e.formToken.status.SetText(e.tokenServer.token.Status)
	//e.formToken.issuedAt.SetText(strconv.FormatInt(e.tokenServer.token.IssuedAt, 10))
	//e.formToken.expiresIn.SetText(strconv.FormatInt(e.tokenServer.token.ExpiresIn, 10))
	//
	//if e.event != nil {
	//	e.event(e.tokenServer.token)
	//}
}
