package main

import (
	"context"
	"time"
)

// Store interface with a Fetch method that includes context support.
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// SpyStore struct for testing, implementing the Store interface.
type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				return // Cancel the operation if context is done.
			default:
				time.Sleep(10 * time.Millisecond) // Simulate work
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err() // Return context cancellation error
	case res := <-data:
		return res, nil // Return the result if operation completed
	}
}
