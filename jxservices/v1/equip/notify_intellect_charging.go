package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipNotifyIntellectChargingRequest struct {
	services.Base
	Data *equipNotifyIntellectChargingRequestData `json:"data"`
}

type equipNotifyIntellectChargingRequestData struct {
	EVSE           EVSE    `json:"evse"`
	IntellectType  uint8   `json:"type"`
	IntellectId    string  `json:"intellectId"`
	StartTime      string  `json:"startTime"`
	EndTime        *string `json:"endTime"`
	EndElectricity *int    `json:"endElectricity"`
	EndSoc         *int    `json:"endSOC"`
	Status         uint8   `json:"status"`
}

func (r *equipNotifyIntellectChargingRequest) GetName() services.Request2ServicesNameType {
	return services.NotifyIntellectCharging
}

func (r *equipNotifyIntellectChargingRequest) TraceId() string {
	return r.MsgID
}

func (equipNotifyIntellectChargingRequest) IsCallback() bool {
	return false
}

func NewNotifyIntellectCharging(sn, id, pod, msgId string, p *services.Protocol) *equipNotifyIntellectChargingRequest {
	return &equipNotifyIntellectChargingRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipNotifyIntellectChargingRequestData{},
	}
}

func NotifyIntellectChargingRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
