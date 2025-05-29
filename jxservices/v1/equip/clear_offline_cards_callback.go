package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipClearOfflineCardsCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipClearOfflineCardsCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.ClearOfflineCards
}

func (r *equipClearOfflineCardsCallbackRequest) TraceId() string {
	return r.MsgID
}

func (r *equipClearOfflineCardsCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipClearOfflineCardsCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipClearOfflineCardsCallbackRequest {
	req := &equipClearOfflineCardsCallbackRequest{
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

func NewEquipClearOfflineCardsCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipClearOfflineCardsCallbackRequest {
	req := &equipClearOfflineCardsCallbackRequest{
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

func ClearOfflineCardsCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
