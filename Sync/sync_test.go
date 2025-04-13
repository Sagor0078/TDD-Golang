
package Sync

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := NewSafeCounter()

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			c.Inc("example")
			wg.Done()
		}()
	}

	wg.Wait()

	if got := c.Value("example"); got != 1000 {
		t.Errorf("Expected 1000, got %d", got)
	} 
}