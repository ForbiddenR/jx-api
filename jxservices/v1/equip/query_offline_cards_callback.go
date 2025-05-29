package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipQueryOfflineCardsCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipQueryOfflineCardsCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.QueryOfflineCards
}

func (r *equipQueryOfflineCardsCallbackRequest) TraceId() string {
	return r.MsgID
}

func (equipQueryOfflineCardsCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipQueryOfflineCardsCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipQueryOfflineCardsCallbackRequest {
	req := &equipQueryOfflineCardsCallbackRequest{
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

func NewEquipQueryOfflineCardsCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipQueryOfflineCardsCallbackRequest {
	req := &equipQueryOfflineCardsCallbackRequest{
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

func QueryOfflineCardsCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
