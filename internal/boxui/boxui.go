package boxui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run(show bool) {
	if show {
		app := tview.NewApplication()

		box := tview.NewBox().
			SetBorder(true).
			SetBorderPadding(1, 1, 1, 1).
			SetTitle(" Learning Go ").
			SetTitleAlign(tview.AlignCenter).
			SetBorderColor(tcell.ColorDeepSkyBlue)

		//list := tview.NewList().
		//	AddItem("List item 1", "Some explanatory text", 'a', nil).
		//	AddItem("List item 2", "Some explanatory text", 'b', nil).
		//	AddItem("List item 3", "Some explanatory text", 'c', nil).
		//	AddItem("List item 4", "Some explanatory text", 'd', nil).
		//	AddItem("Quit", "Press to exit", 'q', func() {
		//		app.Stop()
		//	})

		flex := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(box, 0, 1, false) //.
		//AddItem(list, 0, 3, true)

		if err := app.SetRoot(flex, true).Run(); err != nil {
			panic(err)
		}
	}
}
