package api

import (
	"context"

	testapi "github.com/yoteerakrit/assessment/external/testapi/test"
)

type IApi interface {
	Get(ctx context.Context) (*GetOutput, error)
}

type Client struct {
	client testapi.ClientType
}
