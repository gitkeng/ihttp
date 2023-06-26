package ihttp

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/gitkeng/ihttp/util/fileutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/go-resty/resty/v2"
	"os"
	"time"
)

// IRequester is interface to connect to HTTP endpoint
type IRequester interface {
	Get(path string, params map[string]string) (string, error)
	Post(path string, params map[string]string) (string, error)
	PostJSON(path string, body interface{}) (string, error)
	Put(path string, params map[string]string) (string, error)
	PutJSON(path string, body interface{}) (string, error)
	Delete(path string, params map[string]string) (string, error)
	GetClient() *resty.Client
	Request() *resty.Request
}

// Requester implement IRequester
type Requester struct {
	baseURL string
	ms      *Microservice
	client  *resty.Client
}

// NewRequester return new Requester
func NewRequester(
	ms *Microservice,
	baseURL string,
	timeout time.Duration,
	certFiles ...string) (*Requester, error) {

	if stringutil.IsEmptyString(baseURL) {
		return nil, errors.New("baseURL is required")
	}

	client := resty.New()

	if len(certFiles) > 0 {
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		for _, certFile := range certFiles {
			if found, _ := fileutil.IsFileExist(certFile); !found {
				return nil, fmt.Errorf("certFile %s not found", certFile)
			}

			certRaw, err := os.ReadFile(certFile)
			if err != nil {
				return nil, err
			}

			if ok := rootCAs.AppendCertsFromPEM(certRaw); !ok {
				return nil, errors.New("failed to append cert")
			}
		}

		client.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            rootCAs,
		})
	}

	if timeout > 0 {
		client.SetTimeout(timeout)
	}

	return &Requester{
		baseURL: baseURL,
		ms:      ms,
		client:  client,
	}, nil
}

// Get request using HTTP GET
func (rqt *Requester) Get(path string, params map[string]string) (string, error) {

	url := fmt.Sprint(rqt.baseURL, path)

	resp, err := rqt.client.SetQueryParams(params).R().Get(url)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}

// Delete request using HTTP DELETE
func (rqt *Requester) Delete(path string, params map[string]string) (string, error) {

	url := fmt.Sprint(rqt.baseURL, path)

	resp, err := rqt.client.SetQueryParams(params).R().Delete(url)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil

}

// Post request using HTTP POST
func (rqt *Requester) Post(path string, params map[string]string) (string, error) {

	url := fmt.Sprint(rqt.baseURL, path)

	resp, err := rqt.client.SetQueryParams(params).R().Post(url)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}

// PostJSON request using HTTP POST with JSON body
func (rqt *Requester) PostJSON(path string, jsonBody any) (string, error) {

	url := fmt.Sprint(rqt.baseURL, path)

	resp, err := rqt.client.R().SetBody(jsonBody).Post(url)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil

}

// Put request using HTTP PUT
func (rqt *Requester) Put(path string, params map[string]string) (string, error) {

	url := fmt.Sprint(rqt.baseURL, path)

	resp, err := rqt.client.SetQueryParams(params).R().Put(url)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}

// PutJSON request using HTTP PUT with JSON body
func (rqt *Requester) PutJSON(path string, jsonBody interface{}) (string, error) {

	url := fmt.Sprint(rqt.baseURL, path)

	resp, err := rqt.client.R().SetBody(jsonBody).Put(url)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}

// GetRequest return resty.Request
func (rqt *Requester) Request() *resty.Request {
	return rqt.client.R()
}

func (rqt *Requester) GetClient() *resty.Client {
	return rqt.client
}
