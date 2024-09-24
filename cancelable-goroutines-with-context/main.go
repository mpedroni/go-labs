package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func verySlowRemoteCall() error {
	time.Sleep(2 * time.Second)
	return nil
}

func veryFastRemoteCall() error {
	time.Sleep(50 * time.Millisecond)
	return nil
}

func cancelable(ctx context.Context, f func() error) error {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	ch := make(chan error)
	go func() {
		ch <- f()
	}()

	for {
		select {
		case <-ctx.Done():
			return errors.New("function too slow")

		case <-ch:
			return nil
		}
	}
}

func main() {
	ctx := context.Background()
	start := time.Now()
	slow := cancelable(ctx, verySlowRemoteCall)
	fast := cancelable(ctx, veryFastRemoteCall)

	fmt.Println("time spent: ", time.Since(start))
	fmt.Println("very slow func: ", slow)
	fmt.Println("very fast func: ", fast)
}
