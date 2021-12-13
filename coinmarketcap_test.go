package pricer

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCoinMarketCap(t *testing.T) {
	cmc := NewCoinMarketCap("d438f909-d6c3-4b15-8b32-f3d5e3fe4ed8")
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
			Token:  BSC,
			Target: USD,
		},
	} {
		price, err := cmc.Query(symbol.Token, symbol.Target)
		require.NoError(t, err)
		require.Equal(t, true, decimal.NewFromFloat(price).GreaterThan(decimal.Zero))
	}
}
