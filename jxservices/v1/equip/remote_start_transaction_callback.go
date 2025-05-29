package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipRemoteStartTransactionCallbackRequest{}

type equipRemoteStartTransactionCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipRemoteStartTransactionCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.RemoteStartTransaction
}

func (e *equipRemoteStartTransactionCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipRemoteStartTransactionCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipRemoteStartTransactionCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipRemoteStartTransactionCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipRemoteStartTransactionCallbackResponse{}

type equipRemoteStartTransactionCallbackResponse struct {
	api.Response
}

func (resp *equipRemoteStartTransactionCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRemoteStartTransactionCallbackResponse) GetMsg() string {
	return resp.Msg
}

func RemoteStartTransactionCallbackRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.RemoteStartTransaction)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipRemoteStartTransactionCallbackResponse{})
}