package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipResetOfflineCardsCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipResetOfflineCardsCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.ResetOfflineCards
}

func (equipResetOfflineCardsCallbackRequest) IsCallback() bool {
	return true
}

func (r *equipResetOfflineCardsCallbackRequest) TraceId() string {
	return r.MsgID
}

func NewEquipResetOfflineCardsCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipResetOfflineCardsCallbackRequest {
	return &equipResetOfflineCardsCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipResetOfflineCardsCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipResetOfflineCardsCallbackRequest {
	return &equipResetOfflineCardsCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
}

func ResetOfflineCardsCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
