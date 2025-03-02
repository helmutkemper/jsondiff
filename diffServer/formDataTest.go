package diffServer

import (
	"github.com/rivo/tview"
)

// FormDataTest Gera o formulário de controle e de teste dos dados
type FormDataTest struct {
	dataKeys          *tview.InputField // todo: autocomplete
	amountOfData      *tview.InputField
	interactions      *tview.InputField
	numberOfKeys      *tview.InputField
	deleteOfKeys      *tview.InputField
	screenExplanation *tview.TextView
	status            *tview.TextView
}

// SetStatus define o estado mostrado ao usuário na tela
func (e *FormDataTest) SetStatus(status string) {
	e.status.SetText(status)
}

func (e *FormDataTest) GetDataKeys() string {
	return e.dataKeys.GetText()
}

func (e *FormDataTest) GetDeleteKeys() string {
	return e.deleteOfKeys.GetText()
}

func (e *FormDataTest) GetAmount() string {
	return e.amountOfData.GetText()
}

func (e *FormDataTest) GetInteractions() string {
	return e.interactions.GetText()
}

func (e *FormDataTest) GetNumberOfKeys() string {
	return e.numberOfKeys.GetText()
}

// Init Configura o formulário
func (e *FormDataTest) Init(data DataConfig) {
	fieldWidth := data.GetFieldWidth()

	e.dataKeys = tview.NewInputField().
		SetLabel("Chaves").
		SetPlaceholder("Ex. ID,ID_PESSOA").
		SetFieldWidth(fieldWidth).
		SetText("ID").
		SetChangedFunc(func(text string) {
			data.SetDataKeys(text)
		})

	/*InputFieldInteger = func(text string, ch rune) bool {
		if text == "-" {
			return true
		}
		_, err := strconv.Atoi(text)
		return err == nil
	}*/

	e.amountOfData = tview.NewInputField().
		SetLabel("Quantidade").
		SetPlaceholder("Ex. 100000").
		SetText("100000").
		SetFieldWidth(fieldWidth).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetChangedFunc(func(text string) {
			data.SetAmount(text)
		})

	e.interactions = tview.NewInputField().
		SetLabel("Falhas").
		SetPlaceholder("Ex. 100").
		SetText("100").
		SetFieldWidth(fieldWidth).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetChangedFunc(func(text string) {
			data.SetInteractions(text)
		})

	e.numberOfKeys = tview.NewInputField().
		SetLabel("Interações").
		SetPlaceholder("Ex. 2").
		SetText("2").
		SetFieldWidth(fieldWidth).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetChangedFunc(func(text string) {
			data.SetNumberOfKeys(text)
		})

	e.deleteOfKeys = tview.NewInputField().
		SetLabel("Apagar").
		SetPlaceholder("Ex. 3").
		SetText("3").
		SetFieldWidth(fieldWidth).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetChangedFunc(func(text string) {
			data.SetDeleteKeys(text)
		})

	e.status = tview.NewTextView().
		SetSize(1, fieldWidth).
		SetLabel("Estado")

	e.screenExplanation = tview.NewTextView().
		SetSize(13, fieldWidth).
		SetWrap(true).
		SetText("\n" +
			"Página de Dados e de teste:" +
			"\n" +
			"\n" +
			"Esta página configura o teste do código, " +
			"gerando dados controlados para que o sistema possa ser testado." +
			"\n" +
			"\n" +
			"Chaves:\t\tDefinas as chaves usadas na comparação de dados;\n" +
			"\t\t\t  * As chaves devem ser separadas por ponto e vírgula, sem espaços.\n" +
			"\t\t\t  * As chaves são usadas no dado real.\n" +
			"Quantidade:\tDefine a quantidade de dados aleatórios gerados.\n" +
			"Falhas:\t\tDefine a quantidade de dados com falhas.\n" +
			"Interações:\tDefine a quantidade de falhas dentro de um mesmo dado.\n" +
			"Apagar:\t\tApaga algumas chaves no dado gerado.")
}

// Mount Monta o formulário
func (e *FormDataTest) Mount(form *tview.Form) {
	form.AddFormItem(e.dataKeys)
	form.AddFormItem(e.amountOfData)
	form.AddFormItem(e.interactions)
	form.AddFormItem(e.numberOfKeys)
	form.AddFormItem(e.deleteOfKeys)
	form.AddFormItem(e.status)
	form.AddFormItem(e.screenExplanation)
}
