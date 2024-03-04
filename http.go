package http

import "github.com/valyala/fasthttp"

const (
	MethodPost = fasthttp.MethodPost
	MethodGet  = fasthttp.MethodGet
	MethodPut  = fasthttp.MethodPut
)

func Put(url string, body []byte) *Response {
	req := Request{}
	return req.Put(url, body)
}

func Get(url string) *Response {
	req := Request{}
	return req.Get(url)
}

func Post(url string, body []byte) *Response {
	req := Request{}
	return req.Post(url, body)
}

func Prepare() *Request {
	return &Request{}
}
