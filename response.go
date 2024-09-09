package http

import (
	"bytes"
	"errors"
	"io"

	"github.com/valyala/fasthttp"
)

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

func (r *Response) Body() io.ReadCloser {
	buff := bytes.NewBuffer(r.body)
	return &bodyReader{
		data: buff,
	}
}

func (r *Response) Error() error {
	return r.err
}

func (r *Response) Header() *Header {
	header := new(Header)
	if r.resp != nil {
		header.header = &r.resp.Header
	}
	return header
}

type bodyReader struct {
	data *bytes.Buffer
}

func (b *bodyReader) ReadAll() ([]byte, error) {
	return io.ReadAll(b.data)
}
func (b *bodyReader) Read(p []byte) (n int, err error) {
	if b.data == nil {
		return 0, errors.New(`empty body`)
	}
	return b.data.Read(p)
}

func (b *bodyReader) Close() error {
	return nil
}
