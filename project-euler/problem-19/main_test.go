package main

import (
	"testing"

	"github.com/a1gemmel/project-euler/golang/lib"
	"github.com/stretchr/testify/assert"
)

func Test_Days(t *testing.T) {
	assert.Equal(t, 365, lib.SumInts(daysByMonth))
}
