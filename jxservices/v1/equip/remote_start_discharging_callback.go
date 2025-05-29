package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipRequestStartDischargingRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipRequestStartDischargingRequest) GetName() services.Request2ServicesNameType {
	return services.RequestStartDischargingTransaction
}

func (e *equipRequestStartDischargingRequest) TraceId() string {
	return e.MsgID
}

func (equipRequestStartDischargingRequest) IsCallback() bool {
	return true
}

func NewEquipRequestStartDischargingRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipRequestStartDischargingRequest {
	return &equipRequestStartDischargingRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipRequestStartDischargingRequestError(sn string, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipRequestStartDischargingRequest {
	req := &equipRequestStartDischargingRequest{
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

func RequestStartDischargingCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
