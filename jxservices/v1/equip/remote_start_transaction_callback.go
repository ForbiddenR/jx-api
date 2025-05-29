package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

// var _ services.Request = &equipRemoteStartTransactionCallbackRequest{}

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

func NewEquipRemoteStartTransactionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
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

func NewEquipRemoteStartTransactionCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
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

// var _ services.Response = &equipRemoteStartTransactionCallbackResponse{}

// type equipRemoteStartTransactionCallbackResponse struct {
// 	api.Response
// }

// func (resp *equipRemoteStartTransactionCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipRemoteStartTransactionCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

func RemoteStartTransactionCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
