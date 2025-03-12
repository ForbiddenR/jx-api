package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipSetChargingTimerCallbackRequest{}

type equipSetChargingTimerCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSetChargingTimerCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SetChargingTimer
}

func (e *equipSetChargingTimerCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipSetChargingTimerCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSetChargingTimerCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.SetChargingTimer.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetChargingTimerCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.SetChargingTimer.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSetChargingTimerCallbackResponse{}

type equipSetChargingTimerCallbackResponse struct {
	api.Response
}

func (resp *equipSetChargingTimerCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetChargingTimerCallbackResponse) GetMsg() string {
	return resp.Msg
}

func SetChargingTimerCallbackRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.SetChargingTimer)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetChargingTimerCallbackResponse{})
}
