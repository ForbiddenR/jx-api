package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
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

func NewEquipSetFactoryResetRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetFactoryResetRequest {
	req := &equipSetFactoryResetRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipSetFactoryResetRequestDetail{},
	}
	return req
}

func NewEquipSetFactoryResetRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetFactoryResetRequest {
	req := &equipSetFactoryResetRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipSetFactoryResetRequestDetail{},
	}
	return req
}

// type equipSetFactoryResetResponse struct {
// 	api.Response
// 	Data *equipSetFactoryResetResponseDetail `json:"data"`
// }

// var _ services.Response = &equipSetFactoryResetResponse{}

// func (e *equipSetFactoryResetResponse) GetStatus() int {
// 	return e.Status
// }

// func (e *equipSetFactoryResetResponse) GetMsg() string {
// 	return e.Msg
// }

// type equipSetFactoryResetResponseDetail struct {
// }

func SetFactoryResetRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
