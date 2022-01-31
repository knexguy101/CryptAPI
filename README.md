# CryptAPI
A third-party wrapper for CryptAPI written in golang.

### Ethereum Example
```go
package main

import (
	"fmt"
	"github.com/knexguy101/CryptAPI/clients/noticker"
	"github.com/knexguy101/CryptAPI/tokens/coins"
)

func main() {
	client := noticker.NewNoTickerClient(coins.ETH)

	info, err := client.Info()
	if err != nil {
		fmt.Println(err.CryptError) //The inner error that can be checked against defined error types
		fmt.Println(err.InnerError) //The raw golang error
	}

	fmt.Println(info.Status)
	fmt.Println(info.Coin)
	fmt.Println(info.FeePercent)
}
```
### ERC20 Ticker Example
```go
package main

import (
	"fmt"
	"github.com/knexguy101/CryptAPI/clients/ticker"
	"github.com/knexguy101/CryptAPI/tokens/coins"
	"github.com/knexguy101/CryptAPI/tokens/tickers"
	"github.com/knexguy101/CryptAPI/tokens/tickers/erc20"
)

func main() {
	client := ticker.NewTickerClient(coins.ERC20, tickers.Ticker(erc20.USDC))

	info, err := client.Info()
	if err != nil {
		fmt.Println(err.CryptError) //The inner error that can be checked against defined error types
		fmt.Println(err.InnerError) //The raw golang error
	}

	fmt.Println(info.Status)
	fmt.Println(info.Coin)
	fmt.Println(info.FeePercent)
}
```
