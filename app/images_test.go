package app

import (
	"os"
	"testing"
)

func TestCheckIllegalImage(t *testing.T) {
	if file, err := os.Open(fileName); err == nil {
		defer file.Close()

		result := CheckIllegalImage(*file)
		if !result {
			t.Errorf("check file %s magic number error", fileName)
		}
	} else {
		t.Error(err)
	}
}
