package boxui

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/thegodeveloper/learning-go/internal/registry"
)

var runLock sync.Mutex
var runActive bool

func Run(show bool) {
	if !show {
		return
	}

	app := tview.NewApplication()

	// Left: list of packages
	list := tview.NewList()
	list.ShowSecondaryText(false)
	list.SetBorder(true)
	list.SetTitle(" Packages ")
	list.SetTitleAlign(tview.AlignCenter)

	// Right: output area
	output := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() { app.Draw() })
	output.SetBorder(true)
	output.SetTitle(" Output ")
	output.SetTitleAlign(tview.AlignCenter)

	// Bottom: toolbox bar
	footer := tview.NewTextView()
	footer.SetDynamicColors(true)
	footer.SetRegions(true)
	footer.SetWrap(false)
	footer.SetTextAlign(tview.AlignCenter)
	footer.SetBorder(true)
	footer.SetTitle(" Controls ")
	footer.SetTitleAlign(tview.AlignCenter)
	footer.SetText("[yellow]Enter[white]/[yellow]r[white] Run   [yellow]TAB[white] Switch Focus   [yellow]q[white] Quit")

	// track currently selected name
	var current string
	names := registry.Names()

	// populate list (no inline callback in AddItem)
	for _, name := range names {
		list.AddItem(name, "", 0, nil)
	}

	// update current when selection changes
	list.SetChangedFunc(func(i int, main, secondary string, r rune) {
		current = main
	})

	// keybindings: q to quit, r/Enter to run, Tab to switch focus
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Key() == tcell.KeyEnter:
			if current != "" {
				startRun(app, output, current)
			}
			return nil
		case event.Rune() == 'r' || event.Rune() == 'R':
			if current != "" {
				startRun(app, output, current)
			}
			return nil
		case event.Rune() == 'q' || event.Rune() == 'Q':
			app.Stop()
			return nil
		case event.Key() == tcell.KeyEsc:
			app.Stop()
			return nil
		case event.Key() == tcell.KeyTAB:
			if app.GetFocus() == list {
				app.SetFocus(output)
			} else {
				app.SetFocus(list)
			}
			return nil
		}
		return event
	})

	// Layout: list on left, output on right, footer below
	mainFlex := tview.NewFlex().
		AddItem(list, 30, 1, true).
		AddItem(output, 0, 3, false)

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFlex, 0, 1, true). // main content fills remaining space
		AddItem(footer, 3, 0, false)   // give footer a fixed height (3 lines)

	if err := app.SetRoot(root, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func startRun(app *tview.Application, output *tview.TextView, name string) {
	go func() {
		app.QueueUpdateDraw(func() {
			output.Clear()
			fmt.Fprintf(output, "[green]You selected package: [white]%s\n", name)
		})
	}()
}
