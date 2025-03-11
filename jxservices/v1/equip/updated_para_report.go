package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipUpdatedParaReportRequest struct {
	services.Base
	Data *equipUpdatedParaReportRequestDetail `json:"data"`
}

type equipUpdatedParaReportRequestDetail struct {
	AuthFree           bool     `json:"authFree"`
	BtFastAuth         bool     `json:"btFastAuth"`
	CurrentMode        string   `json:"currentMode"`
	OutputType         string   `json:"outputType"`
	OutputCurrentLimit float64  `json:"outputLimitCurrent"`
	Entities           []Entity `json:"entities,omitempty"`
}

type Entity struct {
	Status      uint8  `json:"status"`
	Cycled      bool   `json:"cycled"`
	StartTime   string `json:"startTime"`
	StopTime    string `json:"stopTime"`
	IntellectId int64  `json:"intellectId"`
}

func (equipUpdatedParaReportRequest) GetName() services.Request2ServicesNameType {
	return services.UpdatedParaReport
}

func (e *equipUpdatedParaReportRequest) TraceId() string {
	return e.MsgID
}

func (equipUpdatedParaReportRequest) IsCallback() bool {
	return false
}

func NewEquipUpdatedParaReportRequest(sn, id, pod, msgId string, p *services.Protocol, authFree, btFastAuth bool, outputType string, outputCurrentLimit float64) *equipUpdatedParaReportRequest {
	return &equipUpdatedParaReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.UpdatedParaReport.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipUpdatedParaReportRequestDetail{
			AuthFree:           authFree,
			BtFastAuth:         btFastAuth,
			OutputType:         outputType,
			OutputCurrentLimit: outputCurrentLimit,
			Entities:           make([]Entity, 2),
		},
	}
}

func UpdatedParaReport(ctx context.Context, req *equipUpdatedParaReportRequest) error {
	return services.Transport(ctx, req)
}
