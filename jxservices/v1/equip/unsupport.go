package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipUnsupportRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
	name     services.Request2ServicesNameType
}

func (e *equipUnsupportRequest) GetName() services.Request2ServicesNameType {
	return e.name
}

func (r *equipUnsupportRequest) TraceId() string {
	return r.MsgID
}

func (r *equipUnsupportRequest) IsCallback() bool {
	return true
}

func NewEquipUnsupportRequest(sn, pod, msgId string, p *services.Protocol, name string, err *apierrors.CallbackError) *equipUnsupportRequest {
	req := &equipUnsupportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Request2ServicesNameType(name).FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		name:     services.Request2ServicesNameType(name),
		Callback: services.NewCBError(err),
	}
	return req
}

func UnsupportRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
