package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipSetFactoryResetRequest{}

type equipSetFactoryResetRequest struct {
	services.Base
	Callback services.CB                        `json:"callback"`
	Data     *equipSetFactoryResetRequestDetail `json:"data"`
}

type equipSetFactoryResetRequestDetail struct {
}

func (*equipSetFactoryResetRequest) GetName() services.Request2ServicesNameType {
	return services.SetFactoryReset
}

func (e *equipSetFactoryResetRequest) TraceId() string {
	return e.MsgID
}

func (equipSetFactoryResetRequest) IsCallback() bool {
	return true
}

func NewEquipSetFactoryResetRequest(sn, id, pod, msgId string, p *services.Protocol, status int) *equipSetFactoryResetRequest {
	req := &equipSetFactoryResetRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.SetFactoryReset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
		Data:     &equipSetFactoryResetRequestDetail{},
	}
	return req
}

func NewEquipSetFactoryResetRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipSetFactoryResetRequest {
	req := &equipSetFactoryResetRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.SetFactoryReset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
		Data:     &equipSetFactoryResetRequestDetail{},
	}
	return req
}

type equipSetFactoryResetResponse struct {
	api.Response
	Data *equipSetFactoryResetResponseDetail `json:"data"`
}

var _ services.Response = &equipSetFactoryResetResponse{}

func (e *equipSetFactoryResetResponse) GetStatus() int {
	return e.Status
}

func (e *equipSetFactoryResetResponse) GetMsg() string {
	return e.Msg
}

type equipSetFactoryResetResponseDetail struct {
}

func SetFactoryResetRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.SetFactoryReset)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetFactoryResetResponse{})
}
