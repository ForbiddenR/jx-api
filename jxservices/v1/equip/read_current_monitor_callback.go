package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipReadCurrentMonitorCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipReadCurrentMonitorCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.ReadCurrentMonitor
}

func (r *equipReadCurrentMonitorCallbackRequest) TraceId() string {
	return r.MsgID
}

func (equipReadCurrentMonitorCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipReadCurrentMonitorCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipReadCurrentMonitorCallbackRequest {
	req := &equipReadCurrentMonitorCallbackRequest{
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

func NewEquipReadCurrentMonitorCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipReadCurrentMonitorCallbackRequest {
	req := &equipReadCurrentMonitorCallbackRequest{
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

func ReadCurrentMonitorCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
