package pricer

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCoinGecko(t *testing.T) {
	cg := NewCoinGecko()
	for _, symbol := range []Symbol{
		{
			Token:  ETH,
			Target: USD,
		},
		{
			Token:  FIL,
			Target: USD,
		},
		{
			Token:  FIL,
			Target: USD,
		},
	} {
		price, err := cg.Query(symbol.Token, symbol.Target)
		require.NoError(t, err)
		require.Equal(t, true, decimal.NewFromFloat(price).GreaterThan(decimal.Zero))
	}
}
