package responses

import "time"

type InfoResponse struct {
	Coin                   string `json:"coin"`
	Logo                   string `json:"logo"`
	Ticker                 string `json:"ticker"`
	MinimumTransaction     int64  `json:"minimum_transaction"`
	MinimumTransactionCoin string `json:"minimum_transaction_coin"`
	MinimumFee             int64  `json:"minimum_fee"`
	MinimumFeeCoin         string `json:"minimum_fee_coin"`
	FeePercent             string `json:"fee_percent"`
	Status                 string `json:"status"`
	Prices                 struct {
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
	PricesUpdated time.Time `json:"prices_updated"`
}