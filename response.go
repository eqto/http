package http

import "github.com/valyala/fasthttp"

type Response struct {
	resp *fasthttp.Response
	err  error
	body []byte
}

func (r *Response) IsError() bool {
	return r.err != nil
}

func (r *Response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.resp.StatusCode()
}

func (r *Response) Body() []byte {
	if r.resp == nil {
		return []byte{}
	}
	if r.body == nil {
		r.body = r.resp.Body()
	}
	return r.body
}

func (r *Response) Error() error {
	return r.err
}
