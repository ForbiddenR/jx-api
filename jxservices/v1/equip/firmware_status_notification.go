package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type FirmwareStatusNotificationType int

const (
	Idle               FirmwareStatusNotificationType = 1
	Downloading        FirmwareStatusNotificationType = 2
	Downloaded         FirmwareStatusNotificationType = 3
	DownloadFailed     FirmwareStatusNotificationType = 4
	Installing         FirmwareStatusNotificationType = 5
	Installed          FirmwareStatusNotificationType = 6
	InstallationFailed FirmwareStatusNotificationType = 7
)

// var _ services.Request = &equipFirmwareStatusNotificationRequest{}

type equipFirmwareStatusNotificationRequest struct {
	services.Base
	Data *equipFirmwareStatusNotificationRequestDetail `json:"data"`
}

type equipFirmwareStatusNotificationRequestDetail struct {
	RequestId *uint32                        `json:"requestId,omitempty"`
	Status    FirmwareStatusNotificationType `json:"status"`
}

func (*equipFirmwareStatusNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.FirmwareStatusNotification
}

func (e *equipFirmwareStatusNotificationRequest) TraceId() string {
	return e.MsgID
}

func (*equipFirmwareStatusNotificationRequest) IsCallback() bool {
	return false
}

func NewEquipFirmwareStatusNotificationRequestOCPP16(sn, pod, msgID string, status FirmwareStatusNotificationType) *equipFirmwareStatusNotificationRequest {
	return &equipFirmwareStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipFirmwareStatusNotificationRequestDetail{
			Status: status,
		},
	}
}

func NewEquipFirmwareStatusNotificationRequest(sn, pod, msgID string, p *services.Protocol, requestID uint32, status FirmwareStatusNotificationType) *equipFirmwareStatusNotificationRequest {
	return &equipFirmwareStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipFirmwareStatusNotificationRequestDetail{
			RequestId: &requestID,
			Status:    status,
		},
	}
}

// var _ services.Response = &equipFirmwareStatusNotificationResponse{}

// type equipFirmwareStatusNotificationResponse struct {
// 	api.Response
// }

// func (resp *equipFirmwareStatusNotificationResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipFirmwareStatusNotificationResponse) GetMsg() string {
// 	return resp.Msg
// }

func FirmwareStatusNotificationRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
