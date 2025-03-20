package bscscan

import "strings"

func (c *Client) GetAccountBalance(address string) (string, error) {
	resp, err := c.client.R().SetQueryParams(map[string]string{
		"module":  "account",
		"action":  "balance",
		"address": address,
		"apikey":  c.apikey,
	}).Get("/")
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (c *Client) GetAccountsBalance(addresses []string) (string, error) {
	resp, err := c.client.R().SetQueryParams(map[string]string{
		"module":  "account",
		"action":  "balancemulti",
		"address": strings.Join(addresses, ","),
		"apikey":  c.apikey,
	}).Get("/")
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (c *Client) GetBEP20TransferEvents(contractaddress, address string) (string, error) {
	resp, err := c.client.R().SetQueryParams(map[string]string{
		"module":          "account",
		"action":          "tokentx",
		"contractaddress": contractaddress,
		"address":         address,
		"apikey":          c.apikey,
	}).Get("/")
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (c *Client) GetBEP721TransferEvents(contractaddress, address string) (string, error) {
	resp, err := c.client.R().SetQueryParams(map[string]string{
		"module":          "account",
		"contractaddress": contractaddress,
		"action":          "tokennfttx",
		"address":         address,
		"apikey":          c.apikey,
	}).Get("/")
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
