package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	"github.com/valyala/fasthttp"
)

var headerContentTypeJson = []byte("application/json")

var client *fasthttp.Client

type headers = map[string]string

func SendRequest(ctx context.Context, url string, protocol any, header map[string]string) ([]byte, error) {
	reqEntityBytes, err := json.Marshal(protocol)
	if err != nil {
		return nil, err
	}
	return sendPostRequest(ctx, url, reqEntityBytes, header)
}

func sendPostRequest(_ context.Context, url string, requestBody []byte, headers headers) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes(headerContentTypeJson)
	req.SetBodyRaw(requestBody)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

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
