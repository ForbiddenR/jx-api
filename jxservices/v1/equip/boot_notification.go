package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipBootNotificationRequest struct {
	services.Base
	Data *equipBootNotificationRequestDetail `json:"data"`
}

type equipBootNotificationRequestDetail struct {
	ModelCode        string  `json:"modelCode"`
	ManufacturerCode string  `json:"manufacturerCode"`
	FirmwareVersion  *string `json:"firmwareVersion,omitempty"`
	Iccid            *string `json:"iccid,omitempty"`
	Imsi             *string `json:"imsi,omitempty"`
	BtName           *string `json:"btName,omitempty"`
	BtMac            *string `json:"btMac,omitempty"`
}

func (equipBootNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.BootNotification
}

func (e *equipBootNotificationRequest) TraceId() string {
	return e.MsgID
}

func (equipBootNotificationRequest) IsCallback() bool {
	return false
}

func NewEquipBootNotificationRequest(sn, pod, msgID string, p *services.Protocol) *equipBootNotificationRequest {
	request := &equipBootNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
	}
	request.Data = &equipBootNotificationRequestDetail{}
	return request
}

func BootNotificationRequest(ctx context.Context, req *equipBootNotificationRequest) error {
	return services.Transport(ctx, req)
}
