package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipResetCallbackRequest{}

type equipResetCallbackRequest struct {
	services.Base
	Callback services.CB                      `json:"callback"`
	Data     *equipResetCallbackRequestDetail `json:"data"`
}

type equipResetCallbackRequestDetail struct {
}

func (equipResetCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.Reset
}

func (e *equipResetCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipResetCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipResetCallbackRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipResetCallbackRequest {
	req := &equipResetCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.Reset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}

	return req
}

func NewEquipResetCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipResetCallbackRequest {
	req := &equipResetCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Reset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipResetResponse{}

type equipResetResponse struct {
	api.Response
	Data *equipResetResponseDetail `json:"data"`
}

func (resp *equipResetResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipResetResponse) GetMsg() string {
	return resp.Msg
}

type equipResetResponseDetail struct {
}

func ResetCallbackRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.Reset)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipResetResponse{})
}
