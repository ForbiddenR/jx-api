package jxservices

import "github.com/ForbiddenR/jxapi/apierrors"

// Request interface needs to be implemented by all api.
type Request interface {
	GetName() Request2ServicesNameType
	TraceId() string
	IsCallback() bool
}

// Response will be implemented by all response struct.
type Response interface {
	GetStatus() int
	GetMsg() string
}

type CallbackRequest interface {
	Request
	SetCallback(cb CB)
}

type ReusedConfig struct {
	Sn       string
	Id       string
	Protocol *Protocol
	Pod      string
	MsgId    string
}

type Option func(option CallbackRequest) error

func WithStatus(status int) Option {
	return func(option CallbackRequest) error {
		option.SetCallback(NewCB(status))
		return nil
	}
}

func WithError(err *apierrors.CallbackError) Option {
	return func(option CallbackRequest) error {
		option.SetCallback(NewCBError(err))
		return nil
	}
}
