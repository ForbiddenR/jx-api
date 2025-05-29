package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipRemoteUpdateAccountBalanceCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipRemoteUpdateAccountBalanceCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.RemoteUpdateAccountBalance
}

func (e *equipRemoteUpdateAccountBalanceCallbackRequest) TraceId() string {
	return e.MsgID
}

func (e *equipRemoteUpdateAccountBalanceCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipRemoteUpdateAccountBalanceCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipRemoteUpdateAccountBalanceCallbackRequest {
	return &equipRemoteUpdateAccountBalanceCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipRemoteUpdateAccountBalanceCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteUpdateAccountBalanceCallbackRequest {
	return &equipRemoteUpdateAccountBalanceCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
}

func RemoteUpdateAccountBalanceCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
