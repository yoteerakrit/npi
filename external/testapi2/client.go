package testapi2

import (
	"github.com/go-resty/resty/v2"
)

type ClientType struct {
	rest *resty.Client
}

func New(domain string, options ...OptionFunc) *ClientType {
	cli := resty.New()

	setOptions(cli, options)

	cli.SetBaseURL(domain)

	return &ClientType{
		rest: cli,
	}
}

func (c *ClientType) NewRequest() *resty.Request {
	return c.rest.R()
}
