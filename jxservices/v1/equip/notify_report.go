package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipNotifyReportRequest{}

type equipNotifyReportRequest struct {
	services.Base
	Data *equipNotifyReportRequestDetail `json:"data"`
}

func (equipNotifyReportRequest) GetName() services.Request2ServicesNameType {
	return services.NotifyReport
}

func (e *equipNotifyReportRequest) TraceId() string {
	return e.MsgID
}

func (equipNotifyReportRequest) IsCallback() bool {
	return false
}

type equipNotifyReportRequestDetail struct {
	RequestId  int64        `json:"requestId"`
	TBC        bool         `json:"tbc"`
	ReportData []ReportData `json:"reportData"`
}

type ReportData struct {
	Component         Component         `json:"component"`
	Key               string            `json:"key"`
	VariableAttribute VariableAttribute `json:"variableAttribute"`
}

func NewEquipNotifyReportRequest(sn, pod, msgID string, p *services.Protocol, requestId int64, tbc bool, reportDatas ...ReportData) *equipNotifyReportRequest {
	return &equipNotifyReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			AccessPod:   pod,
			MsgID:       msgID,
			Protocol:    p,
		},
		Data: &equipNotifyReportRequestDetail{
			RequestId:  requestId,
			TBC:        tbc,
			ReportData: reportDatas,
		},
	}
}

// type equipNotifyReportResponse struct {
// 	api.Response
// }

// func (resp *equipNotifyReportResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipNotifyReportResponse) GetMsg() string {
// 	return resp.Msg
// }

func NotifyReportRequest(ctx context.Context, req *equipNotifyReportRequest) error {
	return services.Transport(ctx, req)
}
