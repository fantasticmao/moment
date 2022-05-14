package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDecodeImage(t *testing.T) {
	file, err := os.Open(testFile)
	assert.Nil(t, err)
	defer file.Close()

	img, imgWidth, imgHeight, imgFormat, err := decodeImage(file)
	assert.Nil(t, err)
	assert.NotNil(t, img)
	assert.Equal(t, 1334, imgWidth)
	assert.Equal(t, 750, imgHeight)
	assert.Equal(t, "png", imgFormat)
}
