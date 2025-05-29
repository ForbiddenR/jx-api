package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipBMSTerminateRequest struct {
	services.Base
}

func (r *equipBMSTerminateRequest) GetName() services.Request2ServicesNameType {
	return services.BmsTerminate
}

func (r *equipBMSTerminateRequest) TraceId() string {
	return r.MsgID
}

func (r *equipBMSTerminateRequest) IsCallback() bool {
	return false
}

func NewEquipBMSTerminateRequest(sn, pod, msgId string, p *services.Protocol) *equipBMSTerminateRequest {
	return &equipBMSTerminateRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipBMSTerminateResponse struct {
	api.Response
}

func (r *equipBMSTerminateResponse) GetStatus() int {
	return r.Status
}

func (r *equipBMSTerminateResponse) GetMsg() string {
	return r.Msg
}

func BMSTerminateRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
