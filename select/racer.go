package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// configure racer with how long to finish the race
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// synchronise processes with select
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		// return error after request times out
		return "", fmt.Errorf("timed out after waiting %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// channel of type struct require no memory allocation
	// making the process abit faster
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// close channel after getting response
		close(ch)
	}()
	return ch
}
