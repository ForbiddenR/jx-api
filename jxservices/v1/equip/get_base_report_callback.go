package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipGetBaseReportCallbackRequest{}

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

func NewEquipGetBaseReportCallbackRequestOCPP16(sn, id, pod, msgId string, status int, length int, unknownKey []string) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    services.OCPP16(),
			Category:    services.GetBaseReport.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
		Data:     &equipGetBaseReportCallbackRequestDetail{},
	}
	// if len(variable) > 0 {
	// 	req.Callback.Variable = variable
	// }

	req.Data.Variable = make([]Variable, 0, length)
	req.Data.Variables = []Variable(nil)

	if len(unknownKey) > 0 {
		req.Data.UnknownKey = unknownKey
	}
	return req
}

func NewEquipGetBaseReportCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.GetBaseReport.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
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
			Category:    services.GetBaseReport.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipGetBaseReportCallbackRequestDetail{},
	}
	return req
}

var _ services.Response = &equipGetBaseReportCallbackResponse{}

type equipGetBaseReportCallbackResponse struct {
	api.Response
	Data *equipGetBaseReportResponseDetail `json:"data"`
}

func (resp *equipGetBaseReportCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetBaseReportCallbackResponse) GetMsg() string {
	return resp.Msg
}

type equipGetBaseReportResponseDetail struct {
}

func GetBaseReportCallbackRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.GetConfiguration)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetBaseReportCallbackResponse{})
}
