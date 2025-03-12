package equip

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/ForbiddenR/jxapi/jxservices"
	"github.com/ForbiddenR/jxapi/rest"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestCallStatusNotificationRequest(t *testing.T) {
	client := &fasthttp.Client{
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
		MaxIdleConnDuration: 10 * time.Second,
		MaxConnsPerHost:     5000,
	}
	req := NewEquipCallStatusNotificationCallbackRequest("", "", "", "", jxservices.IEC002(), 1)
	service, err := url.Parse("http://127.0.0.1:12000/")
	assert.Nil(t, err)
	serviceClient, err := rest.NewRestClient(service, rest.ClientContentConfig{}, client)
	assert.Nil(t, err)
	result := serviceClient.Post().
		RequestURI("test").
		Body(req).
		SetHeader(map[string]string{"test": "test"}).
		Do(context.Background())
	assert.Nil(t, result.Error())

}
