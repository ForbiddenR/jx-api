package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipClearCacheCallbackRequest{}

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

func NewEquipClearCacheCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipClearCacheCallbackRequest {
	req := &equipClearCacheCallbackRequest{
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

func NewEquipClearCacheCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipClearCacheCallbackRequest {
	req := &equipClearCacheCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}

	return req
}

var _ services.Response = &equipClearCacheCallbackResponse{}

type equipClearCacheCallbackResponse struct {
	api.Response
}

func (resp *equipClearCacheCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipClearCacheCallbackResponse) GetMsg() string {
	return resp.Msg
}

func ClearCacheCallbackRequest(ctx context.Context, req services.Request) error {
	// header := services.GetCallbackHeaderValue(services.ClearCache)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipClearCacheCallbackResponse{})
}
