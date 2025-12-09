package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipLoginRequest{}

type equipLoginRequest struct {
	services.Base
	Data *equipLoginRequestDetail `json:"data"`
}

type equipLoginRequestDetail struct {
	RemoteAddress      *string `json:"remoteAddress"`
	ModelCode          string  `json:"modelCode"`
	ManufacturerCode   string  `json:"manufacturerCode"`
	FirmwareVersion    *string `json:"firmwareVersion"`
	Iccid              *string `json:"iccid"`
	Imsi               *string `json:"imsi"`
	ReconnectingReason *string `json:"reason"`
}

func (equipLoginRequest) GetName() services.Request2ServicesNameType {
	return services.Login
}

func (e *equipLoginRequest) TraceId() string {
	return e.MsgID
}

func (equipLoginRequest) IsCallback() bool {
	return false
}

type LoginRequestConfig struct {
	ModelCode        string
	ManufacturerCode string
}

func NewLogin(base services.Base, config *LoginRequestConfig) *equipLoginRequest {
	req := &equipLoginRequest{
		Base: base,
		Data: &equipLoginRequestDetail{
			ModelCode:        config.ModelCode,
			ManufacturerCode: config.ManufacturerCode,
		},
	}
	return req
}

func NewEquipLoginRequest(sn, id, pod, msgId string, p *services.Protocol, modelCode, manufacturerCode string) *equipLoginRequest {
	request := &equipLoginRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
	request.Data = &equipLoginRequestDetail{
		ModelCode:        modelCode,
		ManufacturerCode: manufacturerCode,
	}
	return request
}

var _ services.Response = &equipLoginResponse{}

type equipLoginResponse struct {
	api.Response
}

func (resp *equipLoginResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipLoginResponse) GetMsg() string {
	return resp.Msg
}

func LoginRequest(ctx context.Context, req services.Request) error {
	// header := services.GetSimpleHeaderValue(services.Login)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipLoginResponse{})
}