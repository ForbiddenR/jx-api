package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipAuthorizeTransactionRequest{}

type equipAuthorizeTransactionRequest struct {
	services.Base
	Data *equipAuthorizeTransactionRequestDetail `json:"data"`
}

type equipAuthorizeTransactionRequestDetail struct {
	IdTokenType IdTokenType `json:"idTokenType"`
}

func (r equipAuthorizeTransactionRequest) GetName() services.Request2ServicesNameType {
	return services.Authorize
}

func (r *equipAuthorizeTransactionRequest) TraceId() string {
	return r.MsgID
}

func (equipAuthorizeTransactionRequest) IsCallback() bool {
	return false
}

func NewEquipAuthorizeTransactionRequest(sn, id, pod, msgID string, p *services.Protocol) *equipAuthorizeTransactionRequest {
	return &equipAuthorizeTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.Authorize.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipAuthorizeTransactionRequestDetail{},
	}
}

var _ services.Response = &equipAuthorizeTransactionResponse{}

type equipAuthorizeTransactionResponse struct {
	api.Response
	Data *equipAuthorizeTransactionResponseDetail `json:"data"`
}

func (resp *equipAuthorizeTransactionResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipAuthorizeTransactionResponse) GetMsg() string {
	return resp.Msg
}

type equipAuthorizeTransactionResponseDetail struct {
	IdTokenInfo IdTokenInfo `json:"idTokenInfo" validate:"required"`
}

func AuthorizeTransactionRequest(ctx context.Context, req *equipAuthorizeTransactionRequest) (*equipAuthorizeTransactionResponse, error) {
	// header := services.GetSimpleHeaderValue(services.Authorize)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, &equipAuthorizeTransactionResponse{})
}
