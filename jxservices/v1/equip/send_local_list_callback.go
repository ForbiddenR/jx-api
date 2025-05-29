package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

const (
	SendLocalListAccept          = 0
	SendLocalListFailed          = 1
	SendLocalListVersionMismatch = 2
	SendLocalListNotSupported    = 3
)

var _ services.Request = &equipSendLocalListCallbackRequest{}

type equipSendLocalListCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSendLocalListCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SendLocalList
}

func (e *equipSendLocalListCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipSendLocalListCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSendLocalListCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSendLocalListCallbackRequest {
	req := &equipSendLocalListCallbackRequest{
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

func NewEquipSendLocalListCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSendLocalListCallbackRequest {
	req := &equipSendLocalListCallbackRequest{
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

var _ services.Response = &equipSendLocalListResponse{}

type equipSendLocalListResponse struct {
	api.Response
}

func (resp *equipSendLocalListResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSendLocalListResponse) GetMsg() string {
	return resp.Msg
}

func SendLocalListCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
