package api2

import "github.com/yoteerakrit/assessment/external/testapi2"

type IApi interface {
	GetData() (*GetDataOutput, error)
}

type Client struct {
	client testapi2.ClientType
}
