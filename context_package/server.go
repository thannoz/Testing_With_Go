package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

/*
context has a method Done() which returns a channel which
gets sent a signal when the context is "done" or "cancelled".
We want to listen to that signal and call store.Cancel if
we get it but we want to ignore it if our Store manages to Fetch before it.

To manage this we run Fetch in a goroutine and it
will write the result into a new channel data.
We then use select to effectively race to the two asynchronous
processes and then we either write a response or Cancel.
*/
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
