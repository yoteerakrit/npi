package api2

import "github.com/yoteerakrit/assessment/external/testapi2"

type GetDataOutput struct {
}

func (c *Client) GetData() (*GetDataOutput, error) {
	resp, err := c.client.NewRequest().Get("")
	if err != nil {
		return nil, err
	}

	output, err := testapi2.UnmarshalOutput(resp, &GetDataOutput{})
	if err != nil {
		return nil, err
	}

	return output, err
}
