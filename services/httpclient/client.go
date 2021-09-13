package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HttpClient http.Client

type Options struct {
	ApiUrl string
}

func (ht *HttpClient) Create() *http.Client {
	timeout := time.Second * 10
	client := &http.Client{
		Timeout: timeout,
	}
	return client
}

func (ht *HttpClient) SendRequest(req *http.Request, target interface{}) error {
	client := ht.Create()
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(fmt.Sprintf("error statusCode: %d - %s", res.StatusCode, errRes.Message))
		}
		now := time.Now().Format(time.ANSIC)
		return fmt.Errorf("error statusCode: %d - %s", res.StatusCode, now)
	}

	if err = json.NewDecoder(res.Body).Decode(&target); err != nil {
		return err
	}

	return nil
}

func NewClient() *http.Client {
	httpClient := HttpClient{}
	client := httpClient.Create()
	return client
}
