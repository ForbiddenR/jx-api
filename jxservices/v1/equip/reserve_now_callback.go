package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
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

func NewEquipReserveNowCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ReserveNow.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipReserveNowCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ReserveNow.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

type equipReserveNowCallbackResponse struct {
	api.Response
}

func (resp *equipReserveNowCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipReserveNowCallbackResponse) GetMsg() string {
	return resp.Msg
}

func ReserveNowCallbackRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.ReserveNow)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipReserveNowCallbackResponse{})
}
