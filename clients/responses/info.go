package responses

import "time"

type InfoResponse struct {
	Coin               string    `json:"coin"`
	MinimumTransaction string    `json:"minimum_transaction"`
	MinimumFee         string    `json:"minimum_fee"`
	FeePercent         string    `json:"fee_percent"`
	PricesUpdated      time.Time `json:"prices_updated"`
	Status             string    `json:"status"`
	Prices             string    `json:"prices"`
}
