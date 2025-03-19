package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipCancelIntellectChargeRequest{}

type equipCancelIntellectChargeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipCancelIntellectChargeRequest) GetName() services.Request2ServicesNameType {
	return services.CancelIntellectCharge
}

func (e *equipCancelIntellectChargeRequest) TraceId() string {
	return e.MsgID
}

func (equipCancelIntellectChargeRequest) IsCallback() bool {
	return true
}

func NewEquipCancelIntellectChargeCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipCancelIntellectChargeRequest {
	req := &equipCancelIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.CancelIntellectCharge.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipCancelIntellectChargeCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipCancelIntellectChargeRequest {
	req := &equipCancelIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.CancelIntellectCharge.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipCancelIntellectChargeResponse{}

type equipCancelIntellectChargeResponse struct {
	api.Response
}

func (resp *equipCancelIntellectChargeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipCancelIntellectChargeResponse) GetMsg() string {
	return resp.Msg
}

func CancelIntellectChargeRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.CancelIntellectCharge)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipCancelIntellectChargeResponse{})
}
