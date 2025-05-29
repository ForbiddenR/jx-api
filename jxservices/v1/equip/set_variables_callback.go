package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

const (
	SetVariablesAccept         = 0
	SetVariablesRejected       = 1
	SetVariablesRebootRequired = 2
	SetVariablesNotSupported   = 3
)

type equipSetVariablesCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSetVariablesCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.ChangeConfiguration
}

func (e *equipSetVariablesCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipSetVariablesCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSetVariablesCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetVariablesRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

// var _ services.Response = &equipSetVariablesCallbackResponse{}

// type equipSetVariablesCallbackResponse struct {
// 	api.Response
// }

// func (resp *equipSetVariablesCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipSetVariablesCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

func SetVariablesRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
