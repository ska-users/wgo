package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := ""
		assertStrings(t, got, want)
		assertErrors(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertErrors(t *testing.T, errgot, errwant error) {
	t.Helper()
	if errgot != errwant {
		t.Errorf("got %q want %q given, %q", errgot, errwant, "test")
	}
}
