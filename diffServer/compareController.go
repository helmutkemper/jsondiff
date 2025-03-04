package diffServer

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/rivo/tview"
)

type CompareController struct {
	CompareGrid  CompareGrid
	dataAPointer *Data
	dataBPointer *Data
	indexErr     int
	index        int
	dataKeys     []string
	errorFunc    func(error)
}

func (e *CompareController) SetDataKeys(keys []string) {
	e.dataKeys = keys
}

func (e *CompareController) SetErrorFunc(f func(error)) {
	e.errorFunc = f
}

func (e *CompareController) SetLog(text string) {
	e.CompareGrid.SetLog(text)
}

func (e *CompareController) SetTextA(text string) {
	e.CompareGrid.SetTextA(text)
}

func (e *CompareController) SetTextB(text string) {
	e.CompareGrid.SetTextB(text)
}

func (e *CompareController) SetDataAPointer(data *Data) {
	e.dataAPointer = data
}

func (e *CompareController) SetDataBPointer(data *Data) {
	e.dataBPointer = data
}

func (e *CompareController) Init() {
	e.CompareGrid.SetComparePointer(e.Compare)
	e.CompareGrid.SetNextPointerErr(e.ViewNextErr)
	e.CompareGrid.SetPrevPointerErr(e.ViewPrevErr)
	e.CompareGrid.SetNextPointer(e.ViewNext)
	e.CompareGrid.SetPrevPointer(e.ViewPrev)
	e.CompareGrid.Init()
}

func (e *CompareController) GetGrid() *tview.Grid {
	return e.CompareGrid.GetGrid()
}

func (e *CompareController) ViewNextErr() {
	if e.indexErr < e.GetTotalElements()-1 {
		e.indexErr += 1
	}

	jsA, jsB := e.GetElementsErr(e.indexErr)

	e.SetTextA(string(jsA))
	e.SetTextB(string(jsB))
}

func (e *CompareController) ViewPrevErr() {
	if e.indexErr > 0 {
		e.indexErr -= 1
	}

	jsA, jsB := e.GetElementsErr(e.indexErr)

	e.SetTextA(string(jsA))
	e.SetTextB(string(jsB))
}

func (e *CompareController) ViewNext() {
	if e.index < len(e.dataAPointer.Data)-1 {
		e.index += 1
	}

	jsA, jsB := e.GetElements(e.index)

	e.SetTextA(string(jsA))
	e.SetTextB(string(jsB))
}

func (e *CompareController) ViewPrev() {
	if e.index > 0 {
		e.index -= 1
	}

	jsA, jsB := e.GetElementsErr(e.index)

	e.SetTextA(string(jsA))
	e.SetTextB(string(jsB))
}

func (e *CompareController) jsColor(jsA, jsB []byte) (coloredA, coloredB []byte) {
	jsASplitted := bytes.Split(jsA, []byte("\n"))
	jsBSplitted := bytes.Split(jsB, []byte("\n"))

	for i := 0; i < len(jsASplitted); i++ {
		if bytes.Equal(jsASplitted[i], jsBSplitted[i]) {
			jsASplitted[i] = append(append([]byte("[green]"), jsASplitted[i]...), []byte("[white]")...)
			jsBSplitted[i] = append(append([]byte("[green]"), jsBSplitted[i]...), []byte("[white]")...)
		} else {
			jsASplitted[i] = append(append([]byte("[red]"), jsASplitted[i]...), []byte("[white]")...)
			jsBSplitted[i] = append(append([]byte("[red]"), jsBSplitted[i]...), []byte("[white]")...)
		}
	}

	coloredA = bytes.Join(jsASplitted, []byte("\n"))
	coloredB = bytes.Join(jsBSplitted, []byte("\n"))

	return
}

func (e *CompareController) Compare() {

	if len(e.dataAPointer.Data) == 0 {
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New("CompareController().Compare().error"), errors.New("the A server did not receive data")))
		}
		return
	}

	if len(e.dataBPointer.Data) == 0 {
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New("CompareController().Compare().error"), errors.New("the B server did not receive data")))
		}
		return
	}

	e.indexErr = 0
	e.index = 0

	e.dataAPointer.Compare(e.dataBPointer.Data, e.dataKeys)
	e.SetLog(e.dataAPointer.GetLog())

	jsA, jsB := e.GetElementsErr(0)

	e.SetTextA(string(jsA))
	e.SetTextB(string(jsB))
}

func (e *CompareController) GetTotalElements() int {
	return e.dataAPointer.GetTotalElements()
}

func (e *CompareController) GetElementA() []any {
	return e.dataAPointer.GetElementA()
}

func (e *CompareController) GetElementB() []any {
	return e.dataBPointer.GetElementB()
}

func (e *CompareController) GetElementsErr(index int) (elementA, elementB []byte) {
	elementA, elementB = e.dataAPointer.GetElements(index)
	return e.jsColor(elementA, elementB)
}

func (e *CompareController) GetElements(index int) (elementA, elementB []byte) {
	var err error
	if index >= 0 && index < len(e.dataAPointer.Data) {
		elementA, err = json.MarshalIndent(e.dataAPointer.Data[index], "", "  ")
		if err != nil {
			if e.errorFunc != nil {
				e.errorFunc(errors.Join(errors.New("CompareController().GetElements().json.Unmarshal(A).error"), err))
			}
			return
		}

		elementB, err = json.MarshalIndent(e.dataAPointer.Data[index], "", "  ")
		if err != nil {
			if e.errorFunc != nil {
				e.errorFunc(errors.Join(errors.New("CompareController().GetElements().json.Unmarshal(A).error"), err))
			}
			return
		}

		elementA, elementB = e.jsColor(elementA, elementB)
		return elementA, elementB
	}

	return nil, nil
}

func (e *CompareController) GetLog() string {
	return e.dataAPointer.GetLog()
}
