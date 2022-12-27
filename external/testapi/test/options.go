package testapi

import "github.com/go-resty/resty/v2"

type OptionFunc = func(c *resty.Client)

func setOptions(req *resty.Client, options []OptionFunc) {
	for _, v := range options {
		v(req)
	}
}
