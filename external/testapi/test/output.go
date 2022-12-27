package testapi

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func UnmarshalOutput[T any](response *resty.Response, output *T) (*T, error) {
	if err := json.Unmarshal(response.Body(), output); err != nil {
		return nil, err
	}

	return output, nil
}
