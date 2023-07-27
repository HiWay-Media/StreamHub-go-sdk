package streamhub

import (
	"github.com/go-resty/resty/v2"
)

type streamHubSdk struct {
	url        string
	restClient *resty.Client
	debug      bool
}

type IStreamHubClient interface {
	HealthCheck() error
	IsDebug() bool
}

func (o *streamHubSdk) IsDebug() bool {
	return o.debug
}

func (o *streamHubSdk) HealthCheck() error {
	_, err := o.restyGet(o.url, nil)
	if err != nil {
		return nil
	}
	return nil
}
