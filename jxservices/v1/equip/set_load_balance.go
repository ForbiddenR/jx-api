package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipSetLoadBalanceRequest struct {
	services.Base
	Callback services.CB                       `json:"callback"`
	Data     *equipSetLoadBalanceRequestDetail `json:"data"`
}

type equipSetLoadBalanceRequestDetail struct {
}

func (*equipSetLoadBalanceRequest) GetName() services.Request2ServicesNameType {
	return services.SetLoadBalance
}

func (e *equipSetLoadBalanceRequest) TraceId() string {
	return e.MsgID
}

func (equipSetLoadBalanceRequest) IsCallback() bool {
	return true
}

func NewEquipSetLoadBalanceRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipSetLoadBalanceRequest {
	req := &equipSetLoadBalanceRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.SetLoadBalance.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
		Data:     &equipSetLoadBalanceRequestDetail{},
	}
	return req
}

func NewEquipSetLoadBalanceRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipSetLoadBalanceRequest {
	req := &equipSetLoadBalanceRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.SetLoadBalance.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
		Data:     &equipSetLoadBalanceRequestDetail{},
	}
	return req
}

var _ services.Response = &equipSetLoadBalanceResponse{}

type equipSetLoadBalanceResponse struct {
	api.Response
	Data *equipSetLoadBalanceResponseDetail `json:"data"`
}

func (e *equipSetLoadBalanceResponse) GetStatus() int {
	return e.Status
}

func (e *equipSetLoadBalanceResponse) GetMsg() string {
	return e.Msg
}

type equipSetLoadBalanceResponseDetail struct {
}

func SetLoadBalanceRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.SetLoadBalance)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipSetLoadBalanceResponse{})
}
