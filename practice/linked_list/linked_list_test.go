package linkedlist

import (
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("Creates a list of length 1", func(t *testing.T) {
		want := 1
		l := CreateLinkedList("a")
		got := l.Length

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("Adds onto the end of the list", func(t *testing.T) {
		want := 2
		l := CreateLinkedList("a")
		l.Insert("b")
		got := l.Length
		l.Display()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Deletes node from list", func(t *testing.T) {
		want := 3
		l := CreateLinkedList("a")
		l.Insert("b")
		l.Insert("c")
		l.Insert("d")
		l.Delete("d")
		got := l.Length
		l.Display()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
