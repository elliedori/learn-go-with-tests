package set

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreate(t *testing.T) {
	t.Run("creates empty set for empty input", func(t *testing.T) {
		s := NewSet([]int{})
		got := s.Items
		want := []int{}

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("creates set with all values for non-empty input", func(t *testing.T) {

		s := NewSet([]int{1, 2, 3})
		got := s.Items
		want := []int{1, 2, 3}

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestSetSize(t *testing.T) {
	t.Run("returns 0 for empty set", func(t *testing.T) {
		s := NewSet([]int{})
		got := s.Size()
		want := 0

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns numbers of elements for non-empty set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		got := s.Size()
		want := 3

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestSetMembership(t *testing.T) {
	t.Run("returns true if value is in set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		got := s.Includes(2)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("returns false if value is not in set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		got := s.Includes(7)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestAddValue(t *testing.T) {
	t.Run("does not add value if already in set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		s.Add(1)
		got := s.Items
		want := []int{1, 2, 3}

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("does add value if not already in set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		s.Add(4)
		got := s.Items
		want := []int{1, 2, 3, 4}

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestDeleteValue(t *testing.T) {
	t.Run("deletes value if present in set", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		s.Delete(2)
		got := s.Items
		want := []int{1, 3}

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("does not modify set if the value is not present", func(t *testing.T) {
		s := NewSet([]int{1, 2, 3})
		s.Delete(4)
		got := s.Items
		want := []int{1, 2, 3}

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
