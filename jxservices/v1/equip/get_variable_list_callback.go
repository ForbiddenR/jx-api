package equip

import (
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipGetVariablesCallbackRequest{}

type equipGetVariableListCallbackRequest struct {
	services.Base
	Callback services.CB      `json:"callback"`
	Data     []CustomVariable `json:"data"`
}

func (equipGetVariableListCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.GetConfiguration
}

func (e *equipGetVariableListCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipGetVariableListCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipGetVariableListCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipGetVariableListCallbackRequest {
	req := &equipGetVariableListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipGetVariableListRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipGetVariableListCallbackRequest {
	req := &equipGetVariableListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}
