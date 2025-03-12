package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipClearChargingProfileRequest{}

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

func NewClearChargingProfileCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipClearChargingProfileRequest {
	req := &equipClearChargingProfileRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ClearChargingProfile.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewClearChargingProfileCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipClearChargingProfileRequest {
	req := &equipClearChargingProfileRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ClearChargingProfile.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}

	return req
}

var _ services.Response = &equipClearChargingProfileResponse{}

type equipClearChargingProfileResponse struct {
	api.Response
}

func (resp *equipClearChargingProfileResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipClearChargingProfileResponse) GetMsg() string {
	return resp.Msg
}

func ClearChargingProfileRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.ClearChargingProfile)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipClearChargingProfileResponse{})
}
