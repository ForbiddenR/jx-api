package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

const (
	SetVariablesAccept         = 0
	SetVariablesRejected       = 1
	SetVariablesRebootRequired = 2
	SetVariablesNotSupported   = 3
)

// func OCPP16SetVariablesStatus(status protocol.ChangeConfigurationResponseJsonStatus) int {
// 	switch status {
// 	case protocol.ChangeConfigurationResponseJsonStatusAccepted:
// 		return SetVariablesAccept
// 	case protocol.ChangeConfigurationResponseJsonStatusRejected:
// 		return SetVariablesRejected
// 	case protocol.ChangeConfigurationResponseJsonStatusRebootRequired:
// 		return SetVariablesRebootRequired
// 	default:
// 		return SetVariablesNotSupported
// 	}
// }

var _ services.Request = &equipSetVariablesCallbackRequest{}

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

func NewEquipSetVariablesCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ChangeConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
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
			Category:    services.ChangeConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSetVariablesCallbackResponse{}

type equipSetVariablesCallbackResponse struct {
	api.Response
}

func (resp *equipSetVariablesCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetVariablesCallbackResponse) GetMsg() string {
	return resp.Msg
}

func SetVariablesRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.ChangeConfiguration)
	url := services.GetCallbackURL(req)
	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetVariablesCallbackResponse{})
}
