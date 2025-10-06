package boxui

import (
	"fmt"
	"os"
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

	// keybindings: q to quit, r to rerun, Tab to switch focus
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q', 'Q':
			app.Stop()
			return nil
		case 'r', 'R':
			if current != "" {
				startRun(app, output, current)
			}
			return nil
		}
		// Support Esc to quit
		if event.Key() == tcell.KeyEsc {
			app.Stop()
			return nil
		}
		// Tab switches focus between list and output
		if event.Key() == tcell.KeyTAB {
			if app.GetFocus() == list {
				app.SetFocus(output)
			} else {
				app.SetFocus(list)
			}
			return nil
		}
		return event
	})

	// Layout: list on left (fixed width), output on right
	flex := tview.NewFlex().
		AddItem(list, 30, 1, true).
		AddItem(output, 0, 3, false)

	// initial focus on list
	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

// startRun runs the named package (registry.Packages[name]) and streams stdout to output TextView.
// It prevents concurrent runs; if a run is active it prints a warning into the output pane.
func startRun(app *tview.Application, output *tview.TextView, name string) {
	// Prevent concurrent runs (simple guard)
	runLock.Lock()
	if runActive {
		// already running
		runLock.Unlock()
		app.QueueUpdateDraw(func() {
			fmt.Fprintf(output, "[red]A run is already in progress. Please wait...\n")
		})
		return
	}
	runActive = true
	runLock.Unlock()

	// Ensure we mark not running when finished
	doneCleanup := func() {
		runLock.Lock()
		runActive = false
		runLock.Unlock()
	}

	// Clear output and announce
	app.QueueUpdateDraw(func() {
		output.Clear()
		fmt.Fprintf(output, "[yellow]Running %s...\n\n", name)
	})

	runFunc, ok := registry.Packages[name] // expects map[string]func(bool)
	if !ok {
		app.QueueUpdateDraw(func() {
			fmt.Fprintf(output, "[red]Package %s not found in registry\n", name)
		})
		doneCleanup()
		return
	}

	// Create a pipe and redirect os.Stdout to the write end.
	pr, pw, err := os.Pipe()
	if err != nil {
		app.QueueUpdateDraw(func() {
			fmt.Fprintf(output, "[red]Failed to create pipe: %v\n", err)
		})
		doneCleanup()
		return
	}

	// Save current stdout so we can restore it.
	oldStdout := os.Stdout
	os.Stdout = pw

	// Goroutine: copy from the pipe reader to the TextView
	copyDone := make(chan struct{})
	go func() {
		defer pr.Close()
		buf := make([]byte, 1024)
		for {
			n, err := pr.Read(buf)
			if n > 0 {
				b := make([]byte, n)
				copy(b, buf[:n])
				app.QueueUpdateDraw(func() {
					_, _ = output.Write(b)
				})
			}
			if err != nil {
				break
			}
		}
		close(copyDone)
	}()

	// Goroutine: run the function (so UI stays responsive)
	go func() {
		// When the function returns, close the pipe writer and restore stdout
		defer func() {
			// Close writer to signal reader end-of-file
			_ = pw.Close()
			// restore stdout
			os.Stdout = oldStdout
		}()

		// Call the user package's Run(show bool)
		// Note: since Run(true) may itself spawn goroutines that print to stdout,
		// there is no perfect cancellation here. This is the simple capture approach.
		runFunc(true)

		// Wait until all copied output has been processed
		<-copyDone

		app.QueueUpdateDraw(func() {
			fmt.Fprintf(output, "\n[green]Done running %s\n", name)
		})
		doneCleanup()
	}()

	// all set — function is running and output is being streamed
}
