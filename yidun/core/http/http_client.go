package http

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

type Request interface {
	Method() string
	Url() string
	Headers() map[string]string
	Body() []byte
}

type Response interface {
	StatusCode() int
	Headers() http.Header
	Body() []byte
	SetStatusCode(statusCode int)
	SetHeaders(headers http.Header)
	SetBody(body []byte)
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &HttpClient{
		client: client,
	}
}

func (c *HttpClient) Do(request Request, response Response) error {
	req, err := http.NewRequest(request.Method(), request.Url(), bytes.NewReader(request.Body()))
	if err != nil {
		return err
	}

	headers := http.Header{}
	// 遍历 map，将键值对添加到 http.Header 对象中
	for k, v := range request.Headers() {
		headers.Add(k, v)
	}
	req.Header = headers

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response.SetStatusCode(resp.StatusCode)
	response.SetHeaders(resp.Header)
	response.SetBody(responseBody)

	return nil
}

type HttpRequest struct {
	MethodValue  string
	UrlValue     string
	HeadersValue map[string]string
	BodyValue    []byte
}

func (r *HttpRequest) Method() string {
	return r.MethodValue
}

func (r *HttpRequest) Url() string {
	return r.UrlValue
}

func (r *HttpRequest) Headers() map[string]string {
	return r.HeadersValue
}

func (r *HttpRequest) Body() []byte {
	if r.BodyValue == nil {
		return nil
	}
	return r.BodyValue
}

type CommonResponse struct {
	StatusCodeValue int
	HeadersValue    http.Header
	BodyValue       []byte
}

func (r *CommonResponse) StatusCode() int {
	return r.StatusCodeValue
}

func (r *CommonResponse) Headers() http.Header {
	return r.HeadersValue
}

func (r *CommonResponse) Body() []byte {
	return r.BodyValue
}

func (r *CommonResponse) SetStatusCode(statusCode int) {
	r.StatusCodeValue = statusCode
}

func (r *CommonResponse) SetHeaders(headers http.Header) {
	r.HeadersValue = headers
}

func (r *CommonResponse) SetBody(body []byte) {
	r.BodyValue = body
}

type CustomClient struct {
	client    *http.Client
	userAgent string
}

func (c *CustomClient) Do(request Request, response Response) error {
	req, err := http.NewRequest(request.Method(), request.Url(), bytes.NewBuffer(request.Body()))
	if err != nil {
		return err
	}

	// 遍历 map，将键值对添加到 http.Header 对象中
	for k, v := range request.Headers() {
		req.Header.Set(k, v)
	}
	// 为请求添加 User-Agent 头
	req.Header.Set("User-Agent", c.userAgent)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response.SetStatusCode(resp.StatusCode)
	response.SetHeaders(resp.Header)
	response.SetBody(responseBody)

	return nil
}

func NewCustomClient(client *http.Client) *CustomClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &CustomClient{
		client: client,
	}
}

type RequestOptions struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    interface{}
}

func (c *CustomClient) Request(options *RequestOptions, response Response) error {
	bodyBytes, err := json.Marshal(options.Body)
	if err != nil {
		return err
	}

	request := &HttpRequest{
		MethodValue:  options.Method,
		UrlValue:     options.Url,
		HeadersValue: options.Headers,
		BodyValue:    bodyBytes,
	}

	return c.Do(request, response)
}

type FormOptions struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    map[string][]string
}

func (c *CustomClient) Form(options *FormOptions, response Response) error {

	if options.Body == nil {
		return nil
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, values := range options.Body {
		for _, value := range values {
			_ = writer.WriteField(key, value)
		}
	}
	_ = writer.Close()

	request := &HttpRequest{
		MethodValue:  options.Method,
		UrlValue:     options.Url,
		HeadersValue: options.Headers,
		BodyValue:    body.Bytes(),
	}
	return c.Do(request, response)
}
