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

		case err := <-ch:
			return err
		}
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
