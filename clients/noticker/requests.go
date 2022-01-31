package noticker

import (
	"encoding/json"
	"fmt"
	"github.com/knexguy101/CryptAPI/clients/responses"
	"github.com/knexguy101/CryptAPI/errors"
	"io"
	"net/url"
)

func (c *NoTickerClient) Info() (*responses.InfoResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/info/", c.BaseAPI, c.Coin))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.InfoResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *NoTickerClient) Create(address, callback string) (*responses.CreateResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/create/?callback=%s&address=%s", c.BaseAPI, c.Coin, url.QueryEscape(callback), address))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.CreateResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *NoTickerClient) QRCode(address string) (*responses.QRCodeResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/qrcode/?address=%s", c.BaseAPI, c.Coin, address))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.QRCodeResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *NoTickerClient) Estimate() (*responses.EstimateResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/estimate/", c.BaseAPI, c.Coin))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.EstimateResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *NoTickerClient) Convert(value, from string) (*responses.ConvertResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/convert/?value=%s&from=%s", c.BaseAPI, c.Coin, value, from))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.ConvertResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *NoTickerClient) Logs(callback string) (*responses.PaymentLogsResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/logs/?callback=%s", c.BaseAPI, c.Coin, callback))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.PaymentLogsResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}