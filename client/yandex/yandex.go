package yandex

import (
	"fmt"
	"io"
	"net/http"
)

const (
	Address       = "https://speller.yandex.net/services/spellservice.json"
	PathCheckText = "checkText?text=%s"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) Do(req *http.Request) ([]byte, error) {
	var resp []byte

	response, err := c.httpClient.Do(req)
	if err != nil {
		return resp, fmt.Errorf("request error: %w", err)
	}

	defer response.Body.Close()

	resp, err = io.ReadAll(response.Body)
	if err != nil {
		return resp, fmt.Errorf("reading request body: %w", err)
	}

	if response.StatusCode < http.StatusOK ||
		response.StatusCode >= http.StatusMultipleChoices {
		return resp, fmt.Errorf("docuforce response body: %s", resp)
	}

	return resp, err
}
