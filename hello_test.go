package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Ellie", "")
		want := "Hello, Ellie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in German", func(t *testing.T) {
		got := Hello("Grendel", "German")
		want := "Hallo, Grendel"
		assertCorrectMessage(t, got, want)
	})
}
