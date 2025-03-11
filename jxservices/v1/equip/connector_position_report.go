package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipConnectorPositionReport struct {
	services.Base
	Data *equipConnectorPositionReportDetail `json:"data"`
}

type equipConnectorPositionReportDetail struct {
	Released        bool   `json:"released"`
	ConnectorSerial string `json:"connectorSerial"`
}

func (equipConnectorPositionReport) GetName() services.Request2ServicesNameType {
	return services.ConnectorPositionReport
}

func (e *equipConnectorPositionReport) TraceId() string {
	return e.MsgID
}

func (equipConnectorPositionReport) IsCallback() bool {
	return false
}

func NewEquipConnectorPositionReportRequest(sn, id, pod, msgId string, p *services.Protocol, released bool, connectorId string) *equipConnectorPositionReport {
	return &equipConnectorPositionReport{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ConnectorPositionReport.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipConnectorPositionReportDetail{
			Released:        released,
			ConnectorSerial: connectorId,
		},
	}
}

func ConnectorPositionReport(ctx context.Context, req *equipConnectorPositionReport) error {
	return services.Transport(ctx, req)
}
