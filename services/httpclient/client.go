package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"

	// "io"
	"net/http"
	// "os"
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
	// Print the body to the stdout
	// io.Copy(os.Stdout, res.Body)

	now := time.Now().Format(time.ANSIC)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		// io.Copy(os.Stdout, res.Body)
		var errRes ErrorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(fmt.Sprintf("[error] %v %s %v - %v %v\n", req.Method, req.URL, now, res.Status, errRes.Message))
		}
		return fmt.Errorf("[error] %v %s %v - %v %v\n", req.Method, req.URL, now, res.Status, errRes.Message)
	}

	if err = json.NewDecoder(res.Body).Decode(&target); err != nil {
		return err
	}
	fmt.Printf("[request] %v %s %v - %v\n", req.Method, req.URL, now, res.Status)

	return nil
}

func NewClient() *http.Client {
	httpClient := HttpClient{}
	client := httpClient.Create()
	return client
}
