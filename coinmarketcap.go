package pricer

import (
	"fmt"
	"github.com/valyala/fastjson"
	"gopkg.in/resty.v1"
	"strings"
)

func NewCoinMarketCap(key string) *CoinMarketCap {
	return &CoinMarketCap{key: key}
}

type CoinMarketCap struct {
	key string
}

func (c CoinMarketCap) Query(token, target Token) (price float64, err error) {
	q := resty.New()
	q.RetryCount = 3
	q.SetHeader("X-CMC_PRO_API_KEY", c.key)
	r, err := q.R().Get(
		fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/tools/price-conversion?amount=1&symbol=%s&convert=%s", token, target),
	)
	if err != nil {
		return
	}
	j, err := fastjson.ParseBytes(r.Body())
	if err != nil {
		return
	}
	return j.Get("data").Get("quote").Get(strings.ToUpper(target.String())).Get("price").Float64()
}
