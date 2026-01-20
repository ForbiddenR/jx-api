package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ForbiddenR/jxapi/apierrors"
	"github.com/valyala/fasthttp"
)

var headerContentTypeJson = []byte("application/json")

var client *fasthttp.Client

type HttpOption func(*fasthttp.Request)

func WithHeader[V ~string](key, value V) HttpOption {
	return func(r *fasthttp.Request) {
		r.Header.Set(string(key), string(value))
	}
}

func WithHeaders[V ~map[S]S, S ~string](m V) HttpOption {
	return func(r *fasthttp.Request) {
		for k, v := range m {
			r.Header.Set(string(k), string(v))
		}
	}
}

func SendRequest[T any](ctx context.Context, url string, protocol T, opts ...HttpOption) ([]byte, error) {
	reqEntityBytes, err := json.Marshal(protocol)
	if err != nil {
		return nil, err
	}
	return sendPostRequest(ctx, url, reqEntityBytes, opts...)
}

func sendPostRequest(_ context.Context, url string, requestBody []byte, opts ...HttpOption) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	for _, opt := range opts {
		opt(req)
	}
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes(headerContentTypeJson)
	req.Header.DisableNormalizing()
	req.SetBodyRaw(requestBody)
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	err := client.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		return nil, apierrors.GetFailedRequestDoTimeoutError(err)
	}
	if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
		if statusCode == fasthttp.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, ErrServicesException
	}
	respBody := resp.Body()
	if len(respBody) == 0 {
		return nil, ErrBodyIsNil
	}
	return append([]byte{}, respBody...), nil
}
