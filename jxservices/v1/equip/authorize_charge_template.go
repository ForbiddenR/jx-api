package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipAuthorizeChargeTemplateRequest struct {
	services.Base
	Data *equipAuthorizeChargeTemplateRequestDetail `json:"data"`
}

type equipAuthorizeChargeTemplateRequestDetail struct {
	ChargeTemplateId string `json:"priceSchemeId"`
}

func (r equipAuthorizeChargeTemplateRequest) GetName() services.Request2ServicesNameType {
	return services.AuthorizeChargeTemplate
}

func (r *equipAuthorizeChargeTemplateRequest) TraceId() string {
	return r.MsgID
}

func (equipAuthorizeChargeTemplateRequest) IsCallback() bool {
	return false
}

func NewEquipAuthorizeChargeTemplateRequest(sn, pod, msgId string, p *services.Protocol, chargeTemplateId string) *equipAuthorizeChargeTemplateRequest {
	return &equipAuthorizeChargeTemplateRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipAuthorizeChargeTemplateRequestDetail{
			ChargeTemplateId: chargeTemplateId,
		},
	}
}

type equipAuthorizeChargeTemplateResponse struct {
	api.Response
}

func (resp *equipAuthorizeChargeTemplateResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipAuthorizeChargeTemplateResponse) GetMsg() string {
	return resp.Msg
}

func AuthorizeChargeTemplateRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
