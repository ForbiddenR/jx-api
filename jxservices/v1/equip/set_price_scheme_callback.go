package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipSetPriceSchemeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSetPriceSchemeRequest) GetName() services.Request2ServicesNameType {
	return services.SetPriceScheme
}

func (e *equipSetPriceSchemeRequest) TraceId() string {
	return e.MsgID
}

func (equipSetPriceSchemeRequest) IsCallback() bool {
	return true
}

func NewEquipSetPriceSchemeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetPriceSchemeRequest {
	req := &equipSetPriceSchemeRequest{
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

func NewEquipSetPriceSchemeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetPriceSchemeRequest {
	req := &equipSetPriceSchemeRequest{
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

// type equipSetPriceSchemeResponse struct {
// 	api.Response
// }

// func (resp *equipSetPriceSchemeResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipSetPriceSchemeResponse) GetMsg() string {
// 	return resp.Msg
// }

func SetPriceSchemeRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
