package diffServer

import (
	"github.com/rivo/tview"
)

type CompareGrid struct {
	dataServerA    *tview.TextView
	dataServerB    *tview.TextView
	bottomForm     *tview.Form
	grid           *tview.Grid
	logArea        *tview.TextView
	comparePointer func()
	prevPointerErr func()
	nextPointerErr func()
	prevPointer    func()
	nextPointer    func()
}

func (e *CompareGrid) SetComparePointer(f func()) {
	e.comparePointer = f
}

func (e *CompareGrid) SetPrevPointerErr(f func()) {
	e.prevPointerErr = f
}

func (e *CompareGrid) SetNextPointerErr(f func()) {
	e.nextPointerErr = f
}

func (e *CompareGrid) SetPrevPointer(f func()) {
	e.prevPointer = f
}

func (e *CompareGrid) SetNextPointer(f func()) {
	e.nextPointer = f
}

func (e *CompareGrid) SetLog(text string) {
	e.logArea.SetText(text)
}

func (e *CompareGrid) SetTextA(text string) {
	e.dataServerA.SetText(text)
}

func (e *CompareGrid) SetTextB(text string) {
	e.dataServerB.SetText(text)
}

func (e *CompareGrid) Init() {
	e.dataServerA = tview.NewTextView()
	//e.dataServerA.SetText("Olá mundo!", false)
	e.dataServerA.SetBorder(true)
	e.dataServerA.SetDynamicColors(true)
	e.dataServerA.SetScrollable(true)

	e.dataServerB = tview.NewTextView()
	//e.dataServerB.SetText("Olá mundo!", false)
	e.dataServerB.SetBorder(true)
	e.dataServerB.SetDynamicColors(true)
	e.dataServerB.SetScrollable(true)

	e.bottomForm = tview.NewForm()
	//e.bottom.SetBorder(false)
	//e.bottom.SetTitle("Bottom")

	e.logArea = tview.NewTextView()
	e.logArea.SetText("Compare:\tCompara os dois dados recebidos.\n" +
		"<:\t\t\tMostra o dado recebido.\n" +
		">:\t\t\tMostra o dado recebido.\n" +
		"<:\t\t\tMostra o dado problemático.\n" +
		">:\t\t\tMostra o dado problemático.")
	e.logArea.SetSize(10, -1)
	e.logArea.SetBorder(true)

	e.bottomForm.AddFormItem(e.logArea)
	e.bottomForm.AddButton("Compare", e.comparePointer)
	e.bottomForm.AddButton("<", e.prevPointer)
	e.bottomForm.AddButton(">", e.nextPointer)
	e.bottomForm.AddButton("<", e.prevPointerErr)
	e.bottomForm.AddButton(">", e.nextPointerErr)

	e.grid = tview.NewGrid().
		SetRows(-1, -1).
		SetColumns(-1, -1)

	e.grid.AddItem(e.dataServerA, 0, 0, 1, 1, 0, 0, false)
	e.grid.AddItem(e.dataServerB, 0, 1, 1, 1, 0, 0, false)
	e.grid.AddItem(e.bottomForm, 1, 0, 1, 2, 0, 0, false)
}

func (e *CompareGrid) GetGrid() *tview.Grid {
	return e.grid
}
