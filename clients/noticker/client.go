package noticker

import (
	"github.com/knexguy101/CryptAPI/tokens/coins"
	"net/http"
)

type NoTickerClient struct {
	Client *http.Client
	Coin coins.Coin
	BaseAPI string
}

func NewNoTickerClient(coin coins.Coin) *NoTickerClient {
	return &NoTickerClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: true, //saves your internet
			},
		},
		Coin: coin,
		BaseAPI: "https://api.cryptapi.io",
	}
}

func (c *NoTickerClient) get(endpoint string) (*http.Response, error) {
	return c.Client.Get(endpoint)
}

func (c *NoTickerClient) post(endpoint string) (*http.Response, error) {
	return c.Client.Post(endpoint, "application/json", nil)
}