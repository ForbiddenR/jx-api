package api

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
)

var EsamUrl, ServicesUrl string

type OcppVersion int

const (
	Ocpp16  OcppVersion = 1
	Ocpp201 OcppVersion = 2
)

func ParseOCPPVersion(version string) OcppVersion {
	switch version {
	case "ocpp1.6":
		return Ocpp16
	case "ocpp2.0.1":
		return Ocpp201
	default:
		return Ocpp16
	}
}

const (
	Services = "services"
	Esam     = "esam"
)

const msgID = "msgId"

const (
	Perms = "Perms"
	TraceId = "TraceId"
)

var (
	ErrBodyIsNil = errors.New("body is nil")
	// service-caused error
	ErrServicesException = errors.New("services exception")
	// 404 not found error
	ErrNotFound = errors.New("not found")
)

type Response struct {
	Status    int    `json:"status"`
	Rows      int    `json:"rows"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

func (r *Response) GetStatus() int {
	return r.Status
}

func (r *Response) GetMsg() string {
	return r.Msg
}

func UnmarshalAndVerify(payload []byte, req any, validate *validator.Validate) error {
	if err := json.Unmarshal(payload, req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}

func Decode(payload []byte, req any, validate *validator.Validate) (string, error) {
	msgId := gjson.GetBytes(payload, msgID).String()
	if err := json.Unmarshal(payload, req); err != nil {
		return msgId, err
	}
	if err := validate.Struct(req); err != nil {
		return msgId, err
	}
	return msgId, nil
}
