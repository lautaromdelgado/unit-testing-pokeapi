package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	c := require.New(t)
	result := Add(20, 2)
	expected := 22
	c.Equal(expected, result, "se esperaba %d pero se obtuvo %d", expected, result)
}
