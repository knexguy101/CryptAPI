package ticker

import (
	"cryptAPI/clients/responses"
	"cryptAPI/errors"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

func (c *TickerClient) Info() (*responses.InfoResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/%s/info", c.BaseAPI, c.Coin, c.Ticker))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.InfoResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *TickerClient) Create(address, callback string) (*responses.CreateResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/%s/create?callback=%s&address=%s", c.BaseAPI, c.Coin, c.Ticker, url.QueryEscape(callback), address))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.CreateResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *TickerClient) QRCode(address string) (*responses.QRCodeResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/%s/qrcode?address=%s", c.BaseAPI, c.Coin, c.Ticker, address))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.QRCodeResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *TickerClient) Estimate() (*responses.EstimateResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/%s/estimate", c.BaseAPI, c.Coin, c.Ticker))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.EstimateResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *TickerClient) Convert(value, from string) (*responses.ConvertResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/%s/convert?value=%s&from=%s", c.BaseAPI, c.Coin, c.Ticker, value, from))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.ConvertResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}

func (c *TickerClient) Logs(callback string) (*responses.PaymentLogsResponse, *errors.ErrorData) {
	res, err := c.get(fmt.Sprintf("%s/%s/%s/logs?callback=%s", c.BaseAPI, c.Coin, c.Ticker, callback))
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorRequest, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorReading, err)
	}

	var obj responses.PaymentLogsResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.NewErrorData(errors.ErrorParsing, err)
	}

	return &obj, nil
}