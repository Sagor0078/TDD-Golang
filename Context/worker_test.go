
package Context

import (
	"context"
	"testing"
	"time"
)

func TestDomeSomething_Cancelled(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	result := make(chan string)
	go DoSomething(ctx, result)

	if res := <- result; res != "cancelled" {
		t.Errorf("Expected cancelled, got %s", res)
	}
}