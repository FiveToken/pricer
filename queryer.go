package pricer

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"log"
	"strconv"
	"time"
)

const (
	ETH Token = "eth"
	FIL Token = "fil"
	USD Token = "usd"
)

type Token string

func (t Token) String() string {
	return string(t)
}

type QueryService interface {
	Query(token, target Token) (price float64, err error)
}

type Symbol struct {
	Token  Token
	Target Token
}

func (s Symbol) CacheKey() string {
	return fmt.Sprintf("pricer_%s_%s", s.Token, s.Target)
}

func NewQueryer(duration time.Duration, redis *redis.Client, symbols []Symbol, services []QueryService, ) *Queryer {
	return &Queryer{symbols: symbols, duration: duration, services: services, redis: redis}
}

type Queryer struct {
	symbols  []Symbol
	duration time.Duration
	services []QueryService
	redis    *redis.Client
}

func (q Queryer) Run() {
	for {
		for _, symbol := range q.symbols {
			total := decimal.Decimal{}
			count := int64(0)
			for _, service := range q.services {
				price, err := service.Query(symbol.Token, symbol.Target)
				if err != nil {
					continue
				}
				total = total.Add(decimal.NewFromFloat(price))
				count++
			}
			if count == 0 {
				continue
			}
			log.Printf("query %s price done, services count: %d", symbol.Token, count)
			avgPrice := total.Div(decimal.NewFromInt(count))
			q.redis.Set(context.Background(), symbol.CacheKey(), avgPrice.String(), 0)
		}
		log.Printf("price queryer sleep: %s", q.duration)
		time.Sleep(q.duration)
	}
}

func (q Queryer) exec(symbol Symbol, service QueryService) (price float64, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("panic error: %v", err)
		}
	}()
	return service.Query(symbol.Token, symbol.Target)
}

func (q Queryer) Get(symbol Symbol) (price float64, err error) {
	res := q.redis.Get(context.Background(), symbol.CacheKey())
	if res.Err() != nil {
		return
	}
	return strconv.ParseFloat(res.Val(), 64)
}
