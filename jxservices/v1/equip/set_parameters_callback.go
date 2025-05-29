package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipSetParametersCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSetParametersCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SetParameters
}

func (r *equipSetParametersCallbackRequest) TraceId() string {
	return r.MsgID
}

func (equipSetParametersCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSetParametersCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipSetParametersCallbackRequest {
	req := &equipSetParametersCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetParametersCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipSetParametersCallbackRequest {
	req := &equipSetParametersCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

func SetParametersCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
