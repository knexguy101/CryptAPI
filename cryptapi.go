package cryptAPI

import "net/http"

type CryptAPIClient struct {
	Client *http.Transport
}

