package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}
	got := dict.Search("test")
	want := "this is just a test"
	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
