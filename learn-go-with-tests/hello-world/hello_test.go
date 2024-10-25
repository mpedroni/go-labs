package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bruce", "")
		want := "Hello, Bruce"

		assert(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assert(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Pablo", spanish)
		want := "Hola, Pablo"

		assert(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean-Jacques", french)
		want := "Bonjour, Jean-Jacques"

		assert(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Vladimir", "Russian")
		want := "здравствуйте, Vladimir"

		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	// it tells the test suite to print the func caller line number rather than the line number inside this func if some test fails
	t.Helper()
	if got != want {
		// %q print the string double quoted
		t.Errorf("got %q want %q", got, want)
	}
}
