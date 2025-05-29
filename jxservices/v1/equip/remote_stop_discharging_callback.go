package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipRequestStopDischargingRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipRequestStopDischargingRequest) GetName() services.Request2ServicesNameType {
	return services.RequestStartDischargingTransaction
}

func (e *equipRequestStopDischargingRequest) TraceId() string {
	return e.MsgID
}

func (equipRequestStopDischargingRequest) IsCallback() bool {
	return true
}

func NewEquipRequestStopDischargingRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipRequestStopDischargingRequest {
	return &equipRequestStopDischargingRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipRequestStopDischargingRequestError(sn string, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipRequestStopDischargingRequest {
	req := &equipRequestStopDischargingRequest{
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

func RequestStopDischargingCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
