package testapi

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type IClientType interface {
	NewRequest(ctx context.Context) *resty.Request
}

type ClientType struct {
	resty  *resty.Client
	domain string
}

func New(domain string) *ClientType {
	cli := resty.New()

	cli.SetBaseURL(domain)

	return &ClientType{
		resty: cli,
	}
}

func (c *ClientType) NewRequest(ctx context.Context) *resty.Request {
	return c.resty.R()
}
