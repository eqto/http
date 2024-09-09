package http

import (
	"github.com/valyala/fasthttp"
)

type Header struct {
	header *fasthttp.ResponseHeader
}

func (r *Header) Get(key string) string {
	byteVal := r.header.Peek(key)
	if byteVal != nil {
		return string(byteVal)
	}
	return ``
}

func (r *Header) Set(key, value string) {
	if r.header == nil {
		r.header = &fasthttp.ResponseHeader{}
	}
	r.header.Set(key, value)
}

func (h *Header) Cookie(key string) *Cookie {
	cookie := new(Cookie)
	byteCookie := h.header.PeekCookie(key)
	if cookie.parse(key, byteCookie) {
		return cookie
	}
	return nil
}
