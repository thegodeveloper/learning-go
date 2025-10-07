package boxui

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"

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
	list.SetBorderPadding(0, 0, 1, 0)
	list.SetBorderColor(tcell.ColorAquaMarine)

	// Right: output area
	output := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() { app.Draw() })
	output.SetBorder(true)
	output.SetBorderColor(tcell.ColorAquaMarine)
	output.SetTitle(" Output ")
	output.SetTitleAlign(tview.AlignCenter)

	// Bottom: toolbox bar
	footer := tview.NewTextView()
	footer.SetDynamicColors(true)
	footer.SetRegions(true)
	footer.SetWrap(false)
	footer.SetTextAlign(tview.AlignCenter)
	footer.SetBorder(true)
	footer.SetBorderColor(tcell.ColorAquaMarine)
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

	// Select the first item and set `current`
	if len(names) > 0 {
		list.SetCurrentItem(0)
		current = names[0]
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
	mainFlex.SetBorder(false)

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFlex, 0, 1, true).
		AddItem(footer, 3, 0, false)

	if err := app.SetRoot(root, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func startRun(app *tview.Application, output *tview.TextView, name string) {
	runLock.Lock()
	if runActive {
		runLock.Unlock()
		return
	}

	runActive = true
	runLock.Unlock()

	go func() {
		app.QueueUpdateDraw(func() {
			output.Clear()
			fmt.Fprintf(output, "[green]You selected package: [white]%s\n", name)
		})

		defer func() {
			runLock.Lock()
			runActive = false
			runLock.Unlock()
		}()

		fn, ok := registry.Packages[name]
		if !ok {
			app.QueueUpdateDraw(func() {
				fmt.Fprintf(output, "[red]Error: Package '%s' not found in registry.", name)
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		originalStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		outC := make(chan string)
		go func() {
			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			outC <- buf.String()
		}()

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Fprintf(w, "\n\n[red]FATAL ERROR (panic): %v\n[red]%s", r, string(debug.Stack()))
				}
				wg.Done()
			}()
			fn(true)
		}()

		// Wait for either the function to finish or the timeout to occur.
		doneChan := make(chan struct{})
		go func() {
			wg.Wait()
			close(doneChan)
		}()

		var timedOut bool
		select {
		case <-doneChan:
			// Execution finished normally.
		case <-ctx.Done():
			// Timeout occurred.
			timedOut = true
		}

		os.Stdout = originalStdout
		_ = w.Close()

		capturedOutput := <-outC
		if timedOut {
			capturedOutput += fmt.Sprintf("\n\n[red]Execution timed out after 5 seconds.\n[yellow]This usually means the package is waiting for interactive input, which is not supported in this TUI.")
		}

		app.QueueUpdateDraw(func() {
			output.SetText(capturedOutput)
			output.ScrollToBeginning()
		})

	}()
}
