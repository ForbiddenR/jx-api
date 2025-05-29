package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipChargingHandshakeRequest struct {
	services.Base
}

func (r *equipChargingHandshakeRequest) GetName() services.Request2ServicesNameType {
	return services.ChargingHandshake
}

func (r *equipChargingHandshakeRequest) TraceId() string {
	return r.MsgID
}

func (r *equipChargingHandshakeRequest) IsCallback() bool {
	return false
}

func NewEquipChargingHandshakeRequest(sn, pod, msgId string, p *services.Protocol) *equipChargingHandshakeRequest {
	return &equipChargingHandshakeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipChargingHandshakeResponse struct {
	api.Response
}

func (r *equipChargingHandshakeResponse) GetStatus() int {
	return r.Status
}

func (r *equipChargingHandshakeResponse) GetMsg() string {
	return r.Msg
}

func ChargingHandshakeRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
