package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipChargerTerminateRequest struct {
	services.Base
}

func (r *equipChargerTerminateRequest) GetName() services.Request2ServicesNameType {
	return services.ChargerTerminate
}

func (r *equipChargerTerminateRequest) TraceId() string {
	return r.MsgID
}

func (r *equipChargerTerminateRequest) IsCallback() bool {
	return false
}

func NewEquipChargerTerminateRequest(sn, pod, msgId string, p *services.Protocol) *equipChargerTerminateRequest {
	return &equipChargerTerminateRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipChargerTerminateResponse struct {
	api.Response
}

func (r *equipChargerTerminateResponse) GetStatus() int {
	return r.Status
}

func (r *equipChargerTerminateResponse) GetMsg() string {
	return r.Msg
}

func ChargerTerminateRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
