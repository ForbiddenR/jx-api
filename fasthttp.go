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
	return respBody, nil
}

// not used
//func httpConnError(err error) (string, bool) {
//	errName := ""
//	know := false
//	if err == fasthttp.ErrTimeout {
//		errName = "timeout"
//		know = true
//	} else if err == fasthttp.ErrNoFreeConns {
//		errName = "conn_limit"
//		know = true
//	} else if err == fasthttp.ErrConnectionClosed {
//		errName = "conn_close"
//		know = true
//	} else {
//		errName = reflect.TypeOf(err).String()
//		if errName == "net.OpError" {
//			// Write and Read errors are not so often and in fact they just mean timeout problems
//			errName = "timeout"
//			know = true
//		}
//	}
//	return errName, know
//}

// get request
//func (a *ApiServer) sendGetRequest() {
//	req := fasthttp.AcquireRequest()
//	req.SetRequestURI("http://localhost:8080/")
//	req.Header.SetMethod(fasthttp.MethodGet)
//	resp := fasthttp.AcquireResponse()
//	err := a.client.Do(req, resp)
//	fasthttp.ReleaseRequest(req)
//	if err == nil {
//		fmt.Printf("DEBUG Response: %s\n", resp.Body())
//	} else {
//		fmt.Fprintf(os.Stderr, "Err Connection error: %v\n", err)
//	}
//	fasthttp.ReleaseResponse(resp)
//}
