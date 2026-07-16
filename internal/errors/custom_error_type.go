package errors

import "fmt"

// DivideError Custom error type for division errors
type DivideError struct {
	A, B float64
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide: %.2f by %.2f", e.A, e.B)
}

// divide performs division and demonstrates basic error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivideError{A: a, B: b}
	}
	return a / b, nil
}

func CustomErrorType(show bool) {
	if show {
		fmt.Println("--- Custom Error Type ---")

		total, err := divide(7.0, 0)
		if err != nil {
			fmt.Println("divide error:", err)
			return
		}

		fmt.Println("total:", total)
	}

}
