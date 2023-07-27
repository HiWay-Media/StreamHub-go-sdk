package streamhub

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func NewStreamHubSdk(url string, debug bool) (IStreamHubClient, error){
	client := &streamHubSdk{
		url:        url,
		restClient: resty.New(),
		debug:      debug,
	}
	//
	if debug {
		client.restClient.SetDebug(true)
		client.debug = true
		log.Println("Debug mode is enabled for the stream hub client ")
	}
	return client, nil
}


// Resty Methods

func (o *streamHubSdk) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// get request
func (o *streamHubSdk) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}
