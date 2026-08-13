package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
)

func randomMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := range mat {
		mat[i] = make([]int, cols)
		for j := range mat[i] {
			// Random value: -1, 0, or 1
			mat[i][j] = rand.Intn(3) - 1
		}
	}
	return mat
}

func Matrix(show bool) {
	if show {
		fmt.Println("--- Matrix ---")

		// Just for testing purposes
		n := 2
		m := 3 // common to both matrices
		k := 2

		// Generate matrices
		A := randomMatrix(n, m)
		B := randomMatrix(m, k)

		// Initialize result matrix
		result := make([][]int, n)
		for i := range result {
			result[i] = make([]int, k)
		}

		// Perform matrix multiplication concurrently
		var wg sync.WaitGroup
		for i := range n {
			for j := range k {
				wg.Add(1)
				go func(row, col int) {
					defer wg.Done()
					sum := 0
					for l := range m {
						sum += A[row][l] * B[l][col]
					}
					// Safe: each goroutine writes to a unique cell
					result[row][col] = sum
				}(i, j)
			}
		}
		wg.Wait()

		fmt.Println("Matrix A:")
		for _, row := range A {
			fmt.Println(row)
		}

		fmt.Println("Matrix B:")
		for _, row := range B {
			fmt.Println(row)
		}

		fmt.Println("Result Matrix (A x B):")
		for _, row := range result {
			fmt.Println(row)
		}
	}
}
