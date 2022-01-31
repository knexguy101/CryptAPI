package responses

type QRCodeResponse struct {
	QrCode     string `json:"qr_code"`
	PaymentURI string `json:"payment_uri"`
	Status     string `json:"status"`
}