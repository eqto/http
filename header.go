package http

import (
	"github.com/valyala/fasthttp"
)

type Header struct {
	header *fasthttp.ResponseHeader
}

func (r *Header) Get(key string) string {
	byteVal := r.header.Peek(key)
	if (byteVal != nil) {
		return string(byteVal)
	}
	return ``
}