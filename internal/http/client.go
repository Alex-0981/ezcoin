package http

import "resty.dev/v3"

func NewClient() *resty.Client {
	return resty.New()
}
