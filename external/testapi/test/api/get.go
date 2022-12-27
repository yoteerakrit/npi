package api

import (
	"context"

	testapi "github.com/yoteerakrit/assessment/external/testapi/test"
)

type GetOutput struct {
}

func (c *Client) Get(ctx context.Context) (*GetOutput, error) {
	resp, err := c.client.NewRequest(ctx).Get("/get")
	if err != nil {
		return nil, err
	}

	output, err := testapi.UnmarshalOutput(resp, &GetOutput{})
	if err != nil {
		return nil, err
	}

	return output, nil
}
