package testapi2

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type OptionFunc = func(c *resty.Client)

func setOptions(cli *resty.Client, options []OptionFunc) {
	for _, v := range options {
		v(cli)
	}
}

func WithTimeout(time time.Duration) OptionFunc {
	return func(c *resty.Client) {
		c.SetTimeout(time)
	}
}

func WithRetryCount(count int) OptionFunc {
	return func(c *resty.Client) {
		c.SetRetryCount(count)
	}
}
