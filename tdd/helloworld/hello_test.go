package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "sp")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Elodie", "ja")
		want := "こんにちは、 Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("no name", func(t *testing.T) {
		got := Hello("", "ja")
		want := "こんにちは、 World"
		assertCorrectMessage(t, got, want)
	})
}
