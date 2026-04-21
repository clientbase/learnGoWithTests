package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		dictionary.Add("test", "this is just a test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "this is just a test"

		err := dict.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, def)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("exisiting word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		newDef := "new def"

		dict.Update(word, newDef)

		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, def)

		assertError(t, err, ErrWordDoesNotExist)

	})

}

func TestDelete(t *testing.T) {
	word := "word"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)

	t.Run("non-existant wor", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "test def"}

		err := dict.Delete(word)

		assertError(t, err, nil)

		_, err = dict.Search(word)

		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, def string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should have found word:", err)
	}
	assertStrings(t, got, def)
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
		t.Errorf("got error %q want %q", got, want)
	}
}
