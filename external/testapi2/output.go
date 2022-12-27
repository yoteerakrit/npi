package testapi2

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func UnmarshalOutput[T any](resp *resty.Response, output *T) (*T, error) {
	if err := json.Unmarshal(resp.Body(), output); err != nil {
		return nil, err
	}

	return output, nil
}
