package appdrest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

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
	Analytics           *AnalyticsService
	Application         *ApplicationService
	BusinessTransaction *BusinessTransactionService
	Dashboard           *DashboardService
	MetricData          *MetricDataService
	Snapshot            *SnapshotService
	Tier                *TierService
	Node                *NodeService
	TimeRange           *TimeRangeService
	Configuration       *Configuration
}

type service struct {
	client *Client
}

var log = logging.MustGetLogger("appdrest")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{level} %{shortfunc} %{id:03x}%{color:reset} %{message}`,
)

// NewClient Returns a Client, this is needed for any communication
func NewClient(protocol string, controllerHost string, port int, username string, password string, account string) (*Client, error) {

	// TODO: Let the consumer define the http.Client
	timeout := time.Duration(60 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Timeout: timeout,
		Transport: tr,
	}
	baseURL, err := url.Parse(fmt.Sprintf("%s://%s:%d/", protocol, controllerHost, port))
	if err != nil {
		return nil, err
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
	backend1Leveled.SetLevel(logging.WARNING, "")

	logging.SetBackend(backend1Leveled)

	// TODO: Let the consumer define the logger
	c.log = log
	c.common.client = c

	c.Account = (*AccountService)(&c.common)
	c.Analytics = (*AnalyticsService)(&c.common)
	c.Application = (*ApplicationService)(&c.common)
	c.BusinessTransaction = (*BusinessTransactionService)(&c.common)
	c.MetricData = (*MetricDataService)(&c.common)
	c.Snapshot = (*SnapshotService)(&c.common)
	c.Tier = (*TierService)(&c.common)
	c.Dashboard = (*DashboardService)(&c.common)
	c.Node = (*NodeService)(&c.common)
	c.TimeRange = (*TimeRangeService)(&c.common)
	c.Configuration = (*Configuration)(&c.common)

	c.log.Debug("Created client successfully")
	return c, nil
}

// Rest makes a call using the standard Rest API
func (c *Client) Rest(method string, url string, model interface{}, body interface{}) error {

	req, err := c.newRequest(method, url, body)
	if err != nil {
		return err
	}

	if model == nil {
		err = c.do(req, nil, false)
	} else {
		err = c.do(req, &model, false)
	}

	if err != nil {
		return err
	}
	return nil
}

// RestInternal makes a call using the internal API that requires authorization
func (c *Client) RestInternal(method string, url string, model interface{}, body interface{}) error {

	req, err := c.newRequest(method, url, body)
	if err != nil {
		return err
	}
	err = c.do(req, &model, true)
	if err != nil {
		return err
	}
	return nil
}

// newRequest performs a request.
// The baseURL on the client will be concatenated with the url argument
func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
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
func (c *Client) do(req *http.Request, v interface{}, authorization bool) error {

	req.URL.RawQuery = req.URL.Query().Encode()

	// If we are here, this is an internal call that needs extra authorization
	if authorization == true {
		if len(req.Header["X-CSRF-TOKEN"]) == 0 {
			c.log.Debugf("RESTUI, logging in...")
			err := c.login(req)
			if err != nil {
				return err
			}

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
			Message: fmt.Sprintf("Status Code Error: %d\n", resp.StatusCode),
		}
		c.log.Errorf("Error Request: %v Response: %v\n", req, resp)
		return err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return err

		}
	}
	return nil

}

// DoRawRequest makes an HTTP request and returns the response
func (c *Client) DoRawRequest(method string, url string, body interface{}) ([]byte, error) {

	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		err := &APIError{
			Code:    resp.StatusCode,
			Message: fmt.Sprintf("Error - Status Code: %d\n", resp.StatusCode),
		}
		c.log.Errorf("Error %v, Request: %v, Response: %v", err, req, resp)
		return nil, err
	}

	responseString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseString, nil
}

func (c *Client) login(req *http.Request) error {

	url := "/auth?action=login"

	loginReq, err := c.newRequest("GET", url, nil)
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
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "X-CSRF-TOKEN" {
			csrfToken = cookie.Value
		}
	}
	req.Header.Set("X-CSRF-TOKEN", csrfToken)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cookie", strings.Join(resp.Header["Set-Cookie"], ";"))

	return nil

}
