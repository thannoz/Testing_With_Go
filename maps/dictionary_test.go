package main

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dic.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dic.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		err := dic.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, definition)
	})

}

func assertDefinition(t testing.TB, dic Dictionary, word, definition string) {
	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dic.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{}

		err := dic.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dictionary{word: "test definition"}

	dic.Delete(word)

	_, err := dic.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q, to be delete", word)
	}
}
