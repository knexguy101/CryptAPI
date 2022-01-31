package responses

type CreateResponse struct {
	AddressIn   string `json:"address_in"`
	AddressOut  string `json:"address_out"`
	CallbackURL string `json:"callback_url"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}
