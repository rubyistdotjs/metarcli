package checkwxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	http    *http.Client
	baseUrl string
	apiKey  string
	ctx     context.Context
}

func New(ctx context.Context, apiKey string) *Client {
	return &Client{
		http:    &http.Client{},
		baseUrl: "https://api.checkwx.com/",
		apiKey:  apiKey,
		ctx:     ctx,
	}
}

func (c Client) get(path string, icaoCodes []string, result interface{}) error {
	params := strings.Join(icaoCodes, ",")
	reqUrl, err := url.JoinPath(c.baseUrl, path, params)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	decodeErr := json.NewDecoder(res.Body).Decode(&result)
	if decodeErr != nil {
		return decodeErr
	}

	return nil
}
