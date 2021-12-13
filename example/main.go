package main

import (
	"context"
	"github.com/FiveToken/pricer"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	
	symbols := []pricer.Symbol{
		{Token: pricer.ETH, Target: pricer.USD},
		{Token: pricer.FIL, Target: pricer.USD},
	}
	
	services := []pricer.QueryService{
		pricer.NewCoinGecko(),
		pricer.NewCoinMarketCap("d438f909-d6c3-4b15-8b32-f3d5e3fe4ed8"),
	}
	
	redisClient := redis.NewClient(&redis.Options{
		Addr: "192.168.1.161:6379",
		DB:   0,
	})
	
	res := redisClient.Ping(context.Background())
	if res.Err() != nil {
		log.Panicf("ping redis eror: %s", res.Err())
	}
	queryer := pricer.NewQueryer(
		1*time.Minute,
		redisClient,
		symbols,
		services,
	)
	go queryer.Run()
	for {
		for _, symbol := range symbols {
			ethPrice, err := queryer.Get(symbol)
			if err != nil {
				log.Panicf("get %s price error: %s", symbol.Token, err)
			}
			log.Printf("%s price: %f", symbol.Token, ethPrice)
		}
		time.Sleep(30 * time.Second)
	}
}
