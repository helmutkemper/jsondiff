package diffServer

import (
	"fmt"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

// DataController Gera os dados a serem usados nos testes e recebe os dados reais.
type DataController struct {
	DataTestA   Data
	DataTestB   Data
	DataServerA Data
	DataServerB Data

	fieldWidth int
	formData   FormDataTest
	form       *tview.Form

	dataKeys     []string
	amountOfData int
	interactions int
	numberOfKeys int
	deleteKeys   int
}

func (e *DataController) GetFieldWidth() (fieldWidth int) {
	return e.fieldWidth
}

func (e *DataController) SetDeleteKeys(deleteKeys string) {
	value, err := strconv.Atoi(deleteKeys)
	if err != nil {
		value = 0
	}
	e.deleteKeys = value
}

func (e *DataController) SetDataKeys(keys string) {
	keys = strings.ReplaceAll(keys, " ", "")
	e.dataKeys = strings.Split(keys, ",")
}

func (e *DataController) SetAmount(amount string) {
	value, err := strconv.Atoi(amount)
	if err != nil {
		value = 0
	}
	e.amountOfData = value
}

func (e *DataController) SetInteractions(interactions string) {
	value, err := strconv.Atoi(interactions)
	if err != nil {
		value = 0
	}
	e.interactions = value
}

func (e *DataController) SetNumberOfKeys(numberOfKeys string) {
	value, err := strconv.Atoi(numberOfKeys)
	if err != nil {
		value = 0
	}
	e.numberOfKeys = value
}

// SetStatus define o estado mostrado ao usuário na tela
func (e *DataController) SetStatus(status string) {
	e.formData.SetStatus(status)
}

func (e *DataController) SetStatusNeutro() {
	e.SetStatus(KStatusWaitUserAction)
}

func (e *DataController) updateValuesFromFields() {
	e.SetDataKeys(e.formData.GetDataKeys())
	e.SetAmount(e.formData.GetAmount())
	e.SetInteractions(e.formData.GetInteractions())
	e.SetNumberOfKeys(e.formData.GetNumberOfKeys())
	e.SetDeleteKeys(e.formData.GetDeleteKeys())
}

func (e *DataController) buttonGenerate() {
	e.updateValuesFromFields()

	// Monda os dados do servidor A
	e.DataTestA.Init(e.amountOfData)
	// Copia os dados do servidor A para o servidor B e bagunça os dados
	e.DataTestB.CopyAndPrepare(e.DataTestA, e.interactions, e.numberOfKeys, e.deleteKeys)

	e.UpdateStatus()
}

func (e *DataController) UpdateStatus() {
	e.SetStatus(fmt.Sprintf("Dados de teste: %v, recebido em A: %v, recebido em B: %v", len(e.DataTestA.Data), len(e.DataServerA.Data), len(e.DataServerB.Data)))
}

func (e *DataController) Init(fieldWidth int) {
	e.fieldWidth = fieldWidth

	e.mountFormData()
	e.SetStatus("Gerando dados de teste")
	e.buttonGenerate()
}

func (e *DataController) mountFormData() {

	e.form = tview.NewForm()

	e.formData.Init(e)
	e.formData.Mount(e.form)

	//e.form.AddTextView("Teste:", "", e.fieldWidth, 1, false, false)

	e.form.AddButton("Gerar", e.buttonGenerate)

}
