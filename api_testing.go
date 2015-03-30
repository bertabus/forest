package rat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// ApiTesting provides functions to call a REST api and validate its responses.
type ApiTesting struct {
	BaseUrl string
	client  *http.Client
}

// NewClient returns a new ApiTesting that can be used to send Http requests.
func NewClient(baseUrl string, httpClient *http.Client) *ApiTesting {
	return &ApiTesting{
		BaseUrl: baseUrl,
		client:  httpClient,
	}
}

// NewConfig returns a new RequestConfig.
func (a *ApiTesting) NewConfig(staticPath string) *RequestConfig {
	return NewConfig(staticPath)
}

// GET sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) GET(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("GET", a.BaseUrl+config.Uri, nil)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	if testing.Verbose() {
		print(fmt.Sprintf("\t%v %v %v\n", httpReq.Method, httpReq.URL, headersString(httpReq.Header)))
	}
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// POST sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) POST(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("POST", a.BaseUrl+config.Uri, config.BodyReader)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	if testing.Verbose() {
		print(fmt.Sprintf("\t%v %v %v\n", httpReq.Method, httpReq.URL, headersString(httpReq.Header)))
	}
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// PUT sends a Http request using a config (headers,body,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) PUT(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("PUT", a.BaseUrl+config.Uri, config.BodyReader)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	if testing.Verbose() {
		print(fmt.Sprintf("\t%v %v %v\n", httpReq.Method, httpReq.URL, headersString(httpReq.Header)))
	}
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// DELETE sends a Http request using a config (headers,...)
// The request is logged and any sending error will fail the test.
func (a *ApiTesting) DELETE(t T, config *RequestConfig) *http.Response {
	httpReq, err := http.NewRequest("DELETE", a.BaseUrl+config.Uri, nil)
	if err != nil {
		t.Fatalf("invalid Url:%s", a.BaseUrl+config.Uri)
	}
	copyHeaders(config.HeaderMap, httpReq.Header)
	if testing.Verbose() {
		print(fmt.Sprintf("\t%v %v %v\n", httpReq.Method, httpReq.URL, headersString(httpReq.Header)))
	}
	resp, err := a.client.Do(httpReq)
	CheckError(t, err)
	return resp
}

// Dump is a convenient method to log the full contents of a response
func (a ApiTesting) Dump(t T, resp *http.Response) {
	if resp == nil {
		t.Errorf("no response")
		return
	}
	if resp.ContentLength == 0 {
		t.Logf("empty response")
		return
	}
	if testing.Verbose() {
		for k, v := range resp.Header {
			print(fmt.Sprintf("\t%s : %v\n", k, strings.Join(v, ",")))
		}
	}
	if resp.Body != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		if testing.Verbose() {
			print(string(body))
		}
		resp.Body.Close()
	}
}

func headersString(h http.Header) string {
	if len(h) == 0 {
		return ""
	} else {
		return fmt.Sprintf("%v", h)
	}
}

func copyHeaders(from, to http.Header) {
	for k, list := range from {
		for _, v := range list {
			to.Set(k, v)
		}
	}
}
