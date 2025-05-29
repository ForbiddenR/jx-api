package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipRemoteStartParallelTransactionCallbackRquest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipRemoteStartParallelTransactionCallbackRquest) GetName() services.Request2ServicesNameType {
	return services.RemoteStartParallelTransaction
}

func (r *equipRemoteStartParallelTransactionCallbackRquest) TraceId() string {
	return r.MsgID
}

func (equipRemoteStartParallelTransactionCallbackRquest) IsCallback() bool {
	return true
}

func NewEquipRemoteStartParallelTransactionCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipRemoteStartParallelTransactionCallbackRquest {
	req := &equipRemoteStartParallelTransactionCallbackRquest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipRemoteStartParallelTransactionCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteStartParallelTransactionCallbackRquest {
	req := &equipRemoteStartParallelTransactionCallbackRquest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

func RemoteStartParallelTransactionCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
