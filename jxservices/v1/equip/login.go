package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipLoginRequest{}

type equipLoginRequest struct {
	services.Base
	Data *equipLoginRequestDetail `json:"data"`
}

type equipLoginRequestDetail struct {
	RemoteAddress      *string `json:"remoteAddress"`
	ModelCode          *string `json:"modelCode,omitempty"`
	ManufacturerCode   *string `json:"manufacturerCode,omitempty"`
	FirmwareVersion    *string `json:"firmwareVersion"`
	Iccid              *string `json:"iccid,omitempty"`
	Imsi               *string `json:"imsi,omitempty"`
	ReconnectingReason *string `json:"reason,omitempty"`
	ChargerType        *int16  `json:"chargerType"`
	NetworkLink        *int16  `json:"networkLink"`
	Carrier            *int16  `json:"carrier"`
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

// type LoginRequestConfig struct {
// 	ModelCode        string
// 	ManufacturerCode string
// }

func NewLogin(base services.Base) *equipLoginRequest {
	req := &equipLoginRequest{
		Base: base,
		Data: &equipLoginRequestDetail{},
	}
	return req
}

func NewEquipLoginRequest(sn, pod, msgID string, p *services.Protocol) *equipLoginRequest {
	request := &equipLoginRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipLoginRequestDetail{},
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
	return services.Transport(ctx, req)
}
