package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
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

// func OCPP16GetFirmwareStatus(status protocol.FirmwareStatusNotificationJsonStatus) FirmwareStatusNotificationType {
// 	var result FirmwareStatusNotificationType
// 	switch status {
// 	case protocol.FirmwareStatusNotificationJsonStatusIdle:
// 		result = Idle
// 	case protocol.FirmwareStatusNotificationJsonStatusDownloading:
// 		result = Downloading
// 	case protocol.FirmwareStatusNotificationJsonStatusDownloadFailed:
// 		result = DownloadFailed
// 	case protocol.FirmwareStatusNotificationJsonStatusDownloaded:
// 		result = Downloaded
// 	case protocol.FirmwareStatusNotificationJsonStatusInstalling:
// 		result = Installing
// 	case protocol.FirmwareStatusNotificationJsonStatusInstallationFailed:
// 		result = InstallationFailed
// 	case protocol.FirmwareStatusNotificationJsonStatusInstalled:
// 		result = Installed
// 	}

// 	return result
// }

var _ services.Request = &equipFirmwareStatusNotificationRequest{}

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

func NewEquipFirmwareStatusNotificationRequestOCPP16(sn, id, pod, msgId string, status FirmwareStatusNotificationType) *equipFirmwareStatusNotificationRequest {
	return &equipFirmwareStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    services.OCPP16(),
			Category:    services.FirmwareStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipFirmwareStatusNotificationRequestDetail{
			Status: status,
		},
	}
}

func NewEquipFirmwareStatusNotificationRequest(sn, id, pod, msgId string, p *services.Protocol, requestID uint32, status FirmwareStatusNotificationType) *equipFirmwareStatusNotificationRequest {
	return &equipFirmwareStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.FirmwareStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipFirmwareStatusNotificationRequestDetail{
			RequestId: &requestID,
			Status:    status,
		},
	}
}

var _ services.Response = &equipFirmwareStatusNotificationResponse{}

type equipFirmwareStatusNotificationResponse struct {
	api.Response
}

func (resp *equipFirmwareStatusNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipFirmwareStatusNotificationResponse) GetMsg() string {
	return resp.Msg
}

func FirmwareStatusNotificationRequest(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.FirmwareStatusNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipFirmwareStatusNotificationResponse{})
}
