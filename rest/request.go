package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/ForbiddenR/jxapi/apierrors"
	"github.com/valyala/fasthttp"
)

var headerContentTypeJson = []byte("application/json")

type Request struct {
	c *RESTClient

	req  *fasthttp.Request
	path string

	// output
	err error
}

func NewRequest(c *RESTClient) *Request {
	r := &Request{
		c:   c,
		req: fasthttp.AcquireRequest(),
	}
	return r
}

func (r *Request) Verb(verb string) *Request {
	r.req.Header.SetMethod(verb)
	return r
}

func (r *Request) RequestURI(uri string) *Request {
	if r.err != nil {
		return r
	}
	r.path = uri
	return r
}

func (r *Request) SetHeader(key, value string) *Request {
	r.req.Header.Set(key, value)
	return r
}

func (r *Request) SetParam() *Request {
	r.req.Header.SetContentTypeBytes(headerContentTypeJson)
	r.req.Header.DisableNormalizing()
	return r
}

func (r *Request) Body(body any) *Request {
	if r.err != nil {
		return r
	}
	raw, err := json.Marshal(body)
	if err != nil {
		r.err = err
		return r
	}
	r.req.SetBodyRaw(raw)
	return r
}

// request connects to the server and invokes the provided function when a server response is
// received.
func (r *Request) request(_ context.Context, fn func(*fasthttp.Request, *fasthttp.Response)) error {
	finalURL := &url.URL{}
	if r.c.base != nil {
		*finalURL = *r.c.base
	}
	finalURL.Path = r.path
	r.req.SetRequestURI(finalURL.String())
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(r.req)
	}()
	err := r.c.Client.DoTimeout(r.req, resp, 3*time.Second)
	if err != nil {
		return apierrors.GetFailedRequestDoTimeoutError(err)
	}
	fn(r.req, resp)
	return nil
}

func (r *Request) Do(ctx context.Context) Result {
	var result Result
	err := r.request(ctx, func(req *fasthttp.Request, resp *fasthttp.Response) {
		result = r.transformResponse(resp, req)
	})
	if err != nil {
		return Result{err: err}
	}
	return result
}

func (r *Request) DoRaw(ctx context.Context) ([]byte, error) {
	var result Result
	err := r.request(ctx, func(req *fasthttp.Request, resp *fasthttp.Response) {
		result = r.transformResponse(resp, req)
	})
	if err != nil {
		return nil, err
	}
	return result.body, result.err
}

// transformResponse converts an API response into a structured API object.
func (r *Request) transformResponse(resp *fasthttp.Response, _ *fasthttp.Request) Result {
	var body []byte
	respBody := resp.Body()
	if len(respBody) == 0 {
		return Result{err: ErrBodyIsNil}
	}
	body = make([]byte, len(respBody))
	copy(body, respBody)
	if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
		if statusCode == fasthttp.StatusNotFound {
			return Result{err: ErrNotFound}
		}
		fmt.Println("error code is", statusCode)
		return Result{err: ErrServicesException}
	}
	return Result{body: body}
}

// Error returns any error encountered construcing the reuest, if any.
func (r *Request) Error() error {
	return r.err
}

type Result struct {
	body []byte
	err  error
}

func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

func (r Result) Into(obj any) error {
	if r.err != nil {
		return r.err
	}
	return json.Unmarshal(r.body, obj)
}

func (r Result) Error() error {
	return r.err
}
