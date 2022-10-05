package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// What select lets you do is wait on multiple channels.
// The first one to send a value "wins"
// and the code underneath the case is executed.
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}

}

// Why struct{} and not another type like a bool? Well,
// a chan struct{} is the smallest data type available
// from a memory perspective so we get no allocation versus a bool.
// Since we are closing and not sending anything on the chan, why allocate anything?
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
