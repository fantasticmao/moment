package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const testFile = "./testdata/1.png"

func TestOpenFiles(t *testing.T) {
	files, err := openFiles([]string{testFile})
	assert.Nil(t, err)
	defer closeFiles(files)

	assert.True(t, len(files) > 0)

	file := files[0]
	fileName := filepath.Base(file.Name())
	assert.Equal(t, "1.png", fileName)
}

func TestGetFileName(t *testing.T) {
	file, err := os.Open(testFile)
	assert.Nil(t, err)
	defer file.Close()

	actual := getFileName(file)
	assert.Equal(t, "1", actual)
}

func TestGetFileExtension(t *testing.T) {
	file, err := os.Open(testFile)
	assert.Nil(t, err)
	defer file.Close()

	actual := getFileExtension(file)
	assert.Equal(t, ".png", actual)
}
