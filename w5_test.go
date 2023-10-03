package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Get_Price_W5(t *testing.T) {
	p := 126
	str := GetPrice(p)

	require.Equal(t, str, "Wrong age. Try again")

	p = 4
	str = GetPrice(p)

	require.Equal(t, str, "Free")
	p = 17
	str = GetPrice(p)

	require.Equal(t, str, "30000")
	p = 29
	str = GetPrice(p)

	require.Equal(t, str, "50000")

	p = 71
	str = GetPrice(p)

	require.Equal(t, str, "30000")
}
