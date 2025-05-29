package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

// var _ services.Request = &equipGetBaseReportCallbackRequest{}

type equipGetBaseReportCallbackRequest struct {
	services.Base
	Callback services.CB                              `json:"callback"`
	Data     *equipGetBaseReportCallbackRequestDetail `json:"data"`
}

type equipGetBaseReportCallbackRequestDetail struct {
	// services.CB
	// lagacy
	Variable   []Variable `json:"variable"`
	Variables  []Variable `json:"variables"`
	UnknownKey []string   `json:"unknownKey,omitempty"`
}

type equipGetBaseReportCallbackRequestDetailVariable struct {
	Key      string  `json:"key"`
	Readonly bool    `json:"readonly"`
	Value    *string `json:"value,omitempty"`
}

func NewEquipGetBaseReportCallbackRequestDetailVariable(key string, readonly bool) equipGetBaseReportCallbackRequestDetailVariable {
	return equipGetBaseReportCallbackRequestDetailVariable{
		Key:      key,
		Readonly: readonly,
	}
}

func (g *equipGetBaseReportCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.GetBaseReport
}

func (g *equipGetBaseReportCallbackRequest) TraceId() string {
	return g.MsgID
}

func (g *equipGetBaseReportCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipGetBaseReportCallbackRequestOCPP16(sn, pod, msgID string, status int, length int, unknownKey []string) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipGetBaseReportCallbackRequestDetail{},
	}
	req.Data.Variable = make([]Variable, 0, length)
	req.Data.Variables = []Variable(nil)

	if len(unknownKey) > 0 {
		req.Data.UnknownKey = unknownKey
	}
	return req
}

func NewEquipGetBaseReportCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipGetBaseReportCallbackRequestDetail{},
	}
	return req
}

func NewEquipGetBaseReportRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipGetBaseReportCallbackRequestDetail{},
	}
	return req
}

func GetBaseReportCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
