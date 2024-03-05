package http

import (
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type Request struct {
	method string
	url    string
	bodyReader io.Reader
	header map[string]string
	query  map[string]any
}

func (r *Request) request(timeout time.Duration) (*Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	urlStr := r.url
	if u, e := url.Parse(r.url); e == nil {
		q := u.Query()
		for key, val := range r.query {
			switch val := val.(type) {
			case string:
				q.Set(key, val)
			case int:
				q.Set(key, fmt.Sprintf(`%d`, val))
			}
		}
		u.RawQuery = q.Encode()
		urlStr = u.String()
	}
	req.SetRequestURI(urlStr)

	req.Header.SetMethod(r.method)
	if r.header != nil {
		for key, val := range r.header {
			req.Header.Set(key, val)
		}
	}

	if r.bodyReader != nil {
		body, e := io.ReadAll(r.bodyReader)
		if  e != nil {
			return nil, e
		}
		req.SetBody(body)
	}
	if timeout == 0 {
		timeout = 60 * time.Second
	}
	if ctype, ok := r.header[`Content-Type`]; ok {
		req.Header.SetContentType(ctype)
	} else {
		req.Header.SetContentType(`application/json`)
	}

	httpResp := Response{
		resp: &fasthttp.Response{},
	}

	httpResp.err = fasthttp.DoTimeout(req, resp, timeout)
	if resp != nil {
		resp.CopyTo(httpResp.resp)
	}
	return &httpResp, nil
}

func (r *Request) SetHeader(key, val string) {
	if r.header == nil {
		r.header = make(map[string]string)
	}
	r.header[key] = val
}

func (r *Request) SetQuery(key, val string) {
	r.setQuery(key, val)
}
func (r *Request) SetQueryInt(key string, val int) {
	r.setQuery(key, val)
}

func (r *Request) setQuery(key string, val any) {
	if r.query == nil {
		r.query = make(map[string]any)
	}
	r.query[key] = val
}

func (r *Request) Put(url string, body io.Reader) (*Response, error) {
	r.method = MethodPut
	r.url = url
	r.bodyReader = body
	return r.request(60 * time.Second)
}

func (r *Request) Get(url string) (*Response, error) {
	r.method = MethodGet
	r.url = url
	return r.request(60 * time.Second)
}
func (r *Request) Post(url string, body io.Reader) (*Response, error) {
	r.method = MethodPost
	r.url = url
	r.bodyReader = body
	return r.request(60 * time.Second)
}
