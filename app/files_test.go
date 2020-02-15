package app

import (
	"os"
	"testing"
)

const fileName = "./testdata/1.png"

func TestGetBaseNameWithoutExt(t *testing.T) {
	if file, err := os.Open(fileName); err == nil {
		got := GetBaseNameWithoutExt(*file)
		want := "1"
		assert(t, got, want)
	} else {
		t.Error(err)
	}
}

func TestGetExtension(t *testing.T) {
	if file, err := os.Open(fileName); err == nil {
		got := GetExtension(*file)
		want := ".png"
		assert(t, got, want)
	} else {
		t.Error(err)
	}
}

func assert(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
