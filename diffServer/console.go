package diffServer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/term"
	"time"
)

// Console Monta a aplicação principal
type Console struct {
	app                *tview.Application
	menu               *tview.List
	pages              *tview.Pages
	modal              *tview.Modal
	pageName           string
	mainGrid           *tview.Grid
	connectServerAData ConnectServerData
	connectServerBData ConnectServerData
	dataController     DataController
	compareController  CompareController
	fieldWidth         int
	screenWidth        int
	screenHeight       int
}

func (e *Console) newGridCompare() {
	e.compareController.Init()
}

func (e *Console) newModalError() {
	e.modal = tview.NewModal()
	e.modal.SetText("vivo!")
	e.modal.AddButtons([]string{"Quit"})
	e.modal.SetDoneFunc(
		func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				e.app.Stop()
			}
		},
	)
}

func (e *Console) SetError(err error) {
	e.modal.SetText(err.Error())
	e.app.SetFocus(e.pages)
	e.app.SetFocus(e.modal)
	e.pages.SendToFront("error")
}

func (e *Console) newPages() {
	e.pages = tview.NewPages()
	e.pages.SetBorder(true)
	e.pages.AddPage("error", e.modal, false, true)

	e.pageName = "mainDataTest"
	e.pages.AddPage(e.pageName, e.dataController.form, true, true)

	e.pageName = "dataCompare"
	e.pages.AddPage(e.pageName, e.compareController.GetGrid(), true, true)

	e.pageName = "serverAData"
	e.pages.AddPage(e.pageName, e.connectServerAData.form, true, true)
	e.pageName = "serverBData"
	e.pages.AddPage(e.pageName, e.connectServerBData.form, true, true)

	e.pageName = "serverAToken"
	e.pages.AddPage(e.pageName, e.connectServerAData.ConnectServerToken.form, true, true)
	e.pageName = "serverBToken"
	e.pages.AddPage(e.pageName, e.connectServerBData.ConnectServerToken.form, true, true)

	e.pageName = "serverAData"
	e.pages.SendToFront(e.pageName)
}

func (e *Console) newGrid() {
	e.mainGrid = tview.NewGrid()
	e.mainGrid.SetRows(1)
	e.mainGrid.SetColumns(30, -1) //e.screenWidth-35
	e.mainGrid.SetBorders(true)

	// Layout para telas estreitas (menu e barra lateral ocultos)
	e.mainGrid.AddItem(e.menu, 0, 0, 0, 0, 0, 0, true)
	e.mainGrid.AddItem(e.pages, 1, 0, 1, 3, 0, 0, false)

	// Layout para telas largas
	e.mainGrid.AddItem(e.menu, 1, 0, 1, 1, 0, 100, true)
	e.mainGrid.AddItem(e.pages, 1, 1, 1, 1, 0, 100, false)
}

func (e *Console) newMenu() {
	e.menu = tview.NewList().
		AddItem(
			"ServerA Data",
			"Server A data",
			'a',
			func() {
				e.pages.SendToFront("serverAData")
				e.changeFocus()
			},
		).
		AddItem(
			"ServerA Tk",
			"Server A token",
			'b',
			func() {
				e.pages.SendToFront("serverAToken")
				e.changeFocus()
			},
		).
		AddItem(
			"ServerB Data",
			"Server B data",
			'c',
			func() {
				e.pages.SendToFront("serverBData")
				e.changeFocus()
			},
		).
		AddItem(
			"ServerB Tk",
			"Server B Token",
			'd',
			func() {
				e.pages.SendToFront("serverBToken")
				e.changeFocus()
			},
		).
		AddItem(
			"Data test",
			"Data test config",
			'e',
			func() {
				e.pages.SendToFront("mainDataTest")
				e.dataController.UpdateStatus()
				e.changeFocus()
			},
		).
		AddItem(
			"Data compare",
			"Data compare",
			'f',
			func() {
				e.pages.SendToFront("dataCompare")
				e.changeFocus()
				bf := e.compareController.CompareGrid.bottomForm
				e.app.SetFocus(bf)
			},
		).
		//AddItem("Compare", "Compare json", 'c', compareFunc).
		//AddItem("Sort", "Ajusta o json", 's', completeSortJsonFunc).
		AddItem(
			"Quit",
			"Press to exit",
			'q',
			func() {
				e.app.Stop()
			},
		)
	e.menu.SetBorder(true)
}

func (e *Console) changeFocus() {
	if e.menu.HasFocus() {
		e.app.SetFocus(e.pages)
		//e.pages.SendToFront(e.pageName)
	} else if e.pages.HasFocus() {
		e.app.SetFocus(e.menu)
	}
}

func (e *Console) Init() {
	var err error

	e.screenWidth, e.screenHeight, err = term.GetSize(0)
	if err != nil {
		e.screenWidth = 80
	}

	e.fieldWidth = e.screenWidth - 35 - 20

	e.app = tview.NewApplication()
	e.app.EnablePaste(true)
	e.app.EnableMouse(true)

	e.connectServerAData.SetErrorFunc(e.SetError)
	e.connectServerAData.SetDataSend(&e.dataController.DataTestA)
	e.connectServerAData.SetDataReceiver(&e.dataController.DataServerA)
	e.connectServerAData.Init(e.fieldWidth)

	e.connectServerBData.SetErrorFunc(e.SetError)
	e.connectServerBData.SetDataSend(&e.dataController.DataTestB)
	e.connectServerBData.SetDataReceiver(&e.dataController.DataServerB)
	e.connectServerBData.Init(e.fieldWidth)

	e.dataController.SetUpdateKeys(e.compareController.SetDataKeys)
	e.dataController.SetErrorFunc(e.SetError)
	e.dataController.Init(e.fieldWidth)

	e.compareController.SetErrorFunc(e.SetError)
	e.compareController.SetDataAPointer(&e.dataController.DataServerA)
	e.compareController.SetDataBPointer(&e.dataController.DataServerB)

	e.newGridCompare()
	e.newModalError()
	e.newPages()
	e.newMenu()
	e.newGrid()

	//e.connectServerAData.SetEventOnButtonGetTestData(func() {
	//	os.WriteFile("event.txt", []byte("entrou aqui"), 0655)
	//	e.dataController.UpdateStatus()
	//})

	// tview tem um bug com eventos assíncronos a UI não redesenha
	go func() {
		for {
			time.Sleep(1 * time.Second)
			e.app.Draw()
		}
	}()

	e.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlSpace && event.Modifiers() == tcell.ModCtrl {
			e.changeFocus()
			return nil
		}
		return event
	})

	// Executa a aplicação
	if err := e.app.SetRoot(e.mainGrid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
