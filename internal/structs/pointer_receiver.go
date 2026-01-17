package structs

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// the receiver is being modified by the method
// so we are using a pointer receiver
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// is good practice in the methods definition to be consistent if one is a pointer receiver
// the others methods should be pointer receivers even if they don't change the receiver
func (c *Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func PointerReceiver(show bool) {
	if show {
		var c Counter
		fmt.Println(c.String())

		c.Increment()
		fmt.Println(c.String())
	}
}
