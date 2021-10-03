package jibble

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	services "github.com/markkizz/scheduling-my-boring-stuff/services/httpclient"
)

type JibbleHttpClient interface {
	services.BaseHttpClient
	GetUserAccessToken(ctx context.Context, request UserAccessTokenRequest) (UserAccessTokenResponse, error)
	GetOrganizationId(ctx context.Context) (*OrganizationIdResponse, error)
	NewJibbleClient(options JibbleClientOptions) *JibbleClient
}

type JibbleClient struct {
	identityUrl     string
	timetrackingUrl string
	token           string
	client          *services.HttpClient
}

type JibbleClientOptions struct {
	identityUrl     string
	timetrackingUrl string
	token           string
}

func (jbc *JibbleClient) NewJibbleClient(options JibbleClientOptions) *JibbleClient {
	return &JibbleClient{
		options.identityUrl,
		options.timetrackingUrl,
		options.token,
		&services.HttpClient{},
	}
}

func (jbc *JibbleClient) request(req *http.Request, target interface{}) error {

	if contentType := req.Header.Get("Content-Type"); contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	if jbc.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jbc.token))
	}

	res := jbc.client.SendRequest(req, target)
	return res
}

func (jbc *JibbleClient) GetUserAccessToken(ctx context.Context, request UserAccessTokenRequest) (*UserAccessTokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", "ro.client")
	data.Set("grant_type", "password")
	data.Set("username", request.username)
	data.Set("password", request.password)

	req, err := http.NewRequest(http.MethodPost, jbc.identityUrl+"/connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := &UserAccessTokenResponse{}
	if err := jbc.request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (jbc *JibbleClient) GetOrganizationId(ctx context.Context) (*OrganizationIdResponse, error) {
	req, err := http.NewRequest(http.MethodGet, jbc.identityUrl+"/v1/Organizations", nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	response := &OrganizationIdResponse{}
	if err := jbc.request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (jbc *JibbleClient) GetPersonId(ctx context.Context, request PersonIdRequest) (*PersonIdResponse, error) {
	query := url.Values{}
	query.Add("$filter", "organizationId eq")
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/v1/People?%s+%s", jbc.identityUrl, query.Encode(), request.organizationId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	response := &PersonIdResponse{}
	if err := jbc.request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (jbc *JibbleClient) GetPersonAccessToken(ctx context.Context, request PersonAccessTokenRequest) (*PersonAccessToken, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("client_id", "ro.client")
	writer.WriteField("grant_type", "password")
	writer.WriteField("refresh_token", request.refreshToken)
	writer.WriteField("acr_values", "prsid:"+request.personId)
	writer.WriteField("username", request.username)
	writer.WriteField("password", request.password)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, jbc.identityUrl+"/connect/token", payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req = req.WithContext(ctx)

	response := &PersonAccessToken{}
	if err := jbc.request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (jbc *JibbleClient) Clocking(ctx context.Context, request ClockingRequest) error {
	uniqId := uuid.NewString()
	payload := strings.NewReader(fmt.Sprintf(`{
    "id": "%v",
    "personId": "%v",
    "type": "%v",
    "clientType": "Web",
    "platform": {}
	}`, uniqId, request.PersonID, request.Type))

	req, err := http.NewRequest(http.MethodPost, jbc.timetrackingUrl+"/v1/TimeEntries", payload)
	// ! io will change request body
	// io.Copy(os.Stdout, req.Body)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	if err := jbc.request(req, nil); err != nil {
		return err
	}

	return nil
}
