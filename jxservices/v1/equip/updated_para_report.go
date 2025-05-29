package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipUpdatedParaReportRequest struct {
	services.Base
	Data *equipUpdatedParaReportRequestDetail `json:"data"`
}

type equipUpdatedParaReportRequestDetail struct {
	AuthFree           uint8    `json:"authFree"`
	BtFastAuth         uint8    `json:"btFastAuth"`
	OutputType         string   `json:"outputType"`
	OutputCurrentLimit float64  `json:"outputCurrentLimit"`
	Entities           []Entity `json:"entities"`
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

func NewEquipUpdatedParaReportRequest(sn, pod, msgID string, p *services.Protocol, authFree, btFastAuth uint8, outputType string, outputCurrentLimit float64) *equipUpdatedParaReportRequest {
	return &equipUpdatedParaReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
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
