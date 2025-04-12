

package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search existing word", func(t *testing.T) {
		d := Dictionary{"hello": "world"}
		got, _ := d.Search("hello")
		want := "world"
		assertStrings(t, got, want)
	})

	t.Run("Search unknown word", func(t *testing.T) {
		d := Dictionary{}
		_, err := d.Search("unknown")
		assertError(t, err, ErrNotFound)
	})

	t.Run("Add new word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Add("hello", "world")
		assertError(t, err, nil)
		assertDefinition(t, d, "hello", "world")
	})

	t.Run("Add existing word", func(t *testing.T) {
		d := Dictionary{"hello": "world"}
		err := d.Add("hello", "new")
		assertError(t, err, ErrWordExists)
	})

	t.Run("Update existing word", func(t *testing.T) {
		d := Dictionary{"hello": "world"}
		err := d.Update("hello", "updated")
		assertError(t, err, nil)
		assertDefinition(t, d, "hello", "updated")
	})

	t.Run("Update non-existing word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Update("unknown", "value")
		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("Delete existing word", func(t *testing.T) {
		d := Dictionary{"hello": "world"}
		err := d.Delete("hello")
		assertError(t, err, nil)
		_, err = d.Search("hello")
		assertError(t, err, ErrNotFound)
	})

	t.Run("Delete non-existing word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Delete("missing")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

// Helpers

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

