package responses

type ConvertResponse struct  {
	ValueCoin    string `json:"value_coin"`
	ExchangeRate string `json:"exchange_rate"`
	Status       string `json:"status"`
}