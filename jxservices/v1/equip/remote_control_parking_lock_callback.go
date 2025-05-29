package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipRemoteControlParkingLockCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipRemoteControlParkingLockCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.RemoteControlParkingLock
}

func (r *equipRemoteControlParkingLockCallbackRequest) TraceId() string {
	return r.MsgID
}

func (r *equipRemoteControlParkingLockCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipRemoteControlParkingLockCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipRemoteControlParkingLockCallbackRequest {
	return &equipRemoteControlParkingLockCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipRemoteControlParkingLockCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteControlParkingLockCallbackRequest {
	return &equipRemoteControlParkingLockCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
}

func RemoteControlParkingLockCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
