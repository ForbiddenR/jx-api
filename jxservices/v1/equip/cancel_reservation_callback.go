package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipCancelReservationCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipCancelReservationCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.CancelReservation
}

func (r *equipCancelReservationCallbackRequest) TraceId() string {
	return r.MsgID
}

func (equipCancelReservationCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipCancelReseravtionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
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

func NewEquipCancelReservationCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
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

// var _ services.Response = &equipCancelReservationCallbackResponse{}

// type equipCancelReservationCallbackResponse struct {
// 	api.Response
// }

// func (resp *equipCancelReservationCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipCancelReservationCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

func CancelReservationCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
