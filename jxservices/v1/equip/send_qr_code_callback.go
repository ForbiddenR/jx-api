package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.CallbackRequest = &equipSendQRCodeRequest{}

type equipSendQRCodeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSendQRCodeRequest) GetName() services.Request2ServicesNameType {
	return services.SendQRCode
}

func (e *equipSendQRCodeRequest) TraceId() string {
	return e.MsgID
}

func (equipSendQRCodeRequest) IsCallback() bool {
	return true
}

func (e *equipSendQRCodeRequest) SetCallback(cb services.CB) {
	e.Callback = cb
}

func NewEquipSendQRCodeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSendQRCodeRequest {
	req := &equipSendQRCodeRequest{
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

func NewSendQRCodeCallbackRequest(base services.Base, option services.Option) *equipSendQRCodeRequest {
	req := &equipSendQRCodeRequest{
		Base: base,
	}
	option(req)
	return req
}

func NewEquipSendQRCodeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSendQRCodeRequest {
	req := &equipSendQRCodeRequest{
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

var _ services.Response = &equipSendQRCodeResponse{}

type equipSendQRCodeResponse struct {
	api.Response
}

func (resp *equipSendQRCodeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSendQRCodeResponse) GetMsg() string {
	return resp.Msg
}

func SendQRCodeRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
