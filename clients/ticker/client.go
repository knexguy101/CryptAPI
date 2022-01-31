package ticker

import (
	"github.com/knexguy101/CryptAPI/tokens/coins"
	"github.com/knexguy101/CryptAPI/tokens/tickers"
	"net/http"
)

type TickerClient struct {
	Client *http.Client
	Ticker tickers.Ticker
	Coin coins.Coin
	BaseAPI string
}

func NewTickerClient(coin coins.Coin, ticker tickers.Ticker) *TickerClient {
	return &TickerClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: true, //saves your internet
			},
		},
		Ticker: ticker,
		Coin: coin,
		BaseAPI: "https://api.cryptapi.io",
	}
}

func (c *TickerClient) get(endpoint string) (*http.Response, error) {
	return c.Client.Get(endpoint)
}

func (c *TickerClient) post(endpoint string) (*http.Response, error) {
	return c.Client.Post(endpoint, "application/json", nil)
}