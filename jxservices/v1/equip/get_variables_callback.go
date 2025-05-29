package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

// var _ services.Request = &equipGetVariablesCallbackRequest{}

type equipGetVariablesCallbackRequest struct {
	services.Base
	Callback services.CB                             `json:"callback"`
	Data     *equipGetVariablesCallbackRequestDetail `json:"data"`
}

type equipGetVariablesCallbackRequestDetail struct {
	// services.CB
	// Variable   VariableAttribute `json:"variable"`
	Variable struct {
		Value    string     `json:"value"`
		Readonly Mutability `json:"readonly"`
	} `json:"variable"`
	Variables []Variable `json:"variables"`
	// Value      string     `json:"value"`
	// Readonly   Mutability `json:"readonly"`
	UnknownKey []string `json:"unknownKey,omitempty"`
}

func (equipGetVariablesCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.GetConfiguration
}

func (e *equipGetVariablesCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipGetVariablesCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipGetVariablesCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipGetVariablesCallbackRequest {
	req := &equipGetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipGetVariablesCallbackRequestDetail{},
	}
	req.Data.Variables = []Variable(nil)
	return req
}

func NewEquipGetVariablesRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetVariablesCallbackRequest {
	req := &equipGetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipGetVariablesCallbackRequestDetail{},
	}
	return req
}

// var _ services.Response = &equipGetVariablesCallbackResponse{}

// type equipGetVariablesCallbackResponse struct {
// 	api.Response
// 	Data *equipGetVariablesResponseDetail `json:"data"`
// }

// func (resp *equipGetVariablesCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipGetVariablesCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

// type equipGetVariablesResponseDetail struct {
// }

func GetVariablesCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
