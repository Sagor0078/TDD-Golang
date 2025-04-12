

package Select

import (
	"testing"
	"time"
)

func TestFastestResponce(t *testing.T) {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.second)
		ch1 <- "slow"
	}()

	go fund() {
		time.Sleep(1 * time.second)
		ch2 <- "fast"
	}()


	got := FastestResponce(ch1, ch2)
	want := "fast"

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}