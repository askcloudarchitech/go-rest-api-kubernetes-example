package main

import "testing"

func TestAddStringNumbersTestHappyPath(t *testing.T) {

	got, _ := addStringsAsNumbers("1", "2")
	want := 3

	if got != want {
		t.Errorf("got %x, wanted %x", got, want)
	}

}

func TestLettersExpectError(t *testing.T) {
	got, err := addStringsAsNumbers("a", "2")

	if err == nil {
		t.Errorf("got %x, but expected error", got)
	}
}
