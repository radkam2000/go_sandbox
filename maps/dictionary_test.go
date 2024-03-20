package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		assertError(t, err, ErrNotFound)
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, "new value")
		assertError(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "base value"
		dictionary.Add(key, value)
		newValue := "updated value"
		dictionary.Update(key, newValue)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Update(key, value)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test value"}
	dictionary.Delete(key)
	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()

	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, value)
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
