package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipClearChargingProfileRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipClearChargingProfileRequest) GetName() services.Request2ServicesNameType {
	return services.ClearChargingProfile
}

func (e *equipClearChargingProfileRequest) TraceId() string {
	return e.MsgID
}

func (equipClearChargingProfileRequest) IsCallback() bool {
	return true
}

func NewClearChargingProfileCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipClearChargingProfileRequest {
	req := &equipClearChargingProfileRequest{
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

func NewClearChargingProfileCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipClearChargingProfileRequest {
	req := &equipClearChargingProfileRequest{
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

// var _ services.Response = &equipClearChargingProfileResponse{}

// type equipClearChargingProfileResponse struct {
// 	api.Response
// }

// func (resp *equipClearChargingProfileResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipClearChargingProfileResponse) GetMsg() string {
// 	return resp.Msg
// }

func ClearChargingProfileRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
