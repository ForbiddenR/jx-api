package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipSetPriceSchemeRequest{}

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

func NewEquipSetPriceSchemeCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipSetPriceSchemeRequest {
	req := &equipSetPriceSchemeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetPriceSchemeCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipSetPriceSchemeRequest {
	req := &equipSetPriceSchemeRequest{
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

var _ services.Response = &equipSetIntellectChargeResponse{}

type equipSetPriceSchemeResponse struct {
	api.Response
}

func (resp *equipSetPriceSchemeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetPriceSchemeResponse) GetMsg() string {
	return resp.Msg
}

func SetPriceSchemeRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.SetPriceScheme)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipSetPriceSchemeResponse{})
}