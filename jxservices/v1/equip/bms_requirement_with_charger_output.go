package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipBMSRequirementWithChargerOutputRequest struct {
	services.Base
}

func (r *equipBMSRequirementWithChargerOutputRequest) GetName() services.Request2ServicesNameType {
	return services.BmsRequirementWithChargerOutput
}

func (r *equipBMSRequirementWithChargerOutputRequest) TraceId() string {
	return r.MsgID
}

func (r *equipBMSRequirementWithChargerOutputRequest) IsCallback() bool {
	return false
}

func NewEquipBMSRequirementWithChargerOutputRequest(sn, pod, msgId string, p *services.Protocol) *equipBMSRequirementWithChargerOutputRequest {
	req := &equipBMSRequirementWithChargerOutputRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
	return req
}

type equipBMSRequirementWithChargerOutputResponse struct {
	api.Response
}

func (r *equipBMSRequirementWithChargerOutputResponse) GetStatus() int {
	return r.Status
}

func (r *equipBMSRequirementWithChargerOutputResponse) GetMsg() string {
	return r.Msg
}

func BMSRequirementWithChargerOutputRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
