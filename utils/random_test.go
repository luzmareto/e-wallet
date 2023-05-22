package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(3, 100)
	require.GreaterOrEqual(t, randomInt, int64(3))
	require.LessOrEqual(t, randomInt, int64(100))
}

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.Equal(t, 5, len(owner))
}

func TestRandomMoney(t *testing.T) {
	randomInt := RandomMoney()
	require.GreaterOrEqual(t, randomInt, int64(1000))
	require.LessOrEqual(t, randomInt, int64(100000))
}

func TestRandomCurrency(t *testing.T) {
	availableCurrencies := []string{"RUB", "USD", "CAD", "EUR"}
	curency := RandomCurrency()

	require.Equal(t, 3, len(curency))
	require.Contains(t, availableCurrencies, curency)
}

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	require.Contains(t, email, "@email.com")
}
