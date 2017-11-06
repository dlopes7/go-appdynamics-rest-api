package appdrest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/op/go-logging"
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

	log *logging.Logger

	Account             *AccountService
	Application         *ApplicationService
	BusinessTransaction *BusinessTransactionService
	Dashboard           *DashboardService
	MetricData          *MetricDataService
	Snapshot            *SnapshotService
	Tier                *TierService
}

type service struct {
	client *Client
}

var log = logging.MustGetLogger("appdrest")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{level:.4s} %{shortfunc} %{id:03x}%{color:reset} %{message}`,
)

// NewClient Returns a Client, this is needed for any communication
func NewClient(protocol string, controllerHost string, port int, username string, password string, account string) *Client {

	httpClient := http.DefaultClient
	baseURL, err := url.Parse(fmt.Sprintf("%s://%s:%d/", protocol, controllerHost, port))
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

	backend1 := logging.NewLogBackend(os.Stdout, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.DEBUG, "")

	logging.SetBackend(backend1Leveled)

	c.log = log
	c.common.client = c

	c.Account = (*AccountService)(&c.common)
	c.Application = (*ApplicationService)(&c.common)
	c.BusinessTransaction = (*BusinessTransactionService)(&c.common)
	c.MetricData = (*MetricDataService)(&c.common)
	c.Snapshot = (*SnapshotService)(&c.common)
	c.Tier = (*TierService)(&c.common)
	c.Dashboard = (*DashboardService)(&c.common)

	c.log.Debug("Created client successfully")
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

	c.log.Debugf("Creating request %s %s", method, url.String())
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
	c.log.Debugf("Performed request %v - HTTP %d", req.URL, resp.StatusCode)

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

// DoRestUI makes the http request using authentication. This will break with different versions of AppDynamics
func (c *Client) DoRestUI(req *http.Request, v interface{}) error {

	req.URL.RawQuery = req.URL.Query().Encode()

	if len(req.Header["X-CSRF-TOKEN"]) == 0 {
		c.log.Debugf("RESTUI, logging in...")
		err := c.login(req)
		if err != nil {
			panic(err.Error())
		}

	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	c.log.Debugf("Performed request %v - HTTP %d", req.URL, resp.StatusCode)

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

func (c *Client) login(req *http.Request) error {

	url := "/auth?action=login"

	loginReq, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	loginReq.URL.RawQuery = loginReq.URL.Query().Encode()
	resp, err := c.client.Do(loginReq)
	if err != nil {
		c.log.Errorf("%v", err)
		return err
	}
	c.log.Debugf("Performed request %v - HTTP %d", loginReq.URL, resp.StatusCode)

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		err := &APIError{
			Code:    resp.StatusCode,
			Message: fmt.Sprintf("LOGIN - Status Code Error: %d\nRequest: %v", resp.StatusCode, loginReq),
		}
		return err
	}

	csrfToken := ""
	for _, cookie := range resp.Header["Set-Cookie"] {
		if strings.Contains(cookie, "X-CSRF-TOKEN") {
			csrfToken = strings.Split(cookie, "=")[1]
		}
	}
	req.Header.Set("X-CSRF-TOKEN", csrfToken)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cookie", strings.Join(resp.Header["Set-Cookie"], ";"))

	return nil

}
