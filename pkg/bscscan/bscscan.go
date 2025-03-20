package bscscan

import (
	"resty.dev/v3"
)

type Client struct {
	client *resty.Client
	apikey string
}

func NewClient(apikey string) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL("https://api.bscscan.com/api").
			SetQueryParam("apikey", apikey).SetDebug(true),
		apikey: apikey,
	}
}
