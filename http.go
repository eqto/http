package http

import (
	"io"

	"github.com/valyala/fasthttp"
)

const (
	MethodPost = fasthttp.MethodPost
	MethodGet  = fasthttp.MethodGet
	MethodPut  = fasthttp.MethodPut
)

func Put(url string, body io.Reader) (*Response, error) {
	req := Request{}
	return req.Put(url, body)
}

func Get(url string) (*Response, error) {
	req := Request{}
	return req.Get(url)
}

func Post(url, contentType string, body io.Reader) (*Response, error) {
	req := Request{method: MethodPost}
	req.SetHeader(`Content-Type`, contentType)
	return req.Post(url, body)
}

func NewRequest(method, url string, body io.Reader) *Request {
	return &Request{method: method, url: url, bodyReader: body}
}
