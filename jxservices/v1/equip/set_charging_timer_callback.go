package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipSetChargingTimerCallbackRequest{}

type equipSetChargingTimerCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSetChargingTimerCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SetChargingTimer
}

func (e *equipSetChargingTimerCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipSetChargingTimerCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSetChargingTimerCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
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

func NewEquipSetChargingTimerCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
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

// var _ services.Response = &equipSetChargingTimerCallbackResponse{}

// type equipSetChargingTimerCallbackResponse struct {
// 	api.Response
// }

// func (resp *equipSetChargingTimerCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipSetChargingTimerCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

func SetChargingTimerCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
