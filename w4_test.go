package kiemthu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Get_Price(t *testing.T) {
	p := -1
	str := GetPrice(p)

	require.Equal(t, str, "Wrong age. Try again")

	p = 0
	str = GetPrice(p)

	require.Equal(t, str, "Free")
	p = 5
	str = GetPrice(p)

	require.Equal(t, str, "Free")
	p = 9
	str = GetPrice(p)

	require.Equal(t, str, "Free")

	p = 10
	str = GetPrice(p)

	require.Equal(t, str, "30000")
	p = 15
	str = GetPrice(p)

	require.Equal(t, str, "30000")
	p = 19
	str = GetPrice(p)

	require.Equal(t, str, "30000")

	p = 20
	str = GetPrice(p)

	require.Equal(t, str, "50000")

	p = 40
	str = GetPrice(p)

	require.Equal(t, str, "50000")

	p = 59
	str = GetPrice(p)

	require.Equal(t, str, "50000")
	p = 60
	str = GetPrice(p)

	require.Equal(t, str, "30000")
	p = 90
	str = GetPrice(p)

	require.Equal(t, str, "30000")
	p = 120
	str = GetPrice(p)

	require.Equal(t, str, "30000")
	p = 121
	str = GetPrice(p)

	require.Equal(t, str, "Wrong age. Try again")

	p = 1000
	str = GetPrice(p)

	require.Equal(t, str, "Wrong age. Try again")

	p = 6
	str = GetPrice(p)

	require.Equal(t, str, "Free")

	p = 10
	str = GetPrice(p)

	require.Equal(t, str, "30000")

	p = 23
	str = GetPrice(p)

	require.Equal(t, str, "50000")

	p = 119
	str = GetPrice(p)

	require.Equal(t, str, "30000")
}
