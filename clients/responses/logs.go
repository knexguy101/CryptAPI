package responses

type PaymentLogsResponse struct {
	AddressIn           string      `json:"address_in"`
	AddressOut          string      `json:"address_out"`
	CallbackURL         string      `json:"callback_url"`
	Status              string      `json:"status"`
	NotifyPending       bool        `json:"notify_pending"`
	NotifyConfirmations int         `json:"notify_confirmations"`
	Priority            string      `json:"priority"`
	Callbacks           []Callbacks `json:"callbacks"`
}

type Logs struct {
	RequestURL     string `json:"request_url"`
	Response       string `json:"response"`
	ResponseStatus string `json:"response_status"`
	Timestamp      string `json:"timestamp"`
	NextTry        string `json:"next_try"`
	Success        bool   `json:"success"`
}

type Callbacks struct {
	TxidIn             string  `json:"txid_in"`
	TxidOut            string  `json:"txid_out"`
	Value              int64   `json:"value"`
	ValueCoin          float64 `json:"value_coin"`
	ValueForwarded     int64   `json:"value_forwarded"`
	ValueForwardedCoin float64 `json:"value_forwarded_coin"`
	Confirmations      int     `json:"confirmations"`
	LastUpdate         string  `json:"last_update"`
	Result             string  `json:"result"`
	FeePercent         int     `json:"fee_percent"`
	Fee                int64   `json:"fee"`
	Logs               []Logs  `json:"logs"`
}