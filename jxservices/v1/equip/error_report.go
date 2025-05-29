package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipErrorReportRequest struct {
	services.Base
}

func (r *equipErrorReportRequest) GetName() services.Request2ServicesNameType {
	return services.ErrorReport
}

func (r *equipErrorReportRequest) TraceId() string {
	return r.MsgID
}

func (r *equipErrorReportRequest) IsCallback() bool {
	return false
}

func NewEquipErrorReportRequest(sn, pod, msgId string, p *services.Protocol) *equipErrorReportRequest {
	return &equipErrorReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipErrorReportResponse struct {
	api.Response
}

func (r *equipErrorReportResponse) GetStatus() int {
	return r.Status
}

func (r *equipErrorReportResponse) GetMsg() string {
	return r.Msg
}

func ErrorReportRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
