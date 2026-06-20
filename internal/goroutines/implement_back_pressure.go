package goroutines

import (
	"fmt"
	"net/http"
)

// ImplementBackPressure demonstrates implementing back pressure with goroutines
func ImplementBackPressure(show bool) {
	if !show {
		return
	}
	fmt.Println("-- Implement Back Pressure")
	fmt.Println("\nThis example demonstrates back pressure using a PressureGauge.")
	fmt.Println("In a real application, this would:")
	fmt.Println("1. Create a PressureGauge with a limit (e.g., 10 concurrent requests)")
	fmt.Println("2. Use it to limit concurrent processing in an HTTP handler")
	fmt.Println("3. Return 'Too Many Requests' (429) when the limit is exceeded")
	fmt.Println("\nExample code structure:")
	fmt.Println("  pg := New(10)")
	fmt.Println("  err := pg.Process(func() { /* work here */ })")
	fmt.Println("  if err != nil {")
	fmt.Println("    // Handle too many concurrent requests")
	fmt.Println("  }")
	fmt.Println("\nNote: This example requires an HTTP server, which cannot run in this TUI.")
}

func implementBackPressure() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
