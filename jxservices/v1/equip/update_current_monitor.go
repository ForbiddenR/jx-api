package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipUpdateCurrentMonitorRequest struct {
	services.Base
}

func (r *equipUpdateCurrentMonitorRequest) GetName() services.Request2ServicesNameType {
	return services.UpdateCurrentMonitor
}

func (r *equipUpdateCurrentMonitorRequest) TraceId() string {
	return r.MsgID
}

func (r *equipUpdateCurrentMonitorRequest) IsCallback() bool {
	return false
}

func NewUpdateCurrentMonitorRequest(sn, pod, msgId string, p *services.Protocol) *equipUpdateCurrentMonitorRequest {
	return &equipUpdateCurrentMonitorRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipUpdateCurrentMonitorResponse struct {
	api.Response
}

func (r *equipUpdateCurrentMonitorResponse) GetMsg() string {
	return r.Msg
}

func (r *equipUpdateCurrentMonitorResponse) GetStatus() int {
	return r.Status
}

func UpdateCurrentMonitorRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
