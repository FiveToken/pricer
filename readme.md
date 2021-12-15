# FiveToken Pricer

Service to obtain the market price of Token

## Install

```
go get -u gihub.com/FiveToken/pricer
```

## Feature

* Based on Redis
* Support multi-data sources
* Based on Interface

## Multi-currency support

* FIL
* ETH
* BSC

## Support service provider

* CoinGecko
* CoinMarketCap

## Price acquisition strategy 
easier price strategy:

1. The average value is used when multi-data sources are available simultaneously;
2. The average value is calculated based on the price data of the available service; 
3. The price cache of the last time node When all services are unavailable.

## Usage

**Run price fetching service**

```go
package main

import "fmt"
import "github.com/FiveToken/pricer"

var queryer *pricer.Queryer

func init(){
  symbols := []pricer.Symbol{
		{Token: pricer.ETH, Target: pricer.USD},
		{Token: pricer.FIL, Target: pricer.USD},
		{Token: pricer.BSC, Target: pricer.USD},
	}  
  services := []pricer.QueryService{
      pricer.NewCoinGecko(),
      pricer.NewCoinMarketCap("key"),
    }

  queryer = pricer.NewQueryer(
      10*time.Minute, 
      redisClient,
      symbols,
      services,
    )

   go queryer.Run()	
}
```

**Get price**
```go
func main(){
  fmt.Println(  
    queryer.Get(pricer.Sympol{ Token: pricer.FIL, Target: pricer.USD })
  )
}
```



