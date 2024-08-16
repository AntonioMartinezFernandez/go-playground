package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldBeWhite(t *testing.T) {
	color := "white"
	assert.Equal(t, color, "white")
}
