package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipSyncTimeCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSyncTimeCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SyncTime
}

func (r *equipSyncTimeCallbackRequest) TraceId() string {
	return r.MsgID
}

func (equipSyncTimeCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSyncTimeCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipSyncTimeCallbackRequest {
	req := &equipSyncTimeCallbackRequest{
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

func NewEquipSyncTimeCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipSyncTimeCallbackRequest {
	req := &equipSyncTimeCallbackRequest{
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

func SyncTimeCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
