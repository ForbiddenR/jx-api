package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipBootNotificationRequest{}

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

func NewEquipBootNotificationRequest(sn, id, pod, msgId string, p *services.Protocol) *equipBootNotificationRequest {
	request := &equipBootNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
	request.Data = &equipBootNotificationRequestDetail{}
	return request
}

var _ services.Response = &equipBootNotificationResponse{}

type equipBootNotificationResponse struct {
	api.Response
}

func (resp *equipBootNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipBootNotificationResponse) GetMsg() string {
	return resp.Msg
}

func BootNotificationRequest(ctx context.Context, req *equipBootNotificationRequest) error {
	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipBootNotificationResponse{})
}
