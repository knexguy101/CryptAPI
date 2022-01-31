package responses

import "time"

type InfoResponse struct {
	Coin               string    `json:"coin"`
	MinimumTransaction float64    `json:"minimum_transaction"`
	MinimumFee         float64    `json:"minimum_fee"`
	FeePercent         string    `json:"fee_percent"`
	PricesUpdated      time.Time `json:"prices_updated"`
	Status             string    `json:"status"`
	Prices             struct {
		Usd string `json:"USD"`
		Eur string `json:"EUR"`
		Gbp string `json:"GBP"`
		Cad string `json:"CAD"`
		Jpy string `json:"JPY"`
		Aed string `json:"AED"`
		Dkk string `json:"DKK"`
		Brl string `json:"BRL"`
		Cny string `json:"CNY"`
		Hkd string `json:"HKD"`
		Inr string `json:"INR"`
		Mxn string `json:"MXN"`
		Ugx string `json:"UGX"`
		Pln string `json:"PLN"`
		Php string `json:"PHP"`
		Czk string `json:"CZK"`
		Huf string `json:"HUF"`
		Bgn string `json:"BGN"`
		Ron string `json:"RON"`
	} `json:"prices"`
}