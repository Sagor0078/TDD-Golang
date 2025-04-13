package Context

import (
	"context"
	"fmt"
	"time"
)

func DoSomething(ctx context.Context, result chan<- string) {
	select {
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
		result <- "cancelled"
	case <-simulateWork():
		result <- "completed"
	}
}

func simulateWork() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		// simulate work
		time.Sleep(2 * time.Second)
		close(done)
	}()
	return done
}
