package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()

}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}

		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {

}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			fmt.Println("error occured:", err)
		}
		fmt.Fprint(w, data)
	}
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
// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		data := make(chan string, 1)

// 		go func() {
// 			data <- store.Fetch()
// 		}()

// 		select {
// 		case d := <-data:
// 			fmt.Fprint(w, d)
// 		case <-ctx.Done():
// 			store.Cancel()
// 		}
// 	}
// }
