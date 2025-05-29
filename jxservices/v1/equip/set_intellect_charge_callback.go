package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipSetIntellectChargeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipSetIntellectChargeRequest) GetName() services.Request2ServicesNameType {
	return services.SetIntellectCharge
}

func (e *equipSetIntellectChargeRequest) TraceId() string {
	return e.MsgID
}

func (equipSetIntellectChargeRequest) IsCallback() bool {
	return true
}

func NewEquipSetIntellectChargeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetIntellectChargeRequest {
	req := &equipSetIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetIntellectChargeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetIntellectChargeRequest {
	req := &equipSetIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

// var _ services.Response = &equipSetIntellectChargeResponse{}

// type equipSetIntellectChargeResponse struct {
// 	api.Response
// }

// func (resp *equipSetIntellectChargeResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipSetIntellectChargeResponse) GetMsg() string {
// 	return resp.Msg
// }

func SetIntellectChargeRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
