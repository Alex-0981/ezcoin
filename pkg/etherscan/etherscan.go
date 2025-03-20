package etherscan

import "resty.dev/v3"

type Client struct {
	client *resty.Client
	apikey string
}

func NewClient(apikey string) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL("https://api.etherscan.io").
			SetQueryParams(map[string]string{
				"chainid": "1",
				"apikey":  apikey,
			}),
		apikey: apikey,
	}
}
