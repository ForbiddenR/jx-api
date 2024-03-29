package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipSetPriceSchemeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipSetPriceSchemeRequest) GetName() string {
	return services.SetPriceScheme.String()
}

func NewEquipSetPriceSchemeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetPriceSchemeRequest {
	req := &equipSetPriceSchemeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetPriceScheme.FirstUpper(),
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
			Category:    services.SetPriceScheme.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
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

func SetPriceSchemeRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SetPriceScheme)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetPriceSchemeResponse{})
}
