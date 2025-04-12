package mocking

import (
    "fmt"
    "io"
    "os"
    "time"
)

const finalWord = "Go!"
const countdownStart = 3

const (
    write = "write"
    sleep = "sleep"
)

// Used for testing order of operations
type CountdownOperations interface {
    Sleep()
    Write(p []byte) (n int, err error)
}

// Spy implementation
type SpyCountdownOperations struct {
    Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

// General sleeper interface
type Sleeper interface {
    Sleep()
}

// Configurable sleep implementation
type ConfigurableSleeper struct {
    Duration time.Duration
    SleepFn  func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
    c.SleepFn(c.Duration)
}

// The actual countdown logic
func Countdown(w io.Writer, s Sleeper) {
    for i := countdownStart; i > 0; i-- {
        s.Sleep()
        fmt.Fprintln(w, i)
    }
    s.Sleep() // sleep before Go!
    fmt.Fprint(w, finalWord)
}

// For manual running
func main() {
    sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}
