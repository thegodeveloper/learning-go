package goroutines

func countTo2(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}
