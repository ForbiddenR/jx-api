package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
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

type equipNotifyReportResponse struct {
	api.Response
}

func NewEquipNotifyReportRequest(sn, id, pod, msgId string, p *services.Protocol, requestId int64, tbc bool, reportDatas ...ReportData) *equipNotifyReportRequest {
	return &equipNotifyReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			AccessPod:   pod,
			MsgID:       msgId,
			Protocol:    p,
			Category:    services.NotifyReport.FirstUpper(),
		},
		Data: &equipNotifyReportRequestDetail{
			RequestId:  requestId,
			TBC:        tbc,
			ReportData: reportDatas,
		},
	}
}

func (resp *equipNotifyReportResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipNotifyReportResponse) GetMsg() string {
	return resp.Msg
}

func NotifyReportRequest(ctx context.Context, req *equipNotifyReportRequest) error {
	header := services.GetSimpleHeaderValue(services.NotifyReport)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipNotifyReportResponse{})
}
