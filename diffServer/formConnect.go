package diffServer

import "github.com/rivo/tview"

// FormConnect Monta o formulário de conexão https
type FormConnect struct {
	header     []*tview.InputField
	param      []*tview.InputField
	url        *tview.InputField
	method     *tview.DropDown
	content    *tview.InputField
	fieldWidth *tview.InputField
	status     *tview.TextView
}

// SetStatus define o estado mostrado ao usuário na tela
func (e *FormConnect) SetStatus(status string) {
	e.status.SetText(status)
}

// Init Configura o formulário
func (e *FormConnect) Init(data ConnectDataForm) {
	fieldWidth := data.GetFieldWidth()

	header := data.GetHeader()
	e.header = make([]*tview.InputField, len(header))
	for k := range header {
		e.header[k] = tview.NewInputField().
			SetLabel(header[k].Key).
			SetText(header[k].Value).
			SetFieldWidth(fieldWidth).
			SetChangedFunc(func(text string) {
				data.SetHeader(header[k].Key, header[k].Value)
			})
	}

	param := data.GetParam()
	e.param = make([]*tview.InputField, len(param))
	for k := range param {
		e.param[k] = tview.NewInputField().
			SetLabel(param[k].Key).
			SetText(param[k].Value).
			SetFieldWidth(fieldWidth).
			SetChangedFunc(func(text string) {
				data.SetHeader(param[k].Key, param[k].Value)
			})
	}

	e.url = tview.NewInputField().
		SetLabel("URL").
		SetFieldWidth(fieldWidth).
		SetChangedFunc(func(text string) {
			data.SetUrl(text)
		})

	e.method = tview.NewDropDown().
		SetLabel("Método").
		SetOptions(KMethod, func(text string, _ int) {
			data.SetMethod(text)
		}).SetCurrentOption(0)

	e.content = tview.NewInputField().
		SetLabel("Conteúdo").
		SetFieldWidth(fieldWidth).
		SetChangedFunc(func(text string) {
			data.SetContent(text)
		})

	e.status = tview.NewTextView().
		SetSize(1, fieldWidth).
		SetLabel("Estado")
}

// Mount Monta o formulário
func (e *FormConnect) Mount(form *tview.Form) {
	form.AddFormItem(e.method)
	form.AddFormItem(e.url)
	for k := range e.header {
		form.AddFormItem(e.header[k])
	}
	for k := range e.param {
		form.AddFormItem(e.param[k])
	}
	form.AddFormItem(e.content)
	form.AddFormItem(e.status)
}
