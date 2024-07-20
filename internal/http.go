package internal

import (
	"github.com/imroc/req/v3"
)

type HttpClient struct {
	client *req.Client
}

func NewHttpClient(client *req.Client) *HttpClient {
	return &HttpClient{client: client}
}

func (c *HttpClient) Request(token string) *req.Request {
	return c.client.R().SetHeader("Authorization", token)
}
