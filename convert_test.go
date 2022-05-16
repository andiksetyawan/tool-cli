package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert_Fail(t *testing.T) {
	err := convert("not-found-file.log", "destination.txt")
	assert.NotNil(t, err)
}

func TestConvert_Success(t *testing.T) {
	err := convert("test-log.log", "test-destination.txt")
	assert.Nil(t, err)
}
