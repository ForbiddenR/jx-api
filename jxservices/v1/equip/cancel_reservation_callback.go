package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipCancelReservationCallbackRequest{}

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

func NewEquipCancelReseravtionCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.CancelReservation.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipCancelReservationCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.CancelReservation.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipCancelReservationCallbackResponse{}

type equipCancelReservationCallbackResponse struct {
	api.Response
}

func (resp *equipCancelReservationCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipCancelReservationCallbackResponse) GetMsg() string {
	return resp.Msg
}

func CancelReservationCallbackRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.CancelReservation)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipCancelReservationCallbackResponse{})
}
