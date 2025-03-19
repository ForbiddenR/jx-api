package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipRemoteStopTransactionCallbackRequest{}

type equipRemoteStopTransactionCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (s *equipRemoteStopTransactionCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.RemoteStopTransaction
}

func (s *equipRemoteStopTransactionCallbackRequest) TraceId() string {
	return s.MsgID
}

func (s *equipRemoteStopTransactionCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipRemoteStopTransactionCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipRemoteStopTransactionCallbackRequest {
	req := &equipRemoteStopTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.RemoteStopTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipRemoteStopTransactionCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteStopTransactionCallbackRequest {
	req := &equipRemoteStopTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.RemoteStopTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipRemoteStopTransactionCallbackResponse{}

type equipRemoteStopTransactionCallbackResponse struct {
	api.Response
}

func (resp *equipRemoteStopTransactionCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRemoteStopTransactionCallbackResponse) GetMsg() string {
	return resp.Msg
}

func RemoteStopTransactionRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.RemoteStopTransaction)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipRemoteStopTransactionCallbackResponse{})
}
