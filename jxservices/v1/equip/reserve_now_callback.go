package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipReserveNowCallbackRequest{}

type equipReserveNowCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipReserveNowCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.ReserveNow
}

func (e *equipReserveNowCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipReserveNowCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipReserveNowCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
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

func NewEquipReserveNowCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
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

// type equipReserveNowCallbackResponse struct {
// 	api.Response
// }

// func (resp *equipReserveNowCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipReserveNowCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

func ReserveNowCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
