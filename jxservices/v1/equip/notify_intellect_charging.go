package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
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

func NewEequipNotifyIntellectChargingRequest(base services.Base, connectorId string, typ uint8, intellectId, startTime string, status uint8) *equipNotifyIntellectChargingRequest {
	return &equipNotifyIntellectChargingRequest{
		Base: base,
		Data: &equipNotifyIntellectChargingRequestData{
			EVSE: EVSE{
				Id:          "0",
				ConnectorId: connectorId,
			},
			IntellectType: typ,
			IntellectId:   intellectId,
			StartTime:     startTime,
			Status:        status,
		},
	}
}

func NotifyIntellectChargingRequest(ctx context.Context, req services.Request) error {
	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &api.Response{})
}
