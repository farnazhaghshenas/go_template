package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var MarshalFailed = errors.New("failed to marshal body of request")

func RestCall(value interface{}, method string, url string, body map[string]interface{}, headers map[string]string) error {
	bs, err := json.Marshal(body)
	if err != nil {
		return MarshalFailed
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", " application/json")
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	bs, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &value)
	if err != nil {
		return err
	}
	return nil
}

