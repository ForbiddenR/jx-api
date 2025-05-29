package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipClearCacheCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipClearCacheCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.ClearCache
}

func (e *equipClearCacheCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipClearCacheCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipClearCacheCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipClearCacheCallbackRequest {
	req := &equipClearCacheCallbackRequest{
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

func NewEquipClearCacheCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipClearCacheCallbackRequest {
	req := &equipClearCacheCallbackRequest{
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

// var _ services.Response = &equipClearCacheCallbackResponse{}

// type equipClearCacheCallbackResponse struct {
// 	api.Response
// }

// func (resp *equipClearCacheCallbackResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipClearCacheCallbackResponse) GetMsg() string {
// 	return resp.Msg
// }

func ClearCacheCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
