package moment

import (
	"os"
	"testing"
)

const fileName = "./testdata/1.png"

func TestGetBaseNameWithoutExtension(t *testing.T) {
	file, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	got := getBaseNameWithoutExtension(file)
	want := "1"
	assert(t, got, want)
}

func TestGetExtension(t *testing.T) {
	file, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	got := getExtension(file)
	want := ".png"
	assert(t, got, want)
}

func assert(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %v, but want %v", got, want)
	}
}
