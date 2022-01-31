package responses

type EstimateResponse struct {
	EstimatedCost    string `json:"estimated_cost"`
	EstimatedCostUsd string `json:"estimated_cost_usd"`
	Status           string `json:"status"`
}