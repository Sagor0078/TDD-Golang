
package Select_conn

import (
    "testing"
    "time"
)

func TestFastestResponse(t *testing.T) {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "slow"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "fast"
    }()

    got := FastestResponse(ch1, ch2)
    want := "fast"

    if got != want {
        t.Errorf("expected %q but got %q", want, got)
    }
}
