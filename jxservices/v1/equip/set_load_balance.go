package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
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

func NewEquipSetLoadBalanceRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetLoadBalanceRequest {
	req := &equipSetLoadBalanceRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipSetLoadBalanceRequestDetail{},
	}
	return req
}

func NewEquipSetLoadBalanceRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetLoadBalanceRequest {
	req := &equipSetLoadBalanceRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipSetLoadBalanceRequestDetail{},
	}
	return req
}

// var _ services.Response = &equipSetLoadBalanceResponse{}

// type equipSetLoadBalanceResponse struct {
// 	api.Response
// 	Data *equipSetLoadBalanceResponseDetail `json:"data"`
// }

// func (e *equipSetLoadBalanceResponse) GetStatus() int {
// 	return e.Status
// }

// func (e *equipSetLoadBalanceResponse) GetMsg() string {
// 	return e.Msg
// }

// type equipSetLoadBalanceResponseDetail struct {
// }

func SetLoadBalanceRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
