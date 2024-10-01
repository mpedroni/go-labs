package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func slowRemoteCall() error {
	time.Sleep(2 * time.Second)
	return nil
}

func fastRemoteCall() error {
	time.Sleep(50 * time.Millisecond)
	return nil
}

func fastAndBrokenRemoteCall() error {
	time.Sleep(20 * time.Millisecond)
	return errors.New("fast, but broken")
}

// cancelable act as circuit breaker for the function f, that is, it will wait for the function to finish, but if it takes too long, it will return an error.
func cancelable(ctx context.Context, f func() error) error {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	ch := make(chan error)
	go func() {
		// ATTENTION: once the function begins executing, it will not be interrupted, even if the context is canceled.
		// It is ok for short-lived functions such as HTTP calls, but not for long-running ones, specially it it demands some resource releasing or shutdown gracefully.
		// Furthermore, in a real scenario, the f function receive a context as parameter in order to be able handle cancellation properly if necessary.
		ch <- f()
		close(ch)
	}()

	// select will wait for the first channel to return a value, either the function result or the context.Done signal.
	select {
	case <-ctx.Done():
		return ctx.Err()

	case err := <-ch:
		return err
	}
}

func main() {
	ctx := context.Background()
	start := time.Now()

	slow := cancelable(ctx, slowRemoteCall)
	fast := cancelable(ctx, fastRemoteCall)
	broken := cancelable(ctx, fastAndBrokenRemoteCall)

	fmt.Println("total time spent: ", time.Since(start))
	fmt.Println("slow func: ", slow)
	fmt.Println("fast func: ", fast)
	fmt.Println("broken func: ", broken)
}
