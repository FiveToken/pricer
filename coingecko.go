package pricer

import (
	"fmt"
	"github.com/valyala/fastjson"
	"gopkg.in/resty.v1"
)

func NewCoinGecko() *CoinGecko {
	return &CoinGecko{}
}

type CoinGecko struct {
}

func (c CoinGecko) Query(token, target Token) (price float64, err error) {
	switch token {
	case FIL:
		token = "filecoin"
	case ETH:
		token = "ethereum"
	}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", token, target)
	q := resty.New()
	q.RetryCount = 3
	r, err := q.R().Get(url)
	if err != nil {
		return
	}
	j, err := fastjson.ParseBytes(r.Body())
	if err != nil {
		return
	}
	return j.Get(token.String()).Get(target.String()).Float64()
}
