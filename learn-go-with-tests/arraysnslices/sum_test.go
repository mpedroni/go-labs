package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection with 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum 2 slices", func(t *testing.T) {
		a, b := []int{1, 2, 3}, []int{10, 20, 30}
		got := SumAll(a, b)
		want := []int{6, 60}

		assertSum(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum some slice tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}

		assertSum(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertSum(t, got, want)
	})
}

func assertSum(t *testing.T, got []int, want []int) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
