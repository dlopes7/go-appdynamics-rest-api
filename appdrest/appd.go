package appdrest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// APIError to get HTTP response code to expected errors
type APIError struct {
	Message string
	Code    int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

// Client manages communication with AppDynamics
type Client struct {
	client     *http.Client
	Controller *Controller

	//Shared between different APIs
	common service

	Account             *AccountService
	Application         *ApplicationService
	BusinessTransaction *BusinessTransactionService
	MetricData          *MetricDataService
	Snapshot            *SnapshotService
	Tier                *TierService
}

type service struct {
	client *Client
}

// NewClient Returns a Client, this is needed for any communication
func NewClient(protocol string, controllerHost string, port int, username string, password string, account string) *Client {

	httpClient := http.DefaultClient
	baseURL, err := url.Parse(fmt.Sprintf("%s://%s:%d/controller/", protocol, controllerHost, port))
	if err != nil {
		panic(err.Error())
	}

	restUserName := fmt.Sprintf("%s@%s", username, account)
	controller := &Controller{
		Protocol: protocol,
		Host:     controllerHost,
		Port:     port,
		User:     restUserName,
		Password: password,
		Account:  account,
		BaseURL:  baseURL,
	}

	c := &Client{client: httpClient,
		Controller: controller,
	}

	c.common.client = c

	c.Account = (*AccountService)(&c.common)
	c.Application = (*ApplicationService)(&c.common)
	c.BusinessTransaction = (*BusinessTransactionService)(&c.common)
	c.MetricData = (*MetricDataService)(&c.common)
	c.Snapshot = (*SnapshotService)(&c.common)
	c.Tier = (*TierService)(&c.common)

	return c
}

// NewRequest performs a request.
// The baseURL on the client will be concatenated with the url argument
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := c.Controller.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Controller.User, c.Controller.Password)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do makes the http request
func (c *Client) Do(req *http.Request, v interface{}) error {

	req.URL.RawQuery = req.URL.Query().Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		err := &APIError{
			Code:    resp.StatusCode,
			Message: fmt.Sprintf("Status Code Error: %d\nRequest: %v", resp.StatusCode, req),
		}
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return nil

}
